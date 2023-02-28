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
