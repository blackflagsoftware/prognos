package accountcolumn

import (
	"sort"

	"github.com/blackflagsoftware/prognos/internal/util"
)

type (
	AccountColumnFileData struct{}
)

func (a *AccountColumnFileData) Read(col *AccountColumn) error {
	accs := []AccountColumn{}
	if err := util.OpenFile(ACCOUNTCOLUMN, &accs); err != nil {
		return err
	}
	if len(accs) == 0 {
		col.Id = 0
		return nil
	}
	for _, colObj := range accs {
		if colObj.Id == col.Id {
			col.AccountId = colObj.AccountId
			col.ColumnName = colObj.ColumnName
			col.Position = colObj.Position
			break
		}
	}
	return nil
}

func (a *AccountColumnFileData) List(col *[]AccountColumn, accountId int) error {
	colAll := &[]AccountColumn{}
	if err := util.OpenFile(ACCOUNTCOLUMN, colAll); err != nil {
		return err
	}
	for _, c := range *colAll {
		if c.AccountId == accountId {
			*col = append(*col, c)
		}
	}
	sort.Slice(*col, func(i, j int) bool {
		return (*col)[i].Position < (*col)[j].Position
	})
	return nil
}

func (a *AccountColumnFileData) Create(col AccountColumn) error {
	accs := []AccountColumn{}
	if err := util.OpenFile(ACCOUNTCOLUMN, &accs); err != nil {
		return err
	}
	maxId := 0
	for _, colObj := range accs {
		if colObj.Id > maxId {
			maxId = colObj.Id
		}
	}
	col.Id = maxId + 1
	accs = append(accs, col)
	return util.SaveFile(ACCOUNTCOLUMN, accs)
}

func (a *AccountColumnFileData) Update(col AccountColumn) error {
	accs := []AccountColumn{}
	if err := util.OpenFile(ACCOUNTCOLUMN, &accs); err != nil {
		return err
	}
	for i, colObj := range accs {
		if colObj.Id == col.Id {
			accs[i] = col
			break
		}
	}
	return util.SaveFile(ACCOUNTCOLUMN, accs)
}

func (a *AccountColumnFileData) Delete(col AccountColumn) error {
	accs := []AccountColumn{}
	if err := util.OpenFile(ACCOUNTCOLUMN, &accs); err != nil {
		return err
	}
	for i, colObj := range accs {
		if colObj.Id == col.Id {
			accs = append(accs[:i], accs[i+1:]...)
			break
		}
	}
	return util.SaveFile(ACCOUNTCOLUMN, accs)
}

func (a *AccountColumnFileData) ColumnIdxByName(accountId int, columnName string) int {
	cols := []AccountColumn{}
	if err := a.List(&cols, accountId); err != nil {
		return -1
	}
	for _, c := range cols {
		if c.ColumnName == columnName {
			return c.Position - 1
		}
	}
	return -1
}
