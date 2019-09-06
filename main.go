/*
 * DBqueryBench
 */

package main

import (
	"github.com/gfleury/squaas/worker"
	"log"
	"net/http"
	"time"

	"github.com/gfleury/squaas/api"
	_ "github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/db"
	_ "github.com/gfleury/squaas/ticket"
)

func main() {
	db.InitStorage()
	err := db.DBStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting flow worker")
	wFlow := worker.NewFlowWorker()
	wFlow.BasicWorker.MaxThreads = 2
	wFlow.BasicWorker.MinRunTime = 10 * time.Second
	go wFlow.Run()

	log.Printf("Starting query workers/executors")
	w := worker.NewQueryWorker()
	w.BasicWorker.MaxThreads = 2
	w.BasicWorker.MinRunTime = 10 * time.Second
	go w.Run()

	router := api.NewRouter()

	log.Printf("Starting server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
