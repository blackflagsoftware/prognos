package account

import (
	"fmt"
	"strings"
)

func Read(acc *Account) error {
	if acc.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataRead(acc)
}

func List(acc *[]Account) error {
	return DataList(acc)
}

func Create(acc Account) error {
	if acc.AccountName == "" {
		return fmt.Errorf("Invalid AccountName")
	}
	if acc.OwnerName == "" {
		return fmt.Errorf("Invalid OwnerName")
	}
	if acc.DateFormat == "" {
		return fmt.Errorf("Invalid DateFormat")
	}
	return DataCreate(acc)
}

func Update(acc Account) error {
	// verify the record by Id
	if acc.AccountName == "" {
		return fmt.Errorf("Invalid AccountName")
	}
	if acc.OwnerName == "" {
		return fmt.Errorf("Invalid OwnerName")
	}
	if acc.DateFormat == "" {
		return fmt.Errorf("Invalid DateFormat")
	}
	return DataUpdate(acc)
}

func Delete(acc Account) error {
	if acc.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataDelete(acc)
}

func (a Account) TransformDateFormat() string {
	goFormat := strings.Replace(a.DateFormat, "yyyy", "2006", 1)
	goFormat = strings.Replace(goFormat, "mm", "01", 1)
	goFormat = strings.Replace(goFormat, "dd", "02", 1)
	return goFormat
}
