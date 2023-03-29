package transaction

import (
	"github.com/blackflagsoftware/prognos/internal/util"
)

func DataRead(tra *Transaction) error {
	tras := []Transaction{}
	if err := util.OpenFile(TRANSACTION, &tras); err != nil {
		return err
	}
	if len(tras) == 0 {
		tra.Id = 0
		return nil
	}
	for _, traObj := range tras {
		if traObj.Id == tra.Id {
			tra.AccountId = traObj.AccountId
			tra.CategoryId = traObj.CategoryId
			tra.TxnDate = traObj.TxnDate
			tra.Amount = traObj.Amount
			tra.Description = traObj.Description
			break
		}
	}
	return nil
}

func DataList(tra *[]Transaction) error {
	return util.OpenFile(TRANSACTION, tra)
}

func DataCreate(tra Transaction) error {
	tras := []Transaction{}
	if err := util.OpenFile(TRANSACTION, &tras); err != nil {
		return err
	}
	maxId := 0
	for _, traObj := range tras {
		if traObj.Id > maxId {
			maxId = traObj.Id
		}
	}
	tra.Id = maxId + 1
	tras = append(tras, tra)
	return util.SaveFile(TRANSACTION, tras)
}

func DataUpdate(tra Transaction) error {
	tras := []Transaction{}
	if err := util.OpenFile(TRANSACTION, &tras); err != nil {
		return err
	}
	for i, traObj := range tras {
		if traObj.Id == tra.Id {
			tras[i] = tra
			break
		}
	}
	return util.SaveFile(TRANSACTION, tras)
}

func DataDelete(tra Transaction) error {
	tras := []Transaction{}
	if err := util.OpenFile(TRANSACTION, &tras); err != nil {
		return err
	}
	for i, traObj := range tras {
		if traObj.Id == tra.Id {
			tras = append(tras[:i], tras[i+1:]...)
			break
		}
	}
	return util.SaveFile(TRANSACTION, tras)
}

func DataDeleteAll() error {
	tras := []Transaction{}
	if err := util.OpenFile(TRANSACTION, &tras); err != nil {
		return err
	}
	tras = []Transaction{}
	return util.SaveFile(TRANSACTION, tras)
}
