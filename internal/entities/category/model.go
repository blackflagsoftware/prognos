package category

import "github.com/blackflagsoftware/prognos/config"

type (
	Category struct {
		Id           int    `db:"id"`
		CategoryName string `db:"category_name"`
	}
)

func InitStorage() CategoryDataAdapter {
	if config.UseSQL {
		return InitSQL()
	}
	return &CategoryFileData{}
}

const CATEGORY = "category"
