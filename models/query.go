/*
 * DBworkBench
 */

package models

import (
	"encoding/json"
	"fmt"
	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/ticket"
	"io"
	"io/ioutil"
	"regexp"
	// "strings"

	// "github.com/xwb1989/sqlparser"
	// "github.com/vitessio/vitess/go/vt/sqlparser"
	sqlparser "github.com/lfittl/pg_query_go"
	nodes "github.com/lfittl/pg_query_go/nodes"
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

	tree, err := sqlparser.Parse(q.Query)
	if err != nil {
		return err
	}

	for _, stmt := range tree.Statements {

		switch stmt := stmt.(nodes.RawStmt).Stmt.(type) {
		case nodes.UpdateStmt:
			if stmt.WhereClause == nil {
				return fmt.Errorf("No WHERE found for UPDATE")
			}
			q.HasUpdate = true
		case nodes.SelectStmt:
			q.HasSelect = true
		case nodes.InsertStmt:
			q.HasInsert = true
		case nodes.TransactionStmt:
			if stmt.Kind == nodes.TRANS_STMT_BEGIN {
				hasBegin = true
			} else if hasBegin && stmt.Kind == nodes.TRANS_STMT_COMMIT {
				q.HasTransaction = true
			}
		case nodes.DeleteStmt:
			if stmt.WhereClause == nil {
				return fmt.Errorf("No WHERE found for DELETE")
			}
			q.HasDelete = true
		case nodes.CreateStmt:
			q.HasAlter = true
		}
	}
	return nil
}

func (q *Query) TicketCommentFailed() error {
	httpReferer := config.GetConfig().GetString("mainDomain")
	t, err := ticket.TicketServive.GetTicket(q.TicketID)
	if err != nil {
		return err
	}

	err = t.AddComment(fmt.Sprintf(ticket.TicketServive.GetCommentFormat(), fmt.Sprintf("Query execution on %s {color:red}FAILED{color}: \n", q.ServerName), q.Result.Status, httpReferer+"/frontend/#/queries/edit/"+q.Id.Hex()))
	return err
}

func (q *Query) TicketCommentDone() error {
	httpReferer := config.GetConfig().GetString("mainDomain")
	t, err := ticket.TicketServive.GetTicket(q.TicketID)
	if err != nil {
		return err
	}

	err = t.AddComment(fmt.Sprintf(ticket.TicketServive.GetCommentFormat(), fmt.Sprintf("Query executed on %s with {color:green}SUCCESS{color}: \n Number of affected rows: \n", q.ServerName), q.Result.AffectedRows, httpReferer+"/frontend/#/queries/edit/"+q.Id.Hex()))
	return err
}

func (q *Query) TicketCommentAdded() error {
	httpReferer := config.GetConfig().GetString("mainDomain")
	t, err := ticket.TicketServive.GetTicket(q.TicketID)
	if err != nil {
		return err
	}

	err = t.AddComment(fmt.Sprintf(ticket.TicketServive.GetCommentFormat(), fmt.Sprintf("Added query into querybench, to run at %s: \n", q.ServerName), q.Query, httpReferer+"/frontend/#/queries/edit/"+q.Id.Hex()))
	return err
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
