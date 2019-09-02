package postgresql

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	check "gopkg.in/check.v1"
)

type Suite struct {
}

func (s *Suite) SetUpSuite(c *check.C) {
	e := New("postgres://postgres@localhost/data?sslmode=disable")

	query := `Begin; CREATE TABLE userinfo
		(
			uid serial NOT NULL,
			username character varying(100) NOT NULL,
			departname character varying(500) NOT NULL,
			Created date,
			CONSTRAINT userinfo_pkey PRIMARY KEY (uid)
		)
		WITH (OIDS=FALSE);commit;`

	err := e.Init()
	c.Assert(err, check.IsNil)

	err = e.SetData(query)
	c.Assert(err, check.IsNil)

	_, err = e.Run()
	c.Assert(err, check.IsNil)

}

func (s *Suite) TearDownSuite(c *check.C) {
	e := New("postgres://postgres@localhost/data?sslmode=disable")

	query := `Begin; DROP TABLE userinfo;commit;`

	err := e.Init()
	c.Assert(err, check.IsNil)

	err = e.SetData(query)
	c.Assert(err, check.IsNil)

	data, err := e.Run()
	c.Assert(err, check.IsNil)
	rows, err := data.(sql.Result).RowsAffected()
	c.Assert(err, check.IsNil)
	c.Assert(rows, check.Equals, int64(0))
}

var _ = check.Suite(&Suite{})

func Test(t *testing.T) {
	check.TestingT(t)
}
