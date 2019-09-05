package postgresql

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Executor struct {
	Query    string
	db       *sql.DB
	conninfo string
}

func New(conninfo string) *Executor {
	return &Executor{conninfo: conninfo}
}

func (e *Executor) Init() (err error) {
	e.db, err = sql.Open("postgres", e.conninfo)

	if err != nil {
		return err
	}

	err = e.db.Ping()

	return err
}

func (e *Executor) Run() (result interface{}, err error) {
	defer e.db.Close()

	err = e.db.Ping()
	if err != nil {
		return result, err
	}

	result, err = e.db.Exec(e.Query)
	if err != nil {
		return result, err
	}

	affectedRows, err := result.(sql.Result).RowsAffected()
	if err != nil {
		return result, err
	}

	log.Printf("Affected rows %d", affectedRows)
	return result, err
}

func (e *Executor) SetData(data interface{}) error {
	e.Query = data.(string)
	return nil
}
