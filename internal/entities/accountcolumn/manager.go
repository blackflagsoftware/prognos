package accountcolumn

import "fmt"

type (
	AccountColumnDataAdapter interface {
		Read(*AccountColumn) error
		List(*[]AccountColumn, int) error
		Create(AccountColumn) error
		Update(AccountColumn) error
		Delete(AccountColumn) error
		ColumnIdxByName(int, string) int
	}

	AccountColumnManager struct {
		accountColumnDataAdapter AccountColumnDataAdapter
	}
)

func NewAccountColumnManager(acm AccountColumnDataAdapter) AccountColumnManager {
	return AccountColumnManager{accountColumnDataAdapter: acm}
}

func (a *AccountColumnManager) Read(acc *AccountColumn) error {
	if acc.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return a.accountColumnDataAdapter.Read(acc)
}

func (a *AccountColumnManager) List(acc *[]AccountColumn, accountId int) error {
	return a.accountColumnDataAdapter.List(acc, accountId)
}

func (a *AccountColumnManager) Create(acc AccountColumn) error {
	if acc.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	if acc.ColumnName == "" {
		return fmt.Errorf("Empty ColumnName")
	}
	if acc.ColumnName != "TxnDate" {
		if acc.ColumnName != "Amount" {
			if acc.ColumnName != "Description" {
				if acc.ColumnName != "Category" {
					return fmt.Errorf("Invalid ColumnName, valid options [TxnDate | Amount | Description | Category]")
				}
			}
		}
	}
	if acc.Position < 1 {
		return fmt.Errorf("Invalid PositionId")
	}
	return a.accountColumnDataAdapter.Create(acc)
}

func (a *AccountColumnManager) Update(acc AccountColumn) error {
	if acc.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	if acc.ColumnName == "" {
		return fmt.Errorf("Empty ColumnName")
	}
	if acc.ColumnName != "TxnDate" {
		if acc.ColumnName != "Amount" {
			if acc.ColumnName != "Description" {
				if acc.ColumnName != "Category" {
					return fmt.Errorf("Invalid ColumnName, valid options [TxnDate | Amount | Description | Category]")
				}
			}
		}
	}
	if acc.Position < 1 {
		return fmt.Errorf("Invalid PositionId")
	}
	return a.accountColumnDataAdapter.Update(acc)
}

func (a *AccountColumnManager) Delete(acc AccountColumn) error {
	if acc.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return a.accountColumnDataAdapter.Delete(acc)
}

func (a *AccountColumnManager) ColumnIdxByName(accountId int, columnName string) int {
	return a.accountColumnDataAdapter.ColumnIdxByName(accountId, columnName)
}
