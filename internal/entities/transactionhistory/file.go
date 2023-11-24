package transactionhistory

import (
	"fmt"

	"github.com/blackflagsoftware/prognos/internal/util"
)

type (
	TransactionHistoryFileData struct{}
)

func (t *TransactionHistoryFileData) Read(transactionHistory map[string]int) error {
	if err := util.OpenFile("transactionhistory", &transactionHistory); err != nil {
		fmt.Println("Error transactionhistory:", err)
		return err
	}
	return nil
}

func (t *TransactionHistoryFileData) Create(text string, categoryId int) error {
	transactionHistory := make(map[string]int)
	if err := t.Read(transactionHistory); err != nil {
		return err
	}
	transactionHistory[text] = categoryId
	return util.SaveFile("transactionhistory", transactionHistory)
}
