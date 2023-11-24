package accounttransaction

import (
	"github.com/blackflagsoftware/prognos/internal/util"
)

type (
	AccountTransactionFileData struct{}
)

func (a *AccountTransactionFileData) Exists(accountId int, fileName string) bool {
	acc := &[]AccountTransaction{}
	util.OpenFile(ACCOUNTTRANSACTION, acc)
	for _, a := range *acc {
		if a.AccountId == accountId && a.FileName == fileName {
			return true
		}
	}
	return false
}

func (a *AccountTransactionFileData) Create(acc AccountTransaction) error {
	accs := []AccountTransaction{}
	if err := util.OpenFile(ACCOUNTTRANSACTION, &accs); err != nil {
		return err
	}
	accs = append(accs, acc)
	return util.SaveFile(ACCOUNTTRANSACTION, accs)
}
