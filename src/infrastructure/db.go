package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Sqlhandler struct {
	Conn *sql.DB
}

func NewSqlhandler() *Sqlhandler {
	db := connect()

	return &Sqlhandler{
		Conn: db,
	}
}

func (s *Sqlhandler) Close() {
	s.Close()
}

func connect() *sql.DB {
	connectionURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	db, err := sql.Open("mysql", connectionURL)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
