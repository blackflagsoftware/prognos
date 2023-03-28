package accounttransaction

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	acc "github.com/blackflagsoftware/prognos/internal/entities/account"
	ac "github.com/blackflagsoftware/prognos/internal/entities/accountcolumn"
	tra "github.com/blackflagsoftware/prognos/internal/entities/transaction"
)

func LoadTransactionFile(account acc.Account, filePath string) error {
	// TODO: check if the file has already been loaded, give error if has been loaded
	// load account columns by accountId
	accountColumns := []ac.AccountColumn{}
	if err := ac.DataList(&accountColumns, account.Id); err != nil {
		return fmt.Errorf("Error: unable to get account columns: %v", err)
	}
	// load the file from filePath
	skipHeader := true // TODO: get this from account setting, might be something we set
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error: unable to read file path: %v", err)
	}
	// go through each line and get the data per column that match the account column for amount, date, description, category (optional), check reversesign
	lines := bytes.Split(content, []byte("\n")) // TODO: account settting of line separator
	for i, line := range lines {
		if i == 0 && skipHeader {
			continue
		}
		if len(line) == 0 { // TODO: what if the lines has spaces?
			continue
		}
		// TODO: need account setting of column separator
		lineParts := bytes.Split(line, []byte(","))
		// transaction date
		transactionDateStr := ""
		idx := columnIdxByName(accountColumns, "Transaction Date")
		if idx == -1 {
			fmt.Println("Transaction Date position couldn't be found")
		} else {
			transactionDateStr = string(lineParts[idx])
		}
		dateFormat := accountDateFormat(account)
		transactionDate, err := time.Parse(dateFormat, transactionDateStr)
		if err != nil {
			fmt.Printf("Transaction date parse was invalid [%s]: %s\n", transactionDateStr, err)
		}
		// description
		description := ""
		idx = columnIdxByName(accountColumns, "Description")
		if idx == -1 {
			fmt.Println("Description position couldn't be found")
		} else {
			description = string(lineParts[idx])
		}
		// amount
		amountStr := "0"
		idx = columnIdxByName(accountColumns, "Amount")
		if idx == -1 {
			fmt.Println("Amount position couldn't be found")
		} else {
			amountStr = string(lineParts[idx])
		}
		amount, err := strconv.ParseFloat(amountStr, 64)
		if account.ReverseSign {
			// spending transactions have an amount coming in as "negative", need to make it positive
			// this will make the payments "negative" but that is ok, ignoring payments for now
			amount *= -1
		}
		if err != nil {
			fmt.Printf("Amount [%s] could not be parsed: %s\n", amountStr, err)
		}

		// save off transaction
		transaction := tra.Transaction{
			AccountId:   account.Id,
			TxnDate:     transactionDate,
			Description: description,
			Amount:      amount,
		}
		if err := tra.Create(transaction); err != nil {
			fmt.Printf("Transaction: %v, failed to be created: %s\n", transaction, err)
		}
	}
	return nil
}

// TODO: move these two func(s) into respective manager calls, where would these go?
func columnIdxByName(accountColumn []ac.AccountColumn, columnName string) int {
	for _, c := range accountColumn {
		if c.ColumnName == columnName {
			return c.Position - 1
		}
	}
	return -1
}

func accountDateFormat(account acc.Account) string {
	goFormat := strings.Replace(account.DateFormat, "yyyy", "2006", 1)
	goFormat = strings.Replace(goFormat, "mm", "01", 1)
	goFormat = strings.Replace(goFormat, "dd", "02", 1)
	return goFormat
}
