package transaction

import (
	"fmt"
	"time"
)

func Read(tra *Transaction) error {
	if tra.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataRead(tra)
}

func List(tra *[]Transaction) error {
	return DataList(tra)
}

func Create(tra Transaction) error {
	if tra.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	// if tra.CategoryId < 1 {
	// 	return fmt.Errorf("Invalid CategoryId")
	// }
	return DataCreate(tra)
}

func Update(tra Transaction) error {
	if tra.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	if tra.CategoryId < 1 {
		return fmt.Errorf("Invalid CategoryId")
	}
	return DataUpdate(tra)
}

func Delete(tra Transaction) error {
	if tra.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataDelete(tra)
}

func DeleteAll() error {
	return DataDeleteAll()
}

func Uncategorized(transactions *[]Transaction, accountId int) error {
	return DataUncategorized(transactions, accountId)
}

func TransactionByDate(transactions *[]Transaction, startDate, endDate time.Time) error {
	allTransactions := []Transaction{}
	if err := DataList(&allTransactions); err != nil {
		return err
	}
	for _, t := range allTransactions {
		if (startDate.Equal(t.TxnDate) || startDate.Before(t.TxnDate)) && endDate.After(t.TxnDate) {
			*transactions = append(*transactions, t)
		}
	}
	return nil
}
