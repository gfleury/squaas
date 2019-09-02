package worker

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type BasicWorker struct {
	DataFeed    func() ([]interface{}, error)
	DataProcess func(interface{})
	MaxThreads  int
	ShouldStop  bool
	MinRunTime  time.Duration
	startTime   time.Time
	lastRunTime time.Duration
	wg          sync.WaitGroup
}

func (w *BasicWorker) Run() {

	for !w.ShouldStop {
		w.start()

		dataArray, err := w.DataFeed()

		if err != nil {
			log.Printf("Failed to get data: %s", err.Error())
			w.end()
			continue
		}

		for idx, data := range dataArray {
			if idx >= w.MaxThreads {
				break
			}

			w.wg.Add(1)

			go func(data interface{}) {
				defer w.wg.Done()
				w.DataProcess(data)
			}(data)
		}

		// Wait for all to finish
		w.wg.Wait()

		w.end()
	}
}

func (w *BasicWorker) start() {
	w.startTime = time.Now()
}

func (w *BasicWorker) end() {
	w.lastRunTime = time.Since(w.startTime)
	log.Printf("Worker run took %s", fmtDuration(w.lastRunTime))
	time.Sleep(w.MinRunTime)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}
