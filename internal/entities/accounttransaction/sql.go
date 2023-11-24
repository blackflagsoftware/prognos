package accounttransaction

import (
	"fmt"

	stor "github.com/blackflagsoftware/prognos/internal/storage"
	"github.com/jmoiron/sqlx"
)

type (
	AccountTransactionSqlData struct {
		DB *sqlx.DB
	}
)

func InitSQL() *AccountTransactionSqlData {
	db := stor.PostgresInit()
	return &AccountTransactionSqlData{DB: db}
}

func (a *AccountTransactionSqlData) Exists(accountId int, fileName string) bool {
	sqlExists := "SELECT EXISTS(SELECT account_id FROM account_transaction WHERE account_id = $1 AND file_name = $2)"
	var exists bool
	if err := a.DB.Get(&exists, sqlExists, accountId, fileName); err != nil {
		fmt.Printf("AccountTransaction Exists: unable to get record, %s", err)
	}
	return exists
}

func (a *AccountTransactionSqlData) Create(acc AccountTransaction) error {
	sqlCreate := `
		INSERT INTO account_transaction (
			account_id,
			file_name,
			date_loaded
		) VALUES (
			:account_id,
			:file_name,
			:date_loaded
		)`
	if _, err := a.DB.NamedExec(sqlCreate, acc); err != nil {
		return fmt.Errorf("AccountTransaction Create: unable to create record, %s", err)
	}
	return nil
}
