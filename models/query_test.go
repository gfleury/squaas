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
