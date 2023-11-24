package transactionhistory

import "github.com/blackflagsoftware/prognos/config"

func InitStorage() TransactionHistoryDataAdapter {
	if config.UseSQL {
		return InitSQL()
	}
	return &TransactionHistoryFileData{}
}
