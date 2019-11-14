package ticket

import (
	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/log"
)

type Ticket interface {
	AddComment(comment string) error
	Valid(user string) bool
}

type TicketApi interface {
	GetTicket(id string) (Ticket, error)
	GetCommentFormat() string
}

var TicketServive TicketApi

func init() {
	pconfig := config.GetConfig()
	if pconfig == nil {
		log.Fatalf("Configuration subsystem is not initialized")
	}
	ticketType := pconfig.GetString("ticket.type")
	if ticketType == "jira" {
		TicketServive = NewJiraApi()
	} else {
		TicketServive = NewFakeApi()
	}
}
