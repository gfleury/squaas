package worker

import (
	"time"

	check "gopkg.in/check.v1"

	"github.com/gfleury/squaas/db"
	"github.com/gfleury/squaas/models"
)

func (s *Suite) TestQueryWorker(c *check.C) {

	QueryDB := db.DBStorage.Connection().Model("Query")

	query := &models.Query{}

	err, _ := QueryDB.New(query)
	c.Assert(err, check.IsNil)

	query.Query =
		`Begin; 
	 CREATE TABLE userinfo
		(
			uid serial NOT NULL,
			username character varying(100) NOT NULL,
			departname character varying(500) NOT NULL,
			Created date,
			CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
		)
		WITH (OIDS=FALSE);
	INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;

		DROP TABLE userinfo;
		commit;`

	query.Status = models.StatusApproved
	query.ServerName = "good"

	err = query.Save()
	c.Assert(err, check.IsNil)

	w := NewQueryWorker()

	w.BasicWorker.MaxThreads = 2
	w.BasicWorker.MinRunTime = 10 * time.Second

	go w.Run()
	time.Sleep(2 * time.Second)
	w.ShouldStop.Set(true)

	err = QueryDB.FindId(query.GetId()).Exec(query)
	c.Assert(err, check.IsNil)

	c.Assert(query.Result, check.Equals, models.Result{AffectedRows: 12, Status: "", Success: true})
	c.Assert(query.Status, check.Equals, models.StatusDone)
}

func (s *Suite) TestQueryWorkerWrongQuery(c *check.C) {

	QueryDB := db.DBStorage.Connection().Model("Query")

	query := &models.Query{}

	err, _ := QueryDB.New(query)
	c.Assert(err, check.IsNil)

	query.Query =
		`Begin; 
	 
		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
		
		commit;`

	query.Status = models.StatusApproved
	query.ServerName = "good"

	err = query.Save()
	c.Assert(err, check.IsNil)

	w := NewQueryWorker()
	w.BasicWorker.MaxThreads = 2
	w.BasicWorker.MinRunTime = 10 * time.Second

	go w.Run()
	time.Sleep(2 * time.Second)
	w.ShouldStop.Set(true)

	err = QueryDB.FindId(query.GetId()).Exec(query)
	c.Assert(err, check.IsNil)

	c.Assert(query.Result, check.Equals, models.Result{AffectedRows: 0, Status: "pq: relation \"userinfo\" does not exist", Success: false})
	c.Assert(query.Status, check.Equals, models.StatusFailed)
}

func (s *Suite) TestQueryWorkerBrokenServerConnection(c *check.C) {

	QueryDB := db.DBStorage.Connection().Model("Query")

	query := &models.Query{}

	err, _ := QueryDB.New(query)
	c.Assert(err, check.IsNil)

	query.Query =
		`Begin;

		INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;

		commit;`

	query.Status = models.StatusApproved
	query.ServerName = "broken"

	err = query.Save()
	c.Assert(err, check.IsNil)

	w := NewQueryWorker()
	w.BasicWorker.MaxThreads = 2
	w.BasicWorker.MinRunTime = 10 * time.Second

	go w.Run()
	time.Sleep(2 * time.Second)
	w.ShouldStop.Set(true)

	err = QueryDB.FindId(query.GetId()).Exec(query)
	c.Assert(err, check.IsNil)

	c.Assert(query.Result, check.Equals, models.Result{})
	c.Assert(query.Status, check.Equals, models.StatusApproved)

	err = query.Delete()
	c.Assert(err, check.IsNil)
	err = query.Save()
	c.Assert(err, check.IsNil)
}
