package worker

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/db"
	"github.com/gfleury/squaas/log"
	"github.com/gfleury/squaas/models"
)

type FlowWorker struct {
	BasicWorker
}

type aggregatedQueries struct {
	Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Query         models.Query  `json:"query" bson:"query"`
	GoodApprovals int           `json:"goodApprovals" bson:"goodApprovals"`
}

func GetReadyQueries() (queries []*aggregatedQueries, err error) {

	QueryDB := db.DBStorage.Connection().Model("Query")

	err = QueryDB.Pipe(
		[]bson.M{
			bson.M{"$match": bson.M{
				"deleted": false,
				"status":  models.StatusReady,
				"approvals.approved": bson.M{
					"$eq": true,
				},
			},
			},
			bson.M{
				"$project": bson.M{
					"query": "$$ROOT",
					"goodApprovals": bson.M{
						"$size": bson.M{
							"$filter": bson.M{
								"input": "$approvals",
								"cond":  bson.M{"$eq": []interface{}{"$$this.approved", true}},
							},
						},
					},
				},
			}}).All(&queries)

	if err != nil {
		return
	}

	return
}

func (q *aggregatedQueries) ShouldBeApproved() bool {
	minApproved := config.GetConfig().GetInt("flow.minApproved")
	maxDisapproved := config.GetConfig().GetInt("flow.maxDisapproved")

	return q.GoodApprovals >= minApproved && (len(q.Query.Approvals)-q.GoodApprovals) < maxDisapproved
}

func NewFlowWorker() *FlowWorker {
	w := &FlowWorker{}
	w.BasicWorker.DataFeed = w.DataFeed
	w.BasicWorker.DataProcess = w.DataProcess
	return w
}

func (w *FlowWorker) DataFeed() (data []interface{}, err error) {
	queries, err := GetReadyQueries()
	if err != nil {
		return data, err
	}

	data = make([]interface{}, len(queries))

	for idx, query := range queries {
		data[idx] = query
	}

	return data, err
}

func (w *FlowWorker) DataProcess(data interface{}) {
	aggQuery := data.(*aggregatedQueries)
	log.Printf("FlowWorker: %s: Checking if we must approve this query", aggQuery.Query.Id.Hex())

	if aggQuery.ShouldBeApproved() {
		QueryDB := db.DBStorage.Connection().Model("Query")
		err := QueryDB.FindOne(aggQuery.Query).Exec(&aggQuery.Query)
		if err != nil {
			log.Fatalf("FlowWorker: %s: Cannot get the query on the MongoDB: %s", aggQuery.Query.Id.Hex(), err.Error())
		}

		aggQuery.Query.Status = models.StatusApproved
		err = aggQuery.Query.Save()
		if err != nil {
			log.Fatalf("FlowWorker: %s: Cannot update query: %s", aggQuery.Query.Id.Hex(), err.Error())
		}
	}
}
