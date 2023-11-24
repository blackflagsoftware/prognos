package budgetallocation

import "github.com/blackflagsoftware/prognos/config"

type (
	BudgetAllocation struct {
		CategoryId int     `db:"category_id"`
		Amount     float64 `db:"amount"`
	}
)

func InitStorage() BudgetAllocationDataAdapter {
	if config.UseSQL {
		return InitSQL()
	}
	return &BudgetAllocationFileData{}
}

const BUDGETALLOCATION = "budgetallocation"
