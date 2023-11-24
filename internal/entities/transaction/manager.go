package transaction

import (
	"fmt"
	"time"
)

type (
	TransactionDataAdapter interface {
		Read(*Transaction) error
		List(*[]Transaction) error
		Create(Transaction) error
		Update(Transaction) error
		Delete(Transaction) error
		DeleteAll() error
		Uncategorized(*[]Transaction, int) error
		TransactionByDate(*[]Transaction, time.Time, time.Time) error
	}

	TransactionManager struct {
		transactionDataAdapter TransactionDataAdapter
	}
)

func NewTransactionManager(ta TransactionDataAdapter) *TransactionManager {
	return &TransactionManager{transactionDataAdapter: ta}
}

func (t *TransactionManager) Read(tra *Transaction) error {
	if tra.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return t.transactionDataAdapter.Read(tra)
}

func (t *TransactionManager) List(tra *[]Transaction) error {
	return t.transactionDataAdapter.List(tra)
}

func (t *TransactionManager) Create(tra Transaction) error {
	if tra.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	return t.transactionDataAdapter.Create(tra)
}

func (t *TransactionManager) Update(tra Transaction) error {
	if tra.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	if tra.CategoryId < 1 {
		return fmt.Errorf("Invalid CategoryId")
	}
	return t.transactionDataAdapter.Update(tra)
}

func (t *TransactionManager) Delete(tra Transaction) error {
	if tra.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return t.transactionDataAdapter.Delete(tra)
}

func (t *TransactionManager) DeleteAll() error {
	return t.transactionDataAdapter.DeleteAll()
}

func (t *TransactionManager) Uncategorized(transactions *[]Transaction, accountId int) error {
	return t.transactionDataAdapter.Uncategorized(transactions, accountId)
}

func (t *TransactionManager) TransactionByDate(transactions *[]Transaction, startDate, endDate time.Time) error {
	return t.transactionDataAdapter.TransactionByDate(transactions, startDate, endDate)
}
