package accountcolumn

import "github.com/blackflagsoftware/prognos/config"

type (
	AccountColumn struct {
		Id         int    `db:"id"`
		AccountId  int    `db:"account_id"`
		Position   int    `db:"position"`
		ColumnName string `db:"column_name"`
	}
)

func InitStorage() AccountColumnDataAdapter {
	if config.UseSQL {
		return InitSQL()
	}
	return &AccountColumnFileData{}
}

const ACCOUNTCOLUMN = "accountColumn"
