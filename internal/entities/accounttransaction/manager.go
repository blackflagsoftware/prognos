package accounttransaction

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	acc "github.com/blackflagsoftware/prognos/internal/entities/account"
	ac "github.com/blackflagsoftware/prognos/internal/entities/accountcolumn"
	tra "github.com/blackflagsoftware/prognos/internal/entities/transaction"
	th "github.com/blackflagsoftware/prognos/internal/entities/transactionhistory"
)

func LoadTransactionFile(account acc.Account, filePath string) error {
	fileName := path.Base(filePath) // just get the file name from full path
	// check if the file has already been loaded
	if DataExists(account.Id, fileName) {
		return fmt.Errorf("Error: unable to process file, has already been processed")
	}
	// load the file from filePath
	skipHeader := account.SkipHeader
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error: unable to read file path: %v", err)
	}
	// go through each line and get the data per column that match the account column for amount, date, description, category (optional), check reversesign
	lineSep := strings.TrimSuffix(account.LineSep, " ")
	if lineSep == "" {
		lineSep = "\n"
	}
	lines := bytes.Split(content, []byte(lineSep))
	for i, line := range lines {
		if i == 0 && skipHeader {
			continue
		}
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		elementSep := account.ElementSep
		if elementSep == "" {
			elementSep = ","
		}
		lineParts := bytes.Split(line, []byte(elementSep))
		// transaction date
		transactionDateStr := ""
		idx := ac.ColumnIdxByName(account.Id, "Transaction Date")
		if idx == -1 {
			fmt.Println("Transaction Date position couldn't be found")
		} else {
			transactionDateStr = string(lineParts[idx])
		}
		dateFormat := account.TransformDateFormat()
		transactionDate, err := time.Parse(dateFormat, transactionDateStr)
		if err != nil {
			fmt.Printf("Transaction date parse was invalid [%s]: %s\n", transactionDateStr, err)
		}
		// description
		description := ""
		idx = ac.ColumnIdxByName(account.Id, "Description")
		if idx == -1 {
			fmt.Println("Description position couldn't be found")
		} else {
			description = string(lineParts[idx])
		}
		// category - if no category is available, set it to the description column position
		categoryId := 0 // set to 'Unknown'
		idx = ac.ColumnIdxByName(account.Id, "Category")
		if idx == -1 {
			fmt.Println("Category position couldn't be found")
		} else {
			categoryId = th.FindCategory(string(lineParts[idx]))
		}
		// amount
		amountStr := "0"
		idx = ac.ColumnIdxByName(account.Id, "Amount")
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
			CategoryId:  categoryId,
		}
		if err := tra.Create(transaction); err != nil {
			fmt.Printf("Transaction: %v, failed to be created: %s\n", transaction, err)
		}
	}
	accountTransaction := AccountTransaction{AccountId: account.Id, FileName: fileName, DateLoaded: time.Now().UTC().Format(time.RFC3339)}
	DataCreate(accountTransaction)
	return nil
}

func LoadUncategorizedTransactions(accountId int) []tra.Transaction {
	transactions := []tra.Transaction{}
	if err := tra.Uncategorized(&transactions, accountId); err != nil {
		fmt.Println("LoadUncategorizedTransactions: error", err)
	}
	return transactions
}
