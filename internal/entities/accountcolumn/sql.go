package accountcolumn

import (
	"fmt"

	stor "github.com/blackflagsoftware/prognos/internal/storage"
	"github.com/jmoiron/sqlx"
)

type (
	AccountColumnSqlData struct {
		DB *sqlx.DB
	}
)

func InitSQL() *AccountColumnSqlData {
	db := stor.PostgresInit()
	return &AccountColumnSqlData{DB: db}
}

func (a *AccountColumnSqlData) Read(col *AccountColumn) error {
	sqlGet := `
		SELECT
			id,
			account_id,
			position,
			column_name
		FROM account_column WHERE id = $1`
	if errDB := a.DB.Get(col, sqlGet, col.Id); errDB != nil {
		return fmt.Errorf("AccountColumn Read: unable to get record, %s", errDB)
	}
	return nil
}

func (a *AccountColumnSqlData) List(col *[]AccountColumn, accountId int) error {
	sqlSelect := `
		SELECT
			id,
			account_id,
			position,
			column_name
		FROM account_column WHERE account_id = $1
		ORDER BY position`
	if errDB := a.DB.Select(col, sqlSelect, accountId); errDB != nil {
		return fmt.Errorf("AccountColumn List: unable to get records, %s", errDB)
	}
	return nil
}

func (a *AccountColumnSqlData) Create(col AccountColumn) error {
	sqlCreate := `
		INSERT INTO account_column (
			account_id,
			position,
			column_name
		) VALUES (
			:account_id,
			:position,
			:column_name
		)`
	if _, errDB := a.DB.NamedExec(sqlCreate, col); errDB != nil {
		return fmt.Errorf("AccountColumn Create: unable to create record, %s", errDB)
	}
	return nil
}

func (a *AccountColumnSqlData) Update(col AccountColumn) error {
	sqlUpdate := `
		UPDATE account_column SET 
			account_id = :account_id,
			position = :position,
			column_nam = :column_name
		WHERE id = :id`
	if _, errDB := a.DB.NamedExec(sqlUpdate, col); errDB != nil {
		return fmt.Errorf("AccountColumn Update: unable to update record, %s", errDB)
	}
	return nil
}

func (a *AccountColumnSqlData) Delete(col AccountColumn) error {
	sqlDelete := "DELETE FROM account_column WHERE id = $1"
	if _, errDB := a.DB.Exec(sqlDelete, col.Id); errDB != nil {
		return fmt.Errorf("AccountColumn Delete: unable to delete record, %s", errDB)
	}
	return nil
}

func (a *AccountColumnSqlData) ColumnIdxByName(accountId int, columnName string) int {
	var id int
	sqlGet := `
		SELECT
			position - 1 AS id
		FROM account_column WHERE account_id = $1 AND column_name = $2`
	if errDB := a.DB.Get(&id, sqlGet, accountId, columnName); errDB != nil {
		fmt.Printf("AccountColumn ColumnIdxByName: unable to get record, %s", errDB)
		return -1
	}
	return id
}
