package transaction

import "fmt"

func Read(tra *Transaction) error {
	if tra.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataRead(tra)
}

func List(tra *[]Transaction) error {
	return DataList(tra)
}

func Create(tra Transaction) error {
	if tra.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	// if tra.CategoryId < 1 {
	// 	return fmt.Errorf("Invalid CategoryId")
	// }
	return DataCreate(tra)
}

func Update(tra Transaction) error {
	if tra.AccountId < 1 {
		return fmt.Errorf("Invalid AccountId")
	}
	if tra.CategoryId < 1 {
		return fmt.Errorf("Invalid CategoryId")
	}
	return DataUpdate(tra)
}

func Delete(tra Transaction) error {
	if tra.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataDelete(tra)
}

func DeleteAll() error {
	return DataDeleteAll()
}
