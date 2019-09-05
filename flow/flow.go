package flow

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/db"
	"github.com/gfleury/squaas/models"
)

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
