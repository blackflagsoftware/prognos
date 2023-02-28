package category

import (
	"github.com/blackflagsoftware/prognos/internal/util"
)

func DataRead(cat *Category) error {
	cats := []Category{}
	if err := util.OpenFile("category", &cats); err != nil {
		return err
	}
	if len(cats) == 0 {
		cat.Id = 0
		return nil
	}
	for _, accObj := range cats {
		if accObj.Id == cat.Id {
			cat.CategoryName = accObj.CategoryName
			break
		}
	}
	return nil
}

func DataList(cat *[]Category) error {
	return util.OpenFile("category", cat)
}

func DataCreate(cat Category) error {
	cats := []Category{}
	if err := util.OpenFile("category", &cats); err != nil {
		return err
	}
	maxId := 0
	for _, accObj := range cats {
		if accObj.Id > maxId {
			maxId = accObj.Id
		}
	}
	cat.Id = maxId + 1
	cats = append(cats, cat)
	return util.SaveFile("category", cats)
}

func DataUpdate(cat Category) error {
	cats := []Category{}
	if err := util.OpenFile("category", &cats); err != nil {
		return err
	}
	for i, accObj := range cats {
		if accObj.Id == cat.Id {
			cats[i] = cat
			break
		}
	}
	return util.SaveFile("category", cats)
}

func DataDelete(cat Category) error {
	cats := []Category{}
	if err := util.OpenFile("category", &cats); err != nil {
		return err
	}
	for i, accObj := range cats {
		if accObj.Id == cat.Id {
			cats = append(cats[:i], cats[i+1:]...)
			break
		}
	}
	return util.SaveFile("category", cats)
}
