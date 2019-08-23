/*
 * DBworkBench
 */

package models

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/xwb1989/sqlparser"
	"github.com/zebresel-com/mongodm"
)

type Query struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`

	TicketID string `json:"ticketid" bson:"ticketid"`

	Owner User `json:"owner" bson:"owner"`

	Query string `json:"query" bson:"query"`

	Status string `json:"status,omitempty" bson:"status"`

	Approvals []*User `json:"approvals,omitempty" bson:"approvals"`

	HasSelect bool `json:"hasselect,omitempty" bson:"hasselect"`

	HasDelete bool `json:"hasdelete,omitempty" bson:"hasdelete"`

	HasInsert bool `json:"hasinsert,omitempty" bson:"hasinsert"`

	HasUpdate bool `json:"hasupdate,omitempty" bson:"hasupdate"`

	HasTransaction bool `json:"hastransaction,omitempty" bson:"hastransaction"`
}

func (q *Query) Byte() (objBytes []byte, err error) {
	return json.Marshal(q)
}

func (q *Query) Merge(eq *Query) (err error) {
	q.TicketID = eq.TicketID
	q.Query = eq.Query
	return err
}

func (q *Query) Parse(bodyReader io.Reader) error {
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, q)
	if err != nil {
		return err
	}

	return q.LintSQLQuery()
}

func (q *Query) AddApproval(u *User) {
	q.Approvals = append(q.Approvals, u)
}

func (q *Query) LintSQLQuery() error {
	hasBegin := false

	r := strings.NewReader(q.Query)
	tokens := sqlparser.NewTokenizer(r)
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch stmt.(type) {
		case *sqlparser.Select:
			q.HasSelect = true
		case *sqlparser.Insert:
			q.HasInsert = true
		case *sqlparser.Begin:
			hasBegin = true
		case *sqlparser.Rollback:
		case *sqlparser.Commit:
			if hasBegin {
				q.HasTransaction = true
			}
		case *sqlparser.Delete:
			q.HasDelete = true
		}
	}
	return nil
}
