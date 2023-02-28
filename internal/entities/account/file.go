package account

import (
	"github.com/blackflagsoftware/prognos/internal/util"
)

func DataRead(acc *Account) error {
	accs := []Account{}
	if err := util.OpenFile("account", &accs); err != nil {
		return err
	}
	if len(accs) == 0 {
		acc.Id = 0
		return nil
	}
	for _, accObj := range accs {
		if accObj.Id == acc.Id {
			acc.AccountName = accObj.AccountName
			acc.OwnerName = accObj.OwnerName
			acc.DateFormat = accObj.DateFormat
			acc.ReverseSign = accObj.ReverseSign
			break
		}
	}
	return nil
}

func DataList(acc *[]Account) error {
	return util.OpenFile("account", acc)
}

func DataCreate(acc Account) error {
	accs := []Account{}
	if err := util.OpenFile("account", &accs); err != nil {
		return err
	}
	maxId := 0
	for _, accObj := range accs {
		if accObj.Id > maxId {
			maxId = accObj.Id
		}
	}
	acc.Id = maxId + 1
	accs = append(accs, acc)
	return util.SaveFile("account", accs)
}

func DataUpdate(acc Account) error {
	accs := []Account{}
	if err := util.OpenFile("account", &accs); err != nil {
		return err
	}
	for i, accObj := range accs {
		if accObj.Id == acc.Id {
			accs[i] = acc
			break
		}
	}
	return util.SaveFile("account", accs)
}

func DataDelete(acc Account) error {
	accs := []Account{}
	if err := util.OpenFile("account", &accs); err != nil {
		return err
	}
	for i, accObj := range accs {
		if accObj.Id == acc.Id {
			accs = append(accs[:i], accs[i+1:]...)
			break
		}
	}
	return util.SaveFile("account", accs)
}
