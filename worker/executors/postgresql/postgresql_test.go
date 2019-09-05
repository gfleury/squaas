package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
	check "gopkg.in/check.v1"
)

func (s *Suite) TestExecutorInsert(c *check.C) {
	e := New("postgres://postgres@localhost/data?sslmode=disable")

	query := `Begin; INSERT INTO userinfo(username,departname,created) VALUES('astaxie', 'dsds', '2012-12-09') returning uid;
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
		commit;`

	err := e.Init()
	c.Assert(err, check.IsNil)

	err = e.SetData(query)
	c.Assert(err, check.IsNil)

	data, err := e.Run()
	c.Assert(err, check.IsNil)
	rows, err := data.(sql.Result).RowsAffected()
	c.Assert(err, check.IsNil)
	c.Assert(rows, check.Equals, int64(12))

}
