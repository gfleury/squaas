package worker

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gfleury/squaas/log"
)

type BasicWorker struct {
	DataFeed    func() ([]interface{}, error)
	DataProcess func(interface{})
	MaxThreads  int
	ShouldStop  atomBool
	MinRunTime  time.Duration
	startTime   time.Time
	lastRunTime time.Duration
	wg          sync.WaitGroup
}

// https://gist.github.com/mahan/6256149
type atomBool struct{ flag int32 }

func (b *atomBool) Set(value bool) {
	var i int32 = 0
	if value {
		i = 1
	}
	atomic.StoreInt32(&(b.flag), int32(i))
}

func (b *atomBool) Get() bool {
	return atomic.LoadInt32(&(b.flag)) != 0
}

func (w *BasicWorker) Run() {

	for !w.ShouldStop.Get() {
		w.start()

		dataArray, err := w.DataFeed()

		if err != nil {
			log.Printf("Failed to get data: %s", err.Error())
			w.end(true)
			continue
		}

		// Just jump to the end if there is no data to process.
		if len(dataArray) == 0 {
			w.end(true)
			continue
		}

		shouldRunCount := w.MaxThreads
		for _, data := range dataArray {
			if shouldRunCount <= 0 {
				w.wg.Wait()
				shouldRunCount = w.MaxThreads
			}

			w.wg.Add(1)
			shouldRunCount--

			go func(data interface{}) {
				w.DataProcess(data)
				w.wg.Done()
			}(data)
		}

		// Wait for all to finish
		w.wg.Wait()

		w.end(false)
	}
}

func (w *BasicWorker) start() {
	w.startTime = time.Now()
}

func (w *BasicWorker) end(emptyRun bool) {
	w.lastRunTime = time.Since(w.startTime)
	if !emptyRun {
		log.Printf("Worker run took %s", fmtDuration(w.lastRunTime))
	}
	time.Sleep(w.MinRunTime)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
