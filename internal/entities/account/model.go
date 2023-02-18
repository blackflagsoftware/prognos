package account

type (
	Account struct {
		Id          int
		AccountName string
		OwnerName   string
		DateFormat  string
		ReverseSign bool
	}
)
