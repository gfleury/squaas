/*
 * DBqueryBench
 */

package main

import (
	"log"
	"net/http"

	"github.com/gfleury/squaas/api"
	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/db"
)

func main() {
	config.Init()
	db.InitStorage()
	err := db.DBStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server started")

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
