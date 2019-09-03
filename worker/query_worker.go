package worker

import (
	"database/sql"
	"gopkg.in/mgo.v2/bson"
	"log"

	"github.com/gfleury/squaas/db"
	"github.com/gfleury/squaas/models"
	"github.com/gfleury/squaas/worker/executors/postgresql"
)

type QueryWorker struct {
	BasicWorker
}

func NewQueryWorker() *QueryWorker {
	w := &QueryWorker{}
	w.BasicWorker.DataFeed = w.DataFeed
	w.BasicWorker.DataProcess = w.DataProcess
	return w
}

func (w *QueryWorker) DataFeed() (data []interface{}, err error) {
	var queries []*models.Query

	QueryDB := db.DBStorage.Connection().Model("Query")

	err = QueryDB.Find(bson.M{"deleted": false, "status": models.StatusApproved}).Exec(&queries)

	data = make([]interface{}, len(queries))

	for idx, query := range queries {
		data[idx] = query
	}

	return data, err
}

func (w *QueryWorker) DataProcess(data interface{}) {
	var conninfo *string

	query := data.(*models.Query)
	log.Printf("Running %s", query.Query)

	servers := models.GetDatabases(true)

	for _, server := range servers {
		if server.Name == query.ServerName {
			conninfo = &server.Uri
			break
		}
	}

	if conninfo == nil {
		log.Printf("Database connection information not found, please update configuration. Server: %s", query.ServerName)
		return
	}

	executor := postgresql.New(*conninfo)

	err := executor.Init()

	if err != nil {
		log.Printf("Database connection initialization failed, not running. %s", err.Error())
		return
	}

	err = executor.SetData(query.Query)
	if err != nil {
		log.Printf("Database parameter initialization failed, not running. %s", err.Error())
		return
	}

	query.Status = models.StatusRunning
	err = query.Save()
	if err != nil {
		log.Printf("Query save failed, not running. %s", err.Error())
		return
	}

	result, err := executor.Run()

	query.Result.AffectedRows = 0

	if err != nil {
		log.Printf("Query failed to run. %s", err.Error())
		query.Status = models.StatusFailed
		query.Result.Success = false
		query.Result.Status = err.Error()
		err = query.Save()
		if err != nil {
			log.Printf("Query save after run failure, failed. Query will remain in status 'Running'.")
		}
		return
	}

	affectedRows, err := result.(sql.Result).RowsAffected()
	if err == nil {
		query.Result.AffectedRows = int(affectedRows)
	}
	query.Result.Success = true

	query.Status = models.StatusDone
	err = query.Save()
	if err != nil {
		log.Printf("Query deleting failed, not running. %s", err.Error())
		return
	}
}
