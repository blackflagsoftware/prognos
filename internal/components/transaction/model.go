package transaction

import "time"

type (
	Transaction struct {
		Id          int       `json:"id"`
		AccountId   int       `json:"accountId"`
		CategoryId  int       `json:"categoryId"`
		TxnDate     time.Time `json:"txnDate"`
		Amount      float64   `json:"amount"`
		Description string    `json:"description"`
	}
)
