package worker

import (
	"gopkg.in/mgo.v2/bson"
	"log"

	"github.com/gfleury/squaas/db"
	"github.com/gfleury/squaas/models"
	"github.com/gfleury/squaas/worker/executors"
	"github.com/gfleury/squaas/worker/executors/postgresql"
)

type QueryWorker struct {
	BasicWorker
	executor executors.Executor
}

func New() *QueryWorker {
	w := &QueryWorker{}
	w.executor = postgresql.New("postgres://postgres@localhost/data?sslmode=disable")
	w.BasicWorker.DataFeed = w.DataFeed
	w.BasicWorker.DataProcess = w.DataProcess
	return w
}

func (w *QueryWorker) DataFeed() (queries []interface{}, err error) {
	QueryDB := db.DBStorage.Connection().Model("Query")

	err = QueryDB.Find(bson.M{"deleted": false, "status": models.StatusApproved}).Exec(&queries)

	return queries, err
}

func (w *QueryWorker) DataProcess(data interface{}) {
	query := data.(*models.Query)
	log.Printf("Running %s", query.Query)

	err := w.executor.Init()

	if err != nil {
		log.Printf("Database connection initialization failed, not running.")
		return
	}

	err = w.executor.SetData(query)
	if err != nil {
		log.Printf("Database parameter initialization failed, not running.")
		return
	}

	query.Status = models.StatusRunning
	err = query.Save()
	if err != nil {
		log.Printf("Query deleting failed, not running.")
		return
	}

	_, err = w.executor.Run()
	if err != nil {
		log.Printf("Query failed to run.")
		query.Status = models.StatusFailed
		err = query.Save()
		if err != nil {
			log.Printf("Query saving after run failure, failed. Query will remain in status 'Running'.")
		}
		return
	}

	query.Status = models.StatusDone
	err = query.Save()
	if err != nil {
		log.Printf("Query deleting failed, not running.")
		return
	}
}

func (w *QueryWorker) ShouldExecuteThisQuery(models.Query) error {
	return nil
}
