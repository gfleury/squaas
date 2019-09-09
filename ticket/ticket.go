package ticket

import (
	"github.com/gfleury/squaas/config"
	"log"
)

type Ticket interface {
	AddComment(comment string) error
	Valid(user string) bool
}

type TicketApi interface {
	GetTicket(id string) (Ticket, error)
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