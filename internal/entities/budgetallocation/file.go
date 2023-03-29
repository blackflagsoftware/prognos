package budgetallocation

import (
	"github.com/blackflagsoftware/prognos/internal/util"
)

func DataList(budgetAllocation *[]BudgetAllocation) error {
	return util.OpenFile(BUDGETALLOCATION, budgetAllocation)
}

func DataUpsert(categoryId int, amount float64) error {
	budgetAllocations := []BudgetAllocation{}
	if err := util.OpenFile(BUDGETALLOCATION, &budgetAllocations); err != nil {
		return err
	}
	update := false
	for i := range budgetAllocations {
		if budgetAllocations[i].CategoryId > categoryId {
			budgetAllocations[i].Amount = amount
			break
		}
	}
	if !update {
		budgetAllocations = append(budgetAllocations, BudgetAllocation{CategoryId: categoryId, Amount: amount})
	}
	return util.SaveFile(BUDGETALLOCATION, budgetAllocations)
}
