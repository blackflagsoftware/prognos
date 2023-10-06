package transactionhistory

import (
	"fmt"

	"github.com/blackflagsoftware/prognos/internal/util"
)

func DataRead(transactionHistory map[string]int) error {
	if err := util.OpenFile("transactionhistory", &transactionHistory); err != nil {
		fmt.Println("Error transactionhistory:", err)
		return err
	}
	return nil
}

func DataCreate(text string, categoryId int) error {
	transactionHistory := make(map[string]int)
	if err := Read(transactionHistory); err != nil {
		return err
	}
	transactionHistory[text] = categoryId
	return util.SaveFile("transactionhistory", transactionHistory)
}
