package account

import (
	"fmt"
	"strings"
)

type (
	AccountDataAdapter interface {
		Read(*Account) error
		List(*[]Account) error
		Create(Account) error
		Update(Account) error
		Delete(Account) error
	}

	AccountManager struct {
		accountDataAdapter AccountDataAdapter
	}
)

func NewAccountManager(ama AccountDataAdapter) AccountManager {
	return AccountManager{accountDataAdapter: ama}
}

func (a *AccountManager) Read(acc *Account) error {
	if acc.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return a.accountDataAdapter.Read(acc)
}

func (a *AccountManager) List(acc *[]Account) error {
	return a.accountDataAdapter.List(acc)
}

func (a *AccountManager) Create(acc Account) error {
	if acc.AccountName == "" {
		return fmt.Errorf("Invalid AccountName")
	}
	if acc.OwnerName == "" {
		return fmt.Errorf("Invalid OwnerName")
	}
	if acc.DateFormat == "" {
		return fmt.Errorf("Invalid DateFormat")
	}
	if len(acc.LineSep) == 0 {
		acc.LineSep = "\n"
	}
	if len(acc.ElementSep) == 0 {
		acc.ElementSep = ","
	}
	return a.accountDataAdapter.Create(acc)
}

func (a *AccountManager) Update(acc Account) error {
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
	if len(acc.LineSep) == 0 {
		acc.LineSep = "\n"
	}
	if len(acc.ElementSep) == 0 {
		acc.ElementSep = ","
	}
	return a.accountDataAdapter.Update(acc)
}

func (a *AccountManager) Delete(acc Account) error {
	if acc.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return a.accountDataAdapter.Delete(acc)
}

func (a Account) TransformDateFormat() string {
	goFormat := strings.Replace(a.DateFormat, "yyyy", "2006", 1)
	goFormat = strings.Replace(goFormat, "mm", "01", 1)
	goFormat = strings.Replace(goFormat, "dd", "02", 1)
	return goFormat
}
