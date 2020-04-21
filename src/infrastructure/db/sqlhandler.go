package db

import (
	"database/sql"
	"fmt"
	"os"
	"task-api/src/interfaces"

	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	Conn *sql.DB
}

type Tx struct {
	Tx *sql.Tx
}

type Result struct {
	Result sql.Result
}

type Row struct {
	Rows *sql.Rows
}

func NewSqlhandler() interfaces.SQLHandler {
	db := connect()

	return &SQLHandler{
		Conn: db,
	}
}

func (s *SQLHandler) Close() {
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

func (s *SQLHandler) Begin() (interfaces.Tx, error) {
	t, err := s.Conn.Begin()
	if err != nil {
		return nil, err
	}

	tx := &Tx{t}
	return tx, nil
}

func (s *SQLHandler) Query(query string, args ...interface{}) (interfaces.Row, error) {
	rows, err := s.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	row := &Row{}
	row.Rows = rows

	return row, nil
}

func (s *SQLHandler) Exec(query string, args ...interface{}) (interfaces.Result, error) {
	result, err := s.Conn.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t Tx) Commit() error {
	err := t.Tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (t Tx) Rollback() error {
	err := t.Tx.Rollback()

	if err != nil {
		return err
	}

	return nil
}

func (t Tx) Exec(query string, args ...interface{}) (interfaces.Result, error) {
	result, err := t.Tx.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r Result) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r Result) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

func (r Row) Scan(value ...interface{}) error {
	return r.Rows.Scan(value...)
}

func (r Row) Next() bool {
	return r.Rows.Next()
}

func (r Row) Close() error {
	return r.Rows.Close()
}

func (r Row) Err() error {
	return r.Rows.Err()
}
