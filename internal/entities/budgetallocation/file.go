package budgetallocation

import (
	"github.com/blackflagsoftware/prognos/internal/util"
)

type (
	BudgetAllocationFileData struct{}
)

func (b *BudgetAllocationFileData) List(budgetAllocation *[]BudgetAllocation) error {
	return util.OpenFile(BUDGETALLOCATION, budgetAllocation)
}

func (b *BudgetAllocationFileData) Upsert(categoryId int, amount float64) error {
	budgetAllocations := []BudgetAllocation{}
	if err := util.OpenFile(BUDGETALLOCATION, &budgetAllocations); err != nil {
		return err
	}
	update := false
	for i := range budgetAllocations {
		if budgetAllocations[i].CategoryId == categoryId {
			budgetAllocations[i].Amount = amount
			update = true
			break
		}
	}
	if !update {
		budgetAllocations = append(budgetAllocations, BudgetAllocation{CategoryId: categoryId, Amount: amount})
	}
	return util.SaveFile(BUDGETALLOCATION, budgetAllocations)
}
