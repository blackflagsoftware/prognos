package accountcolumn

import (
	"github.com/blackflagsoftware/prognos/internal/util"
)

func DataRead(col *AccountColumn) error {
	accs := []AccountColumn{}
	if err := util.OpenFile("AccountColumn", &accs); err != nil {
		return err
	}
	if len(accs) == 0 {
		col.Id = 0
		return nil
	}
	for _, colObj := range accs {
		if colObj.Id == col.Id {
			col.AccountId = colObj.AccountId
			col.ColumnId = colObj.ColumnId
			col.Position = colObj.Position
			break
		}
	}
	return nil
}

// func DataList(col *[]AccountColumn) error {
// 	return util.OpenFile("accountcolumn", col)
// }

func DataCreate(col AccountColumn) error {
	accs := []AccountColumn{}
	if err := util.OpenFile("accountcolumn", &accs); err != nil {
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
	return util.SaveFile("accountcolumn", accs)
}

func DataUpdate(col AccountColumn) error {
	accs := []AccountColumn{}
	if err := util.OpenFile("account", &accs); err != nil {
		return err
	}
	for i, colObj := range accs {
		if colObj.Id == col.Id {
			accs[i] = col
			break
		}
	}
	return util.SaveFile("accountcolumn", accs)
}

func DataDelete(col AccountColumn) error {
	accs := []AccountColumn{}
	if err := util.OpenFile("account", &accs); err != nil {
		return err
	}
	for i, colObj := range accs {
		if colObj.Id == col.Id {
			accs = append(accs[:i], accs[i+1:]...)
			break
		}
	}
	return util.SaveFile("account", accs)
}
