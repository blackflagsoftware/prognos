package budgetallocation

type (
	BudgetAllocationDataAdapter interface {
		List(*[]BudgetAllocation) error
		Upsert(int, float64) error
	}

	BudgetAllocationManager struct {
		budgetAllocationDataAdapter BudgetAllocationDataAdapter
	}
)

func NewBudgetAllocationManager(ba BudgetAllocationDataAdapter) BudgetAllocationManager {
	return BudgetAllocationManager{budgetAllocationDataAdapter: ba}
}

func (b *BudgetAllocationManager) List(budgetAllocation *[]BudgetAllocation) error {
	return b.budgetAllocationDataAdapter.List(budgetAllocation)
}

func (b *BudgetAllocationManager) Upsert(categoryId int, amount float64) error {
	return b.budgetAllocationDataAdapter.Upsert(categoryId, amount)
}
