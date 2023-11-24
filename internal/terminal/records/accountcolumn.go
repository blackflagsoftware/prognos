package records

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	acc "github.com/blackflagsoftware/prognos/internal/entities/account"
	ac "github.com/blackflagsoftware/prognos/internal/entities/accountcolumn"
	"github.com/blackflagsoftware/prognos/internal/util"
)

func AccountColumnMenu() {
	for {
		util.ClearScreen()
		messages := []string{"** AccountColumns **", "Please select your function"}
		prompts := []string{"(c) Create", "(r) Read", "(u) Update", "(d) Delete", "(l) List"}
		acceptablePrompts := []string{"c", "r", "u", "d", "l"}
		exitString := "e"
		selection := util.BasicPrompt(messages, prompts, acceptablePrompts, exitString)

		if selection == "e" {
			break
		}
		switch selection {
		case "c":
			createAccountColumn()
		case "r":
			readAccountColumn()
		case "u":
			updateAccountColumn()
		case "d":
			deleteAccountColumn()
		case "l":
			listAccountColumn()
		}
	}
}

func createAccountColumn() {
	accountcolumn := ac.AccountColumn{}
	for {
		util.ClearScreen()
		fmt.Println("** AccountColumn - Create **")
		fmt.Println("* - Required")
		fmt.Println("")
		PrintAccounts()
		accountcolumn.AccountId = util.ParseInputIntWithMessage("AccountId*: ")
		accountcolumn.ColumnName = util.ParseInputWithMessage("ColumnName [TxnDate | Amount | Description | Category]*: ")
		accountcolumn.Position = util.ParseInputIntWithMessage("Position*: ")
		err := ac.Create(accountcolumn)
		if err != nil {
			fmt.Printf("AccountColumn was not added: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Add another accountcolumn") {
			break
		}
	}
}

func readAccountColumn() {
	accountcolumn := &ac.AccountColumn{}
	for {
		util.ClearScreen()
		getAccountColumn(accountcolumn)
		addlText := ""
		if accountcolumn.Id == 0 {
			addlText = "Record not found"
		}
		fmt.Printf("AccountColumn Details: %s\n", addlText)
		fmt.Println("")
		if accountcolumn.Id != 0 {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "Id\tColumnName\tPosition")
			fmt.Fprintln(writer, "----\t--------\t--------")
			fmt.Fprintf(writer, "%d\t%s\t%d\n", accountcolumn.Id, accountcolumn.ColumnName, accountcolumn.Position)
			writer.Flush()
		}
		fmt.Println("")

		if !util.AskYesOrNo("Read another accountcolumn") {
			break
		}
	}
}

func updateAccountColumn() {
	origAccountColumn := &ac.AccountColumn{}
	newAccountColumn := ac.AccountColumn{}
	for {
		util.ClearScreen()
		fmt.Println("** AccountColumn - Update **")
		fmt.Println("Saved value in [], press enter to keep")
		fmt.Println("* - Required")
		fmt.Println("")
		getAccountColumn(origAccountColumn)
		newAccountColumn.Id = origAccountColumn.Id
		newAccountColumn.AccountId = util.ParseInputIntWithMessageCompare(fmt.Sprintf("AccountId [%d]*: ", origAccountColumn.AccountId), origAccountColumn.AccountId)
		newAccountColumn.ColumnName = util.ParseInputStringWithMessageCompare(fmt.Sprintf("ColumnName [%s] [TxnDate | Amount | Description | Category]*: ", origAccountColumn.ColumnName), origAccountColumn.ColumnName)
		newAccountColumn.Position = util.ParseInputIntWithMessageCompare(fmt.Sprintf("Position [%d]*: ", origAccountColumn.Position), origAccountColumn.Position)
		err := ac.Update(newAccountColumn)
		if err != nil {
			fmt.Printf("AccountColumn was not updated: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Update another accountcolumn") {
			break
		}
	}
}

func deleteAccountColumn() {
	accountcolumn := ac.AccountColumn{}
	for {
		util.ClearScreen()
		accountcolumn.Id = util.ParseInputIntWithMessage("Enter AccountColumn Id to delete: ")
		err := ac.Delete(accountcolumn)
		if err != nil {
			fmt.Printf("AccountColumn was not deleted: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Delete another accountcolumn") {
			break
		}
	}
}

func listAccountColumn() {
	PrintAccounts()
	accountId := util.ParseInputIntWithMessage("Filter by Account: ")
	fmt.Println("")
	accountcolumns := &[]ac.AccountColumn{}
	ac.List(accountcolumns, accountId)
	fmt.Println("AccountColumns - List")
	fmt.Println("")
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "Id\tColumnName\tPosition")
	fmt.Fprintln(writer, "----\t--------\t--------")
	for _, accountcolumn := range *accountcolumns {
		fmt.Fprintf(writer, "%d\t%s\t%d\n", accountcolumn.Id, accountcolumn.ColumnName, accountcolumn.Position)
	}
	writer.Flush()
	fmt.Println("")
	fmt.Print("Press 'enter' to continue ")
	util.ParseInput()
}

func getAccountColumn(accountcolumn *ac.AccountColumn) {
	for {
		accountcolumn.Id = util.ParseInputIntWithMessage("Enter AccountColumn Id: ")
		err := ac.Read(accountcolumn)
		if err != nil {
			fmt.Printf("AccountColumn was not added: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		break
	}
}

func PrintAccounts() {
	accounts := []acc.Account{}
	as := acc.InitStorage()
	am := acc.NewAccountManager(as)
	err := am.List(&accounts)
	if err != nil {
		fmt.Println("Could not read account:", err)
		return
	}
	line := []string{}
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	for i := range accounts {
		if i%4 == 0 {
			fmt.Fprintln(writer, strings.Join(line, "\t"))
			line = []string{}
		}
		line = append(line, fmt.Sprintf("%d - %s", accounts[i].Id, accounts[i].AccountName))
	}
	fmt.Fprintln(writer, strings.Join(line, "\t"))
	writer.Flush()
	fmt.Println("")
}
