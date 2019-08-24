/*
 * DBqueryBench
 */

package main

import (
	"github.com/gfleury/dbquerybench/config"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gfleury/dbquerybench/api"
	"github.com/gfleury/dbquerybench/db"
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
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})

	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
}
