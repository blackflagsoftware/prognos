package category

import "fmt"

type (
	CategoryDataAdapter interface {
		Read(*Category) error
		List(*[]Category) error
		Create(Category) error
		Update(Category) error
		Delete(Category) error
	}

	CategoryManager struct {
		categoryDataAdapter CategoryDataAdapter
	}
)

func NewCategoryManager(ca CategoryDataAdapter) CategoryManager {
	return CategoryManager{categoryDataAdapter: ca}
}

func (c *CategoryManager) Read(cat *Category) error {
	if cat.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return c.categoryDataAdapter.Read(cat)
}

func (c *CategoryManager) List(cat *[]Category) error {
	return c.categoryDataAdapter.List(cat)
}

func (c *CategoryManager) Create(cat Category) error {
	if cat.CategoryName == "" {
		return fmt.Errorf("Invalid CategoryName")
	}
	return c.categoryDataAdapter.Create(cat)
}

func (c *CategoryManager) Update(cat Category) error {
	if cat.CategoryName == "" {
		return fmt.Errorf("Invalid CategoryName")
	}
	return c.categoryDataAdapter.Update(cat)
}

func (c *CategoryManager) Delete(cat Category) error {
	if cat.Id < 1 {
		return fmt.Errorf("Invalid Id")
	}
	return c.categoryDataAdapter.Delete(cat)
}

// this checks the list of categories for possible duplicate
// if not found, create a new record
// but will have to run through list again to get new id for incoming 'cat'
func (c *CategoryManager) CheckAndCreate(cat *Category) error {
	cats := []Category{}
	if err := c.List(&cats); err != nil {
		return fmt.Errorf("Unable to add new category: %s", err)
	}
	for _, c := range cats {
		if c.CategoryName == cat.CategoryName {
			cat.Id = c.Id
			return nil
		}
	}
	if err := c.Create(*cat); err != nil {
		return fmt.Errorf("Unable to add new category: %s", err)
	}
	// TODO: not efficient, refactor for video, do it here and above
	if err := c.List(&cats); err != nil {
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
