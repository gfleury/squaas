/*
 * DBqueryBench
 */

package main

import (
	"flag"
	"github.com/gfleury/squaas/worker"
	"log"
	"net/http"
	"time"

	"github.com/gfleury/squaas/api"
	_ "github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/db"
	_ "github.com/gfleury/squaas/ticket"
)

var web, workers, all bool

func init() {
	flag.BoolVar(&all, "all", true, "For running api+frontned+workers")
	flag.BoolVar(&web, "web", false, "For running only api+frontned")
	flag.BoolVar(&workers, "workers", false, "For running only the workers")
}

func main() {
	flag.Parse()

	if workers || web {
		all = false
	}

	db.InitStorage()
	err := db.DBStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	if workers || all {
		startWorkers(all || web)
	}

	if web || all {
		startWeb()
	}

}

func startWorkers(all bool) {
	log.Printf("Starting flow worker")
	wFlow := worker.NewFlowWorker()
	wFlow.BasicWorker.MaxThreads = 2
	wFlow.BasicWorker.MinRunTime = 10 * time.Second
	go wFlow.Run()

	log.Printf("Starting query workers/executors")
	w := worker.NewQueryWorker()
	w.BasicWorker.MaxThreads = 2
	w.BasicWorker.MinRunTime = 10 * time.Second
	if !all {
		w.Run()
	} else {
		go w.Run()
	}
}

func startWeb() {
	router := api.NewRouter()

	log.Printf("Starting server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
