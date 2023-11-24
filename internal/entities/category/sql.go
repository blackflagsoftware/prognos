package category

import (
	"fmt"

	stor "github.com/blackflagsoftware/prognos/internal/storage"
	"github.com/jmoiron/sqlx"
)

type (
	CategorySqlData struct {
		DB *sqlx.DB
	}
)

func InitSQL() *CategorySqlData {
	db := stor.PostgresInit()
	return &CategorySqlData{DB: db}
}

func (c *CategorySqlData) Read(cat *Category) error {
	sqlGet := "SELECT id, category_name FROM category WHERE id = $1"
	if errDB := c.DB.Get(cat, sqlGet, cat.Id); errDB != nil {
		return fmt.Errorf("Category Read: unable to get record, %s", errDB)
	}
	return nil
}

func (c *CategorySqlData) List(cat *[]Category) error {
	sqlSelect := "SELECT id, category_name FROM category"
	if errDB := c.DB.Select(cat, sqlSelect); errDB != nil {
		return fmt.Errorf("Category List: unable to select records, %s", errDB)
	}
	return nil
}

func (c *CategorySqlData) Create(cat Category) error {
	sqlCreate := "INSERT INTO category (category_name) VALUES (:category_name)"
	if _, errDB := c.DB.NamedExec(sqlCreate, cat); errDB != nil {
		return fmt.Errorf("Category Create: unable to create record, %s", errDB)
	}
	return nil
}

func (c *CategorySqlData) Update(cat Category) error {
	sqlUpdate := "UPDATE category SET category_name = :category_name WHERE id = :id"
	if _, errDB := c.DB.NamedExec(sqlUpdate, cat); errDB != nil {
		return fmt.Errorf("Category Update: unable to update record, %s", errDB)
	}
	return nil
}

func (c *CategorySqlData) Delete(cat Category) error {
	sqlDelete := "DELETE category WHERE id = $1"
	if _, errDB := c.DB.Exec(sqlDelete, cat.Id); errDB != nil {
		return fmt.Errorf("Category Delete: unable to delete record, %s", errDB)
	}
	return nil
}
