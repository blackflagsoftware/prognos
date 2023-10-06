package category

import "fmt"

func Read(cat *Category) error {
	if cat.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataRead(cat)
}

func List(cat *[]Category) error {
	return DataList(cat)
}

func Create(cat Category) error {
	if cat.CategoryName == "" {
		return fmt.Errorf("Invalid CategoryName")
	}
	return DataCreate(cat)
}

func Update(cat Category) error {
	if cat.CategoryName == "" {
		return fmt.Errorf("Invalid CategoryName")
	}
	return DataUpdate(cat)
}

func Delete(cat Category) error {
	if cat.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return DataDelete(cat)
}

// this checks the list of categories for possible duplicate
// if not found, create a new record
// but will have to run through list again to get new id for incoming 'cat'
func CheckAndCreate(cat *Category) error {
	cats := []Category{}
	if err := List(&cats); err != nil {
		return fmt.Errorf("Unable to add new category: %s", err)
	}
	for _, c := range cats {
		if c.CategoryName == cat.CategoryName {
			cat.Id = c.Id
			return nil
		}
	}
	if err := Create(*cat); err != nil {
		return fmt.Errorf("Unable to add new category: %s", err)
	}
	// TODO: not efficient, refactor for video
	if err := List(&cats); err != nil {
		return fmt.Errorf("Unable to add new category: %s", err)
	}
	for _, c := range cats {
		if c.CategoryName == cat.CategoryName {
			cat.Id = c.Id
			return nil
		}
	}
	return nil
}
