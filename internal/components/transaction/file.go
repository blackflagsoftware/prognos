package transaction

import (
	"github.com/blackflagsoftware/prognos/config"
	"github.com/blackflagsoftware/prognos/internal/util"
)

type (
	FileTransaction struct{}
)

func (d *FileTransaction) Read(tra *Transaction) error {
	tras := []Transaction{}
	if err := util.OpenFile(config.FilePath, &tras); err != nil {
		return err
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

func (d *FileTransaction) List(tra *[]Transaction) error {
	return util.OpenFile(config.FilePath, tra)
}

func (d *FileTransaction) Create(tra Transaction) error {
	tras := []Transaction{}
	if err := util.OpenFile(config.FilePath, &tras); err != nil {
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
	return util.SaveFile(config.FilePath, tras)
}

func (d *FileTransaction) Update(tra Transaction) error {
	tras := []Transaction{}
	if err := util.OpenFile(config.FilePath, &tras); err != nil {
		return err
	}
	for i, traObj := range tras {
		if traObj.Id == tra.Id {
			tras[i] = tra
			break
		}
	}
	return util.SaveFile(config.FilePath, tras)
}

func (d *FileTransaction) Delete(tra Transaction) error {
	tras := []Transaction{}
	if err := util.OpenFile(config.FilePath, &tras); err != nil {
		return err
	}
	for i, traObj := range tras {
		if traObj.Id == tra.Id {
			tras = append(tras[:i], tras[i+1:]...)
			break
		}
	}
	return util.SaveFile(config.FilePath, tras)
}
