/*
 * DBworkBench
 */

package models

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/xwb1989/sqlparser"
	"github.com/zebresel-com/mongodm"
)

type Status string

const (
	StatusReady     Status = "ready"
	StatusDone      Status = "done"
	StatusPending   Status = "pending"
	StatusApproved  Status = "approved"
	StatusRunning   Status = "running"
	StatusFailed    Status = "failed"
	StatusParseOnly Status = "PARSEONLY"
)

type Approvals struct {
	User     *User `json:"user" bson:"user"`
	Approved bool  `json:"approved" bson:"approved"`
}

type Result struct {
	AffectedRows int    `json:"affectedrows" bson:"affectedrows"`
	Success      bool   `json:"success" bson:"success"`
	Status       string `json:"status" bson:"status"`
}

type Query struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`

	TicketID string `json:"ticketid" bson:"ticketid"`

	Owner User `json:"owner" bson:"owner"`

	ServerName string `json:"servername" bson:"servername"`

	Query string `json:"query" bson:"query"`

	Status Status `json:"status" bson:"status"`

	Approvals []Approvals `json:"approvals" bson:"approvals"`

	HasSelect bool `json:"hasselect,omitempty" bson:"hasselect"`

	HasDelete bool `json:"hasdelete,omitempty" bson:"hasdelete"`

	HasInsert bool `json:"hasinsert,omitempty" bson:"hasinsert"`

	HasUpdate bool `json:"hasupdate,omitempty" bson:"hasupdate"`

	HasAlter bool `json:"hasalter,omitempty" bson:"hasalter"`

	HasTransaction bool `json:"hastransaction,omitempty" bson:"hastransaction"`

	Result Result `json:"result,omitempty" bson:"result"`
}

func (q *Query) Byte() (objBytes []byte, err error) {
	return json.Marshal(q)
}

func (q *Query) Merge(eq *Query) (err error) {
	if q.Status == StatusPending {
		q.TicketID = eq.TicketID
		q.Query = eq.Query
		q.Status = eq.Status
	}

	if q.Status == StatusReady && eq.Status == StatusPending {
		q.Approvals = nil
		q.Status = eq.Status
	}

	return q.LintSQLQuery()
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

func (q *Query) AddApproval(u *User, approve bool) {
	for id, approval := range q.Approvals {
		if approval.User.Name == u.Name {
			q.Approvals[id].Approved = approve
			return
		}
	}
	q.Approvals = append(q.Approvals, Approvals{u, approve})
}

func (q *Query) LintSQLQuery() error {
	q.HasAlter = false
	q.HasSelect = false
	q.HasInsert = false
	q.HasTransaction = false
	q.HasDelete = false

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
		case *sqlparser.Update:
			if stmt.(*sqlparser.Update).Where == nil {
				return fmt.Errorf("No WHERE found for UPDATE")
			}
			q.HasUpdate = true
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
			if stmt.(*sqlparser.Delete).Where == nil {
				return fmt.Errorf("No WHERE found for DELETE")
			}
			q.HasDelete = true
		case *sqlparser.DDL:
			q.HasAlter = true
		}
	}
	return nil
}

var validID = regexp.MustCompile(`^[0-9a-fA-F]{24}$`)

func IsValidObjectId(id string) bool {
	return validID.MatchString(id)
}

func (s Status) Valid() bool {
	switch s {
	case StatusReady:
		return true
	case StatusDone:
		return true
	case StatusPending:
		return true
	case StatusApproved:
		return true
	case StatusRunning:
		return true
	case StatusFailed:
		return true
	case StatusParseOnly:
		return true
	default:
		return false
	}
}
