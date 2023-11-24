package transaction

import (
	"fmt"
	"time"

	stor "github.com/blackflagsoftware/prognos/internal/storage"
	"github.com/jmoiron/sqlx"
)

type (
	TransactionSqlData struct {
		DB *sqlx.DB
	}
)

func InitSQL() *TransactionSqlData {
	db := stor.PostgresInit()
	return &TransactionSqlData{DB: db}
}

func (t *TransactionSqlData) Read(tra *Transaction) error {
	sqlGet := `
		SELECT
			t.id,
			t.account_id,
			CONCAT('[', a.id, '] ', a.account_name) AS account_name,
			t.category_id,
			CONCAT('[', c.id, '] ', c.category_name) AS category_name,
			t.txn_date,
			t.amount,
			t.description
		FROM transaction AS t
		LEFT JOIN account AS a ON t.account_id = a.id
		LEFT JOIN category AS c ON t.category_id = c.id
		WHERE t.id = $1`
	if errDB := t.DB.Get(tra, sqlGet, tra.Id); errDB != nil {
		return fmt.Errorf("Transaction Read: unable to read record, %s", errDB)
	}
	return nil
}

func (t *TransactionSqlData) List(tra *[]Transaction) error {
	sqlSelect := `
		SELECT
			t.id,
			t.account_id,
			CONCAT('[', a.id, '] ', a.account_name) AS account_name,
			t.category_id,
			CONCAT('[', c.id, '] ', c.category_name) AS category_name,
			t.txn_date,
			t.amount,
			t.description
		FROM transaction AS t
		LEFT JOIN account AS a ON t.account_id = a.id
		LEFT JOIN category AS c ON t.category_id = c.id`
	if errDB := t.DB.Select(tra, sqlSelect); errDB != nil {
		return fmt.Errorf("Transaction Select: unable to select records, %s", errDB)
	}
	return nil
}

func (t *TransactionSqlData) Create(tra Transaction) error {
	sqlCreate := `
		INSERT INTO transaction (
			account_id,
			category_id,
			txn_date,
			amount,
			description
		) VALUES (
			:account_id,
			:category_id,
			:txn_date,
			:amount,
			:description
		)`
	if _, errDB := t.DB.NamedExec(sqlCreate, tra); errDB != nil {
		return fmt.Errorf("Transaction Create: unable to create record, %s", errDB)
	}
	return nil
}

func (t *TransactionSqlData) Update(tra Transaction) error {
	sqlUpdate := `
		UPDATE transaction SET 
			account_id = :account_id,
			category_id = :category_id,
			txn_date = :txn_date,
			amount = :amount,
			description = :description
		WHERE id = :id`
	if _, errDB := t.DB.NamedExec(sqlUpdate, tra); errDB != nil {
		return fmt.Errorf("Transaction Update: unable to update record, %s", errDB)
	}
	return nil
}

func (t *TransactionSqlData) Delete(tra Transaction) error {
	sqlDelete := "DELETE FROM transaction WHERE .id = $1"
	if _, errDB := t.DB.Exec(sqlDelete, tra.Id); errDB != nil {
		return fmt.Errorf("Transaction Delete: unable to delete record, %s", errDB)
	}
	return nil
}

func (t *TransactionSqlData) DeleteAll() error {
	sqlDelete := "DELETE FROM transaction"
	if _, errDB := t.DB.Exec(sqlDelete); errDB != nil {
		return fmt.Errorf("Transaction DeleteAll: unable to delete records, %s", errDB)
	}
	return nil
}

func (t *TransactionSqlData) Uncategorized(transactions *[]Transaction, accountId int) error {
	sqlSelect := `
		SELECT
			t.id,
			t.account_id,
			CONCAT('[', a.id, '] ', a.account_name) AS account_name,
			t.category_id,
			t.txn_date,
			t.amount,
			t.description
		FROM transaction AS t
		INNER JOIN account AS a ON t.account_id = a.id
		WHERE t.account_id = $1 AND category_id < 1`
	if errDB := t.DB.Select(transactions, sqlSelect, accountId); errDB != nil {
		return fmt.Errorf("Transaction Uncategorized: unable to select records, %s", errDB)
	}
	return nil
}

func (t *TransactionSqlData) TransactionByDate(transactions *[]Transaction, startDate time.Time, endDate time.Time) error {
	sqlSelect := `
		SELECT
			t.id,
			t.account_id,
			CONCAT('[', a.id, '] ', a.account_name) AS account_name,
			t.category_id,
			CONCAT('[', c.id, '] ', c.category_name) AS category_name,
			t.txn_date,
			t.amount,
			t.description
		FROM transaction AS t
		LEFT JOIN account AS a ON t.account_id = a.id
		LEFT JOIN category AS c ON t.category_id = c.id
		WHERE txn_date >= $1 AND txn_date < $2`
	if errDB := t.DB.Select(transactions, sqlSelect, startDate, endDate); errDB != nil {
		return fmt.Errorf("Transaction TransactionByDate: unable to select records, %s", errDB)
	}
	return nil
}
