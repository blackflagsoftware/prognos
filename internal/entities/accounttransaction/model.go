package accounttransaction

import "github.com/blackflagsoftware/prognos/config"

type (
	AccountTransaction struct {
		AccountId  int    `json:"account_id"`
		FileName   string `json:"file_name"`
		DateLoaded string `json:"date_loaded"`
	}
)

func InitStorage() AccountTransactionDataAdapter {
	if config.UseSQL {
		return InitSQL()
	}
	return &AccountTransactionFileData{}
}

const ACCOUNTTRANSACTION = "accounttransaction"
