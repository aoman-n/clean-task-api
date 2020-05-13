package interfaces

import (
	"task-api/src/entity/repository"
)

type SQLHandler interface {
	Query(string, ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row
	Exec(string, ...interface{}) (Result, error)
	Close()
	TransactAndReturnData(txFunc func(repository.Transaction) (interface{}, error)) (data interface{}, err error)
	Transactionable()
	FromTransaction(tx repository.Transaction) SQLHandler
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
	Err() error
}

type Row interface {
	Scan(...interface{}) error
}
