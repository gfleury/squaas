/*
 * DBworkBench
 */

package db

import (
	"log"

	"github.com/gfleury/squaas/config"
	"github.com/gfleury/squaas/models"

	"github.com/zebresel-com/mongodm"
	mgo "gopkg.in/mgo.v2"
)

var (
	DBStorage *Storage

	mongoConnectionURL string
)

func loadConfig() error {

	mongoConnectionURL = config.GetConfig().GetString("mongo.url")
	if mongoConnectionURL == "" {
		mongoConnectionURL = "mongodb://127.0.0.1:27017/squaas"
	}

	return nil
}

type Storage struct {
	db *mongodm.Connection
}

func (s *Storage) Init() error {

	localMap := map[string]map[string]string{
		"en-US": {
			"validation.field_required":           "Field '%s' is required.",
			"validation.field_invalid":            "Field '%s' has an invalid value.",
			"validation.field_invalid_id":         "Field '%s' contains an invalid object id value.",
			"validation.field_minlen":             "Field '%s' must be at least %v characters long.",
			"validation.field_maxlen":             "Field '%s' can be maximum %v characters long.",
			"validation.entry_exists":             "%s already exists for value '%v'.",
			"validation.field_not_exclusive":      "Only one of both fields can be set: '%s'' or '%s'.",
			"validation.field_required_exclusive": "Field '%s' or '%s' required.",
			"validation.field_invalid_relation11": "Field '%s' has wrong relation. Expected an array.",
			"validation.field_invalid_relation1n": "Field '%s' has wrong relation. No array expected.",
		},
	}

	dialInfo, err := mgo.ParseURL(mongoConnectionURL)

	if err != nil {
		return err
	}

	// Configure the mongodm connection and specify localisation map
	dbConfig := &mongodm.Config{
		DatabaseHosts:    dialInfo.Addrs,
		DatabaseName:     dialInfo.Database,
		DatabaseUser:     dialInfo.Username,
		DatabasePassword: dialInfo.Password,
		Locals:           localMap["en-US"],
	}

	// Connect and check for error
	db, err := mongodm.Connect(dbConfig)

	if err != nil {
		return err
	}

	s.db = db

	s.db.Register(&models.Query{}, "queries")

	// Create Indexes
	Query := s.db.Model("Query")
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = Query.EnsureIndex(index)
	if err != nil {
		return err
	}

	index = mgo.Index{
		Key:        []string{"status", "owner", "ticketid", "approvals.user"},
		Background: true,
		Sparse:     true,
	}
	err = Query.EnsureIndex(index)

	return err
}

func (s *Storage) Connection() *mongodm.Connection {
	err := s.db.Session.Ping()
	if err != nil {
		log.Printf("Ping failed, trying to reconnect, previous error: %s", err.Error())
		s.db.Session.Refresh()
	}
	return s.db
}

func InitStorage() {
	err := loadConfig()
	if err != nil {
		log.Fatalf("Database configuration load failed: %s", err.Error())
	}
	DBStorage = &Storage{}
}
