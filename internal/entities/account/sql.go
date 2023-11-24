package account

import (
	"fmt"

	stor "github.com/blackflagsoftware/prognos/internal/storage"
	"github.com/jmoiron/sqlx"
)

type (
	AccountSqlData struct {
		DB *sqlx.DB
	}
)

func InitSQL() *AccountSqlData {
	db := stor.PostgresInit()
	return &AccountSqlData{DB: db}
}

func (a *AccountSqlData) Read(acc *Account) error {
	sqlGet := `
		SELECT
			id,
			account_name,
			owner_name,
			date_format,
			reverse_sign,
			skip_header,
			line_sep,
			element_sep
		FROM account WHERE id = $1`
	if errDB := a.DB.Get(acc, sqlGet, acc.Id); errDB != nil {
		return fmt.Errorf("Account Get: unable to get record, %s", errDB)
	}
	return nil
}

func (a *AccountSqlData) List(acc *[]Account) error {
	sqlSelect := `
		SELECT
			id,
			account_name,
			owner_name,
			date_format,
			reverse_sign,
			skip_header,
			line_sep,
			element_sep
		FROM account`
	if errDB := a.DB.Select(acc, sqlSelect); errDB != nil {
		return fmt.Errorf("Account List: unable to select records, %s", errDB)
	}
	return nil
}

func (a *AccountSqlData) Create(acc Account) error {
	sqlCreate := `
		INSERT INTO account (
			account_name,
			owner_name,
			date_format,
			reverse_sign,
			skip_header,
			line_sep,
			element_sep
		) VALUES (
			:account_name,
			:owner_name,
			:date_format,
			:reverse_sign,
			:skip_header,
			:line_sep,
			:element_sep
		)`
	if _, errDB := a.DB.NamedExec(sqlCreate, acc); errDB != nil {
		return fmt.Errorf("Account Create: unable to create record, %s", errDB)
	}
	return nil
}

func (a *AccountSqlData) Update(acc Account) error {
	sqlUpdate := `
		UPDATE account SET
			account_name = :account_name, 
			owner_name = :owner_name,
			date_format = :date_format,
			reverse_sign = :reverse_sign,
			skip_header = :skip_header,
			line_sep = :line_sep,
			element_sep = :element_sep
		WHERE id = :id`
	if _, errDB := a.DB.NamedExec(sqlUpdate, acc); errDB != nil {
		return fmt.Errorf("Account Update: unable to update record, %s", errDB)
	}
	return nil
}

func (a *AccountSqlData) Delete(acc Account) error {
	sqlDelete := "DELETE FROM account WHERE id = $1"
	if _, errDB := a.DB.Exec(sqlDelete, acc.Id); errDB != nil {
		return fmt.Errorf("Account Delete: unable to delete record, %s", errDB)
	}
	return nil
}
