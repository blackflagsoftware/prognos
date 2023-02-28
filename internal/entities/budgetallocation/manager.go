package budgetallocation

func List(budgetAllocation *[]BudgetAllocation) error {
	return DataList(budgetAllocation)
}

func Upsert(categoryId int, amount float64) error {
	return DataUpsert(categoryId, amount)
}
