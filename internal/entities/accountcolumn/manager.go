package accountcolumn

import "fmt"

func Read(acc *AccountColumn) error {
	if acc.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataRead(acc)
}

func List(acc *[]AccountColumn, accountId int) error {
	return DataList(acc, accountId)
}

func Create(acc AccountColumn) error {
	if acc.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	if acc.ColumnName == "" {
		return fmt.Errorf("Empty ColumnName")
	}
	if acc.ColumnName != "TxnDate" {
		if acc.ColumnName != "Amount" {
			if acc.ColumnName != "Description" {
				return fmt.Errorf("Invalid ColumnName, valid options [TxnDate | Amount | Description]")
			}
		}
	}
	if acc.Position < 1 {
		return fmt.Errorf("Invalid PositionId")
	}
	return DataCreate(acc)
}

func Update(acc AccountColumn) error {
	if acc.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	if acc.ColumnName == "" {
		return fmt.Errorf("Empty ColumnName")
	}
	if acc.ColumnName != "TxnDate" {
		if acc.ColumnName != "Amount" {
			if acc.ColumnName != "Description" {
				return fmt.Errorf("Invalid ColumnName, valid options [TxnDate | Amount | Description]")
			}
		}
	}
	if acc.Position < 1 {
		return fmt.Errorf("Invalid PositionId")
	}
	return DataUpdate(acc)
}

func Delete(acc AccountColumn) error {
	if acc.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataDelete(acc)
}

func ColumnIdxByName(accountId int, columnName string) int {
	return DataColumnIdxByName(accountId, columnName)
}
