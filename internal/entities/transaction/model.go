package transaction

import (
	"time"

	"github.com/blackflagsoftware/prognos/config"
)

type (
	Transaction struct {
		Id           int       `db:"id" json:"id"`
		AccountId    int       `db:"account_id" json:"accountId"`
		AccountName  string    `db:"account_name" json:"-"`
		CategoryId   int       `db:"category_id" json:"categoryId"`
		CategoryName string    `db:"category_name" json:"-"`
		TxnDate      time.Time `db:"txn_date" json:"txnDate"`
		Amount       float64   `db:"amount" json:"amount"`
		Description  string    `db:"description" json:"description"`
	}
)

func InitStorage() TransactionDataAdapter {
	if config.UseSQL {
		return InitSQL()
	}
	return &TransactionFileData{}
}

const TRANSACTION = "transaction"
