package accounttransaction

type (
	AccountTransaction struct {
		AccountId  int    `json:"account_id"`
		FileName   string `json:"file_name"`
		DateLoaded string `json:"date_loaded"`
	}
)

const ACCOUNTTRANSACTION = "accounttransaction"
