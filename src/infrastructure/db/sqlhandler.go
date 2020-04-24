package db

import (
	"database/sql"
	"fmt"
	"os"
	"task-api/src/interfaces"
	"task-api/src/usecase"

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

type Rows struct {
	Rows *sql.Rows
}

type Row struct {
	Row *sql.Row
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
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println("connection url: ", connectionURL)
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

func (s *SQLHandler) Query(query string, args ...interface{}) (interfaces.Rows, error) {
	rows, err := s.Conn.Query(query, args...)

	if err != nil {
		return nil, err
	}

	row := &Rows{}
	row.Rows = rows

	return row, nil
}

func (s *SQLHandler) QueryRow(query string, args ...interface{}) interfaces.Row {
	r := s.Conn.QueryRow(query, args...)

	row := &Row{}
	row.Row = r

	return row
}

func (s *SQLHandler) Exec(query string, args ...interface{}) (interfaces.Result, error) {
	result, err := s.Conn.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SQLHandler) TransactAndReturnData(txFunc func(usecase.Transaction) (interface{}, error)) (data interface{}, err error) {
	tx, err := s.Conn.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	data, err = txFunc(&Tx{Tx: tx})
	return
}

func (s *SQLHandler) Transactionable() {
	return
}

func (s *SQLHandler) FromTransaction(tx usecase.Transaction) interfaces.SQLHandler {
	sqlhandler, ok := tx.(*SQLHandler)
	if !ok {
		return s
	}

	return sqlhandler
}

func (t *Tx) Transactionable() {
	return
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

func (r Rows) Scan(value ...interface{}) error {
	return r.Rows.Scan(value...)
}

func (r Rows) Next() bool {
	return r.Rows.Next()
}

func (r Rows) Close() error {
	return r.Rows.Close()
}

func (r Rows) Err() error {
	return r.Rows.Err()
}

func (r Row) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}
