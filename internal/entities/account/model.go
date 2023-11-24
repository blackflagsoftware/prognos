package account

import (
	"github.com/blackflagsoftware/prognos/config"
)

type (
	Account struct {
		Id          int    `db:"id"`
		AccountName string `db:"account_name"`
		OwnerName   string `db:"owner_name"`
		DateFormat  string `db:"date_format"`
		ReverseSign bool   `db:"reverse_sign"`
		SkipHeader  bool   `db:"skip_header"`
		LineSep     string `db:"line_sep"`
		ElementSep  string `db:"element_sep"`
	}
)

const ACCOUNT = "account"

func InitStorage() AccountDataAdapter {
	if config.UseSQL {
		return InitSQL()
	}
	return &AccountFileData{}
}
