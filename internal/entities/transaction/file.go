package transaction

import (
	"fmt"

	acc "github.com/blackflagsoftware/prognos/internal/entities/account"
	cat "github.com/blackflagsoftware/prognos/internal/entities/category"
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
			accountName := accountIdToName(traObj.AccountId)
			tra.AccountName = fmt.Sprintf("[%d] %s", traObj.AccountId, accountName)
			tra.CategoryId = traObj.CategoryId
			tra.TxnDate = traObj.TxnDate
			tra.Amount = traObj.Amount
			tra.Description = traObj.Description
			categoryName := categoryIdToName(traObj.CategoryId)
			// format with name(id)
			tra.CategoryName = fmt.Sprintf("[%d] %s", traObj.CategoryId, categoryName)
			break
		}
	}
	return nil
}

func DataList(tra *[]Transaction) error {
	err := util.OpenFile(TRANSACTION, tra)
	if err != nil {
		return err
	}
	// go through each transaction record and get category
	for i := range *tra {
		categoryName := categoryIdToName((*tra)[i].CategoryId)
		(*tra)[i].CategoryName = fmt.Sprintf("[%d] %s", (*tra)[i].CategoryId, categoryName)
		accountName := accountIdToName((*tra)[i].AccountId)
		(*tra)[i].AccountName = fmt.Sprintf("[%d] %s", (*tra)[i].AccountId, accountName)
	}
	return nil
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

func DataUncategorized(transactions *[]Transaction, accountId int) error {
	tras := []Transaction{}
	if err := util.OpenFile("transaction", &tras); err != nil {
		return err
	}
	for _, tra := range tras {
		if tra.AccountId == accountId && tra.CategoryId < 1 {
			// less than 1 is uncategorized
			*transactions = append(*transactions, tra)
		}
	}
	return nil
}

func categoryIdToName(catId int) string {
	c := cat.Category{Id: catId}
	if err := cat.DataRead(&c); err != nil {
		return ""
	}
	return c.CategoryName
}

func accountIdToName(accId int) string {
	a := acc.Account{Id: accId}
	if err := acc.DataRead(&a); err != nil {
		return ""
	}
	return a.AccountName
}
