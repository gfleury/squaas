package worker

import (
	"sort"
	"sync"
	"time"

	check "gopkg.in/check.v1"
)

type testWorker struct {
	BasicWorker
	toProcessData []string
	processedData []string
	m             sync.Mutex
}

func (w *testWorker) DataFeed() ([]interface{}, error) {
	defer w.m.Unlock()
	w.m.Lock()
	ret := make([]interface{}, len(w.toProcessData))
	for idx, item := range w.toProcessData {
		ret[idx] = item
	}
	return ret, nil
}

func (w *testWorker) DataProcess(d interface{}) {
	defer w.m.Unlock()
	w.m.Lock()
	w.processedData = append(w.processedData, d.(string))

	for index, data := range w.toProcessData {
		if data == d.(string) {
			w.toProcessData = append(w.toProcessData[:index], w.toProcessData[index+1:]...)
			return
		}
	}

}

func (s *Suite) TestBasicWorker(c *check.C) {
	data := []string{"one", "two", "three", "four", "five"}
	w := testWorker{
		BasicWorker{MaxThreads: 2, MinRunTime: 10 * time.Second},
		data,
		[]string{},
		sync.Mutex{},
	}

	w.BasicWorker.DataFeed = w.DataFeed
	w.BasicWorker.DataProcess = w.DataProcess

	go w.Run()
	time.Sleep(2 * time.Second)
	w.ShouldStop.Set(true)

	w.m.Lock()
	sort.Slice(w.processedData, func(i, j int) bool {
		return w.processedData[i] < w.processedData[j]
	})
	w.m.Unlock()

	c.Assert(w.processedData, check.DeepEquals, []string{"one", "two"})
}
