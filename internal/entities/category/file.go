package category

import (
	"github.com/blackflagsoftware/prognos/internal/util"
)

type (
	CategoryFileData struct{}
)

func (c *CategoryFileData) Read(cat *Category) error {
	cats := []Category{}
	if err := util.OpenFile(CATEGORY, &cats); err != nil {
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

func (c *CategoryFileData) List(cat *[]Category) error {
	return util.OpenFile(CATEGORY, cat)
}

func (c *CategoryFileData) Create(cat Category) error {
	cats := []Category{}
	if err := util.OpenFile(CATEGORY, &cats); err != nil {
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
	return util.SaveFile(CATEGORY, cats)
}

func (c *CategoryFileData) Update(cat Category) error {
	cats := []Category{}
	if err := util.OpenFile(CATEGORY, &cats); err != nil {
		return err
	}
	for i, accObj := range cats {
		if accObj.Id == cat.Id {
			cats[i] = cat
			break
		}
	}
	return util.SaveFile(CATEGORY, cats)
}

func (c *CategoryFileData) Delete(cat Category) error {
	cats := []Category{}
	if err := util.OpenFile(CATEGORY, &cats); err != nil {
		return err
	}
	for i, accObj := range cats {
		if accObj.Id == cat.Id {
			cats = append(cats[:i], cats[i+1:]...)
			break
		}
	}
	return util.SaveFile(CATEGORY, cats)
}
