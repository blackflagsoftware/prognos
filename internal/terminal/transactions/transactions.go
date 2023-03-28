package transactions

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/blackflagsoftware/prognos/config"
	acc "github.com/blackflagsoftware/prognos/internal/entities/account"
	at "github.com/blackflagsoftware/prognos/internal/entities/accounttransaction"
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
			if err := acc.DataRead(&account); err != nil {
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
			if err := at.LoadTransactionFile(account, loadFile); err != nil {
				fmt.Println(err)
			}
			fmt.Println("File loaded, press 'enter' to continue")
			util.ParseInput()
			break
		}
	}
}
