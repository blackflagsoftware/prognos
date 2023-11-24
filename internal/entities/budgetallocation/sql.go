package budgetallocation

import (
	"fmt"

	stor "github.com/blackflagsoftware/prognos/internal/storage"
	"github.com/jmoiron/sqlx"
)

type (
	BudgetAllocationSqlData struct {
		DB *sqlx.DB
	}
)

func InitSQL() *BudgetAllocationSqlData {
	db := stor.PostgresInit()
	return &BudgetAllocationSqlData{DB: db}
}

func (b *BudgetAllocationSqlData) List(budgetAllocation *[]BudgetAllocation) error {
	sqlSelect := `
		SELECT
			category_id,
			amount
		FROM budget_allocation`
	if errDB := b.DB.Select(budgetAllocation, sqlSelect); errDB != nil {
		return fmt.Errorf("BudgetAllocation List: unable to select records, %s", errDB)
	}
	return nil
}

func (b *BudgetAllocationSqlData) Upsert(categoryId int, amount float64) error {
	sqlExists := "SELECT EXISTS(SELECT category_id FROM budget_allocation WHERE category_id = $1)"
	sqlInsert := "INSERT INTO budget_allocation (category_id, amount) VALUES (:category_id, :amount)"
	sqlUpdate := "UPDATE budget_allocation SET amount = :amount WHERE category_id = :category_id"
	sqlQuery := ""
	var exists bool
	if errDB := b.DB.Get(&exists, sqlExists, categoryId); errDB != nil {
		return fmt.Errorf("BudgetAllocation Upsert: unable to get exists, %s", errDB)
	}
	sqlQuery = sqlInsert
	if exists {
		sqlQuery = sqlUpdate
	}
	ba := BudgetAllocation{CategoryId: categoryId, Amount: amount}
	if _, errDB := b.DB.NamedExec(sqlQuery, ba); errDB != nil {
		return fmt.Errorf("BudgetAllocation Upsert: unable to insert/update record, %s", errDB)
	}
	return nil
}
