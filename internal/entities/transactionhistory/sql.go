package transactionhistory

import (
	"fmt"

	stor "github.com/blackflagsoftware/prognos/internal/storage"
	"github.com/blackflagsoftware/prognos/internal/util"
	"github.com/jmoiron/sqlx"
)

type (
	TransactionHistorySqlData struct {
		DB *sqlx.DB
	}
)

func InitSQL() *TransactionHistorySqlData {
	db := stor.PostgresInit()
	return &TransactionHistorySqlData{DB: db}
}

func (t *TransactionHistorySqlData) Read(transactionHistory map[string]int) error {
	if err := util.OpenFile("transactionhistory", &transactionHistory); err != nil {
		fmt.Println("Error transactionhistory:", err)
		return err
	}
	return nil
}

func (t *TransactionHistorySqlData) Create(text string, categoryId int) error {
	transactionHistory := make(map[string]int)
	if err := t.Read(transactionHistory); err != nil {
		return err
	}
	transactionHistory[text] = categoryId
	return util.SaveFile("transactionhistory", transactionHistory)
}
