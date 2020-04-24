package usecase

type Transaction interface {
	Transactionable()
}

type TransactionHandler interface {
	TransactAndReturnData(func(Transaction) (interface{}, error)) (interface{}, error)
}
