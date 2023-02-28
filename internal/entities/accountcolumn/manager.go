package accountcolumn

import "fmt"

func Read(acc *AccountColumn) error {
	if acc.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataRead(acc)
}

// func List(acc *[]AccountColumn) error {
// 	return DataList(acc)
// }

func Create(acc AccountColumn) error {
	if acc.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	if acc.ColumnId < 1 {
		return fmt.Errorf("Invalid ColumnId")
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
	if acc.ColumnId < 1 {
		return fmt.Errorf("Invalid ColumnId")
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
