package models

import (
	check "gopkg.in/check.v1"
)

func (s *Suite) TestModelSQLLint(c *check.C) {
	q := &Query{
		TicketID: "BLAH-123",
		Query:    "SELECT * FROM TABLEX; SELECT COUNT(*) FROM TABLEY;",
	}

	err := q.LintSQLQuery()
	c.Assert(err, check.IsNil)

	c.Assert(q.HasSelect, check.Equals, true)
	c.Assert(q.HasDelete, check.Equals, false)
	c.Assert(q.HasInsert, check.Equals, false)
	c.Assert(q.HasUpdate, check.Equals, false)
	c.Assert(q.HasTransaction, check.Equals, false)
}

func (s *Suite) TestModelSQLLintInsertSelect(c *check.C) {
	q := &Query{
		TicketID: "BLAH-123",
		Query:    "SELECT * FROM TABLEX; INSERT INTO a VALUES (1);",
	}

	err := q.LintSQLQuery()
	c.Assert(err, check.IsNil)

	c.Assert(q.HasSelect, check.Equals, true)
	c.Assert(q.HasDelete, check.Equals, false)
	c.Assert(q.HasInsert, check.Equals, true)
	c.Assert(q.HasUpdate, check.Equals, false)
	c.Assert(q.HasTransaction, check.Equals, false)
}

func (s *Suite) TestModelSQLLintTransaction(c *check.C) {
	q := &Query{
		TicketID: "BLAH-123",
		Query:    "BEGIN; INSERT INTO a VALUES (1); COMMIT;",
	}

	err := q.LintSQLQuery()
	c.Assert(err, check.IsNil)

	c.Assert(q.HasSelect, check.Equals, false)
	c.Assert(q.HasDelete, check.Equals, false)
	c.Assert(q.HasInsert, check.Equals, true)
	c.Assert(q.HasUpdate, check.Equals, false)
	c.Assert(q.HasTransaction, check.Equals, true)
}

func (s *Suite) TestModelSQLLintBrokenTransaction(c *check.C) {
	q := &Query{
		TicketID: "BLAH-123",
		Query:    "BEGIN; INSERT INTO a VALUES (1);INSERT INTO a VALUES (1);INSERT INTO a VALUES (1);",
	}

	err := q.LintSQLQuery()
	c.Assert(err, check.IsNil)

	c.Assert(q.HasSelect, check.Equals, false)
	c.Assert(q.HasDelete, check.Equals, false)
	c.Assert(q.HasInsert, check.Equals, true)
	c.Assert(q.HasUpdate, check.Equals, false)
	c.Assert(q.HasTransaction, check.Equals, false)
}

func (s *Suite) TestStatus(c *check.C) {
	st := StatusReady
	c.Assert(st.Valid(), check.Equals, true)
	st = "Wronglyz"
	c.Assert(st.Valid(), check.Equals, false)
}

func (s *Suite) TestMongoHexId(c *check.C) {
	st := "5d6900e11b4db412b3c5f7b1"
	c.Assert(IsValidObjectId(st), check.Equals, true)
	st = "Wronglyz"
	c.Assert(IsValidObjectId(st), check.Equals, false)
}

func (s *Suite) TestQueryAddRepeatedApprovals(c *check.C) {
	q := &Query{
		TicketID: "BLAH-123",
		Query:    "BEGIN; INSERT INTO a VALUES (1);INSERT INTO a VALUES (1);INSERT INTO a VALUES (1);",
	}

	q.AddApproval(&User{Name: "root"}, true)
	q.AddApproval(&User{Name: "root"}, true)
	q.AddApproval(&User{Name: "root"}, false)
	q.AddApproval(&User{Name: "root"}, false)

	c.Assert(len(q.Approvals), check.Equals, 1)
}

func (s *Suite) TestQueryAddDifferentApprovals(c *check.C) {
	q := &Query{
		TicketID: "BLAH-123",
		Query:    "BEGIN; INSERT INTO a VALUES (1);INSERT INTO a VALUES (1);INSERT INTO a VALUES (1);",
	}

	q.AddApproval(&User{Name: "root"}, true)
	q.AddApproval(&User{Name: "user1"}, true)
	q.AddApproval(&User{Name: "user3"}, false)
	q.AddApproval(&User{Name: "user9"}, false)

	c.Assert(len(q.Approvals), check.Equals, 4)
}

func (s *Suite) TestQueryUpdateTicketWithComment(c *check.C) {
	q := &Query{
		TicketID: "BLAH-123",
		Query:    "BEGIN; INSERT INTO a VALUES (1);INSERT INTO a VALUES (1);INSERT INTO a VALUES (1);",
	}

	err := q.TicketCommentAdded("")
	c.Assert(err, check.IsNil)
	err = q.TicketCommentDone()
	c.Assert(err, check.IsNil)
	err = q.TicketCommentFailed()
	c.Assert(err, check.IsNil)

}
