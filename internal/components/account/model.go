package account

type (
	Account struct {
		Id          int64
		AccountName string
		OwnerName   string
		DateFormat  string
		ReverseSign bool
	}
)
