package transactions

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/blackflagsoftware/prognos/config"
	acc "github.com/blackflagsoftware/prognos/internal/entities/account"
	at "github.com/blackflagsoftware/prognos/internal/entities/accounttransaction"
	cat "github.com/blackflagsoftware/prognos/internal/entities/category"
	tra "github.com/blackflagsoftware/prognos/internal/entities/transaction"
	th "github.com/blackflagsoftware/prognos/internal/entities/transactionhistory"
	rec "github.com/blackflagsoftware/prognos/internal/terminal/records"
	"github.com/blackflagsoftware/prognos/internal/util"
)

func TransactionsLoad() {
	for {
		util.ClearScreen()
		loadFile := util.ParseInputWithMessage("Please enter file name to load in 'project_path/data/test_data' (e to exit): ")
		if loadFile == "" {
			fmt.Println("Empty file name, press 'enter' to try again")
			util.ParseInput()
			continue
		}
		if strings.ToLower(loadFile) == "e" {
			break
		}
		loadFile = path.Clean(fmt.Sprintf("%s/../test_data/%s", config.FilePath, loadFile))
		if _, err := os.Stat(loadFile); os.IsNotExist(err) {
			fmt.Printf("File not found in: %s, press 'enter' to try again\n", loadFile)
			util.ParseInput()
			continue
		}
		for {
			util.ClearScreen()
			rec.PrintAccounts()
			accountIdStr := util.ParseInputWithMessage("Which account to load these under: ")
			// catch anything other than a int
			accountId, err := strconv.ParseInt(accountIdStr, 10, 64)
			if err != nil {
				fmt.Println("Invalid selection, press 'enter' to try again")
				util.ParseInput()
				continue
			}
			account := acc.Account{Id: int(accountId)}
			as := acc.InitStorage()
			am := acc.NewAccountManager(as)
			if err := am.Read(&account); err != nil {
				fmt.Println("Invalid selection, press 'enter' to try again")
				util.ParseInput()
				continue
			}
			// not a valid account
			if account.Id == 0 {
				fmt.Println("Invalid selection, press 'enter' to try again")
				util.ParseInput()
				continue
			}
			ats := at.InitStorage()
			atm := at.NewAccountTransactionManager(ats)
			if err := atm.LoadTransactionFile(account, loadFile); err != nil {
				fmt.Println(err)
			}
			fmt.Println("File loaded, press 'enter' to continue")
			util.ParseInput()
			break
		}
	}
}

func TransactionsUncategorized() {
	ats := at.InitStorage()
	atm := at.NewAccountTransactionManager(ats)
	cs := cat.InitStorage()
	cm := cat.NewCategoryManager(cs)
	ts := tra.InitStorage()
	tm := tra.NewTransactionManager(ts)
	ths := th.InitStorage()
	thm := th.NewTransactionHistoryManager(ths)
	for {
		util.ClearScreen()
		rec.PrintAccounts()
		accountId := util.ParseInputIntWithMessage("Filter by Account (0 - exit): ")
		if accountId == 0 {
			break
		}
		// load categories
		categories := []cat.Category{}
		if err := cm.List(&categories); err != nil {
			fmt.Println("TransactionsUncategorized: failed to load categories, error:", err)
			break
		}
		// load transactions
		transactions := atm.LoadUncategorizedTransactions(accountId)
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	newTransaction:
		for i, t := range transactions {
			util.ClearScreen()
			fmt.Fprintln(writer, "Txn Date\tAmount\tDescription")
			fmt.Fprintln(writer, "--------\t------\t-----------")
			fmt.Fprintf(writer, "%s\t%0.2f\t%s\n", t.TxnDate, t.Amount, t.Description)
			writer.Flush()
			for {
				rec.PrintCategories(categories)
				fmt.Println("(s) Skip")
				fmt.Println("(n) New Category")
				selection := util.ParseInputWithMessage("Category Id to assign to this transaction: ")
				if strings.ToLower(selection) == "s" {
					continue newTransaction
				}
				if strings.ToLower(selection) == "n" {
					// TODO: makes the loop too long; refactor for video
					for {
						newCategory := util.ParseInputWithMessage("Enter in the new category (e - exit): ")
						if strings.ToLower(newCategory) == "e" {
							break
						}
						c := cat.Category{CategoryName: newCategory}
						if err := cm.CheckAndCreate(&c); err != nil {
							fmt.Println(err)
							yesNo := util.AskYesOrNo("Try again (y/n): ")
							if yesNo {
								continue
							}
						}
						// apply the new category to our existing list
						categories = append(categories, c)
						// TODO: refactor this; for video
						transactions[i].CategoryId = c.Id
						if err := tm.Update(transactions[i]); err != nil {
							fmt.Println("TransactionsUncategorized: updating transaction:", err)
						}
						if err := thm.Create(transactions[i].Description, c.Id); err != nil {
							fmt.Println("TransactionsUncategorized: creating transaction history:", err)
						}
						break
					}
					continue newTransaction
				}
				selectionInt, err := strconv.Atoi(selection)
				if err != nil {
					fmt.Println("Invalid selection, press 'enter' to try again")
					util.ParseInput()
					continue
				}
				transactions[i].CategoryId = selectionInt
				if err := tm.Update(transactions[i]); err != nil {
					fmt.Println("TransactionsUncategorized: updating transaction:", err)
				}
				if err := thm.Create(transactions[i].Description, selectionInt); err != nil {
					fmt.Println("TransactionsUncategorized: creating transaction history:", err)
				}
				break
			}
		}

	}
}
