package db

import (
	"database/sql"
	"fmt"
	"os"
	"task-api/src/interfaces"
	"task-api/src/usecase"

	_ "github.com/go-sql-driver/mysql"
)

func NewSqlhandler() interfaces.SQLHandler {
	db := connect()

	return &SQLHandler{
		Conn: db,
	}
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

/* SQLHandler --------------- */
// トランザクションのためにTxを持たせている
// Txを保持している場合にはそれを使い、なければConnを使う
type SQLHandler struct {
	Conn *sql.DB
	Tx   *sql.Tx
}

func (s *SQLHandler) Close() {
	s.Close()
}

// Conn or Tx
func (s *SQLHandler) Query(query string, args ...interface{}) (interfaces.Rows, error) {
	var rows *sql.Rows
	var err error

	if s.Tx != nil {
		rows, err = s.Tx.Query(query, args...)
	} else {
		rows, err = s.Conn.Query(query, args...)
	}

	if err != nil {
		return nil, err
	}

	return &Rows{Rows: rows}, nil
}

// Conn or Tx
func (s *SQLHandler) QueryRow(query string, args ...interface{}) interfaces.Row {
	var r *sql.Row
	if s.Tx != nil {
		r = s.Tx.QueryRow(query, args...)
	} else {
		r = s.Conn.QueryRow(query, args...)
	}

	return &Row{Row: r}
}

// Conn or Tx
func (s *SQLHandler) Exec(query string, args ...interface{}) (interfaces.Result, error) {
	var result sql.Result
	var err error
	if s.Tx != nil {
		result, err = s.Tx.Exec(query, args...)
	} else {
		result, err = s.Conn.Exec(query, args...)
	}

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
	// トランザクションfieldをもたせてそのTxを使って処理をさせる
	data, err = txFunc(&SQLHandler{Tx: tx})
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

/* Result --------------- */
type Result struct {
	Result sql.Result
}

func (r Result) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r Result) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

/* Rows --------------- */
type Rows struct {
	Rows *sql.Rows
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

/* Row --------------- */
type Row struct {
	Row *sql.Row
}

func (r Row) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}
