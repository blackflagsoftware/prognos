package account

import (
	"github.com/blackflagsoftware/prognos/internal/util"
)

func DataRead(acc *Account) error {
	accs := []Account{}
	if err := util.OpenFile(ACCOUNT, &accs); err != nil {
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
			acc.SkipHeader = accObj.SkipHeader
			acc.LineSep = accObj.LineSep
			acc.ElementSep = accObj.ElementSep
			break
		}
	}
	return nil
}

func DataList(acc *[]Account) error {
	return util.OpenFile(ACCOUNT, acc)
}

func DataCreate(acc Account) error {
	accs := []Account{}
	if err := util.OpenFile(ACCOUNT, &accs); err != nil {
		return err
	}
	maxId := 0
	for _, accObj := range accs {
		if accObj.Id > maxId {
			maxId = accObj.Id
		}
	}
	if len(acc.LineSep) == 0 {
		acc.LineSep = "\n"
	}
	if len(acc.ElementSep) == 0 {
		acc.ElementSep = ","
	}
	acc.Id = maxId + 1
	accs = append(accs, acc)
	return util.SaveFile(ACCOUNT, accs)
}

func DataUpdate(acc Account) error {
	accs := []Account{}
	if err := util.OpenFile(ACCOUNT, &accs); err != nil {
		return err
	}
	for i, accObj := range accs {
		if accObj.Id == acc.Id {
			if len(acc.LineSep) == 0 {
				acc.LineSep = "\n"
			}
			if len(acc.ElementSep) == 0 {
				acc.ElementSep = ","
			}
			accs[i] = acc
			break
		}
	}
	return util.SaveFile(ACCOUNT, accs)
}

func DataDelete(acc Account) error {
	accs := []Account{}
	if err := util.OpenFile(ACCOUNT, &accs); err != nil {
		return err
	}
	for i, accObj := range accs {
		if accObj.Id == acc.Id {
			accs = append(accs[:i], accs[i+1:]...)
			break
		}
	}
	return util.SaveFile(ACCOUNT, accs)
}
