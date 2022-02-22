package transaction

import "time"

type (
	Transaction struct {
		Id          int
		AccountId   int
		CategoryId  int
		TxnDate     time.Time
		Amount      float64
		Description string
	}
)
