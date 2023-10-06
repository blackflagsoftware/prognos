package transaction

import "time"

type (
	Transaction struct {
		Id           int       `json:"id"`
		AccountId    int       `json:"accountId"`
		AccountName  string    `json:"-"`
		CategoryId   int       `json:"categoryId"`
		CategoryName string    `json:"-"`
		TxnDate      time.Time `json:"txnDate"`
		Amount       float64   `json:"amount"`
		Description  string    `json:"description"`
	}
)

const TRANSACTION = "transaction"
