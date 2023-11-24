package records

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	a "github.com/blackflagsoftware/prognos/internal/entities/account"
	"github.com/blackflagsoftware/prognos/internal/util"
)

var (
	accountManager a.AccountManager
)

func init() {
	as := a.InitStorage()
	accountManager = a.NewAccountManager(as)
}

func AccountMenu() {
	for {
		util.ClearScreen()
		messages := []string{"** Accounts **", "Please select your function"}
		prompts := []string{"(c) Create", "(r) Read", "(u) Update", "(d) Delete", "(l) List"}
		acceptablePrompts := []string{"c", "r", "u", "d", "l"}
		exitString := "e"
		selection := util.BasicPrompt(messages, prompts, acceptablePrompts, exitString)

		if selection == "e" {
			break
		}
		switch selection {
		case "c":
			createAccount()
		case "r":
			readAccount()
		case "u":
			updateAccount()
		case "d":
			deleteAccount()
		case "l":
			listAccount()
		}
	}
}

func createAccount() {
	account := a.Account{}
	for {
		util.ClearScreen()
		fmt.Println("** Account - Create **")
		fmt.Println("* - Required")
		fmt.Println("")
		account.AccountName = util.ParseInputWithMessage("Account Name*: ")
		account.OwnerName = util.ParseInputWithMessage("Owner Name*: ")
		account.DateFormat = util.ParseInputWithMessage("Date Format*: ")
		account.ReverseSign = util.ParseInputBoolWithMessage("Reverse Sign*: ")
		account.SkipHeader = util.ParseInputBoolWithMessage("Skip Header: ")
		account.LineSep = util.ParseInputWithMessage("Line Seprator (default '\\n'): ")
		account.ElementSep = util.ParseInputWithMessage("Element Seprator (default ','): ")
		err := accountManager.Create(account)
		if err != nil {
			fmt.Printf("Account was not added: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Add another account") {
			break
		}
	}
}

func readAccount() {
	account := &a.Account{}
	for {
		util.ClearScreen()
		getAccount(account)
		addlText := ""
		if account.Id == 0 {
			addlText = "Record not found"
		}
		fmt.Printf("Account Details: %s\n", addlText)
		fmt.Println("")
		if account.Id != 0 {
			lineSep := strings.TrimSuffix(account.LineSep, " ")
			if lineSep == "\n" {
				lineSep = "new-line"
			}
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "Id\tAccountName\tOwnerName\tDateFormat\tReverseSign\tSkipHeader\tLineSeparator\tElementSeparator")
			fmt.Fprintln(writer, "----\t---------\t--------\t--------\t--------\t----------\t--------------\t----------------")
			fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%t\t%t\t%s\t%s\n", account.Id, account.AccountName, account.OwnerName, account.DateFormat, account.ReverseSign, account.SkipHeader, lineSep, account.ElementSep)
			writer.Flush()
		}
		fmt.Println("")
		if !util.AskYesOrNo("Read another account") {
			break
		}
	}
}

func updateAccount() {
	origAccount := &a.Account{}
	newAccount := a.Account{}
	for {
		util.ClearScreen()
		fmt.Println("** Account - Update **")
		fmt.Println("Saved value in [], press enter to keep")
		fmt.Println("* - Required")
		fmt.Println("")
		getAccount(origAccount)
		newAccount.Id = origAccount.Id
		newAccount.AccountName = util.ParseInputStringWithMessageCompare(fmt.Sprintf("Account Name [%s]*: ", origAccount.AccountName), origAccount.AccountName)
		newAccount.OwnerName = util.ParseInputStringWithMessageCompare(fmt.Sprintf("Owner Name [%s]*: ", origAccount.OwnerName), origAccount.OwnerName)
		newAccount.DateFormat = util.ParseInputStringWithMessageCompare(fmt.Sprintf("Date Format [%s]*: ", origAccount.DateFormat), origAccount.DateFormat)
		newAccount.ReverseSign = util.ParseInputBoolWithMessageCompare(fmt.Sprintf("Reverse Sign[%t]*: ", origAccount.ReverseSign), origAccount.ReverseSign)
		newAccount.SkipHeader = util.ParseInputBoolWithMessageCompare(fmt.Sprintf("Skip Header[%t]*: ", origAccount.SkipHeader), origAccount.SkipHeader)
		newAccount.LineSep = util.ParseInputStringWithMessageCompare(fmt.Sprintf("Line Separator[%s]: ", origAccount.LineSep), string(origAccount.LineSep))
		newAccount.ElementSep = util.ParseInputStringWithMessageCompare(fmt.Sprintf("Element Separator[%s]: ", origAccount.ElementSep), string(origAccount.ElementSep))
		err := accountManager.Update(newAccount)
		if err != nil {
			fmt.Printf("Account was not updated: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Update another account") {
			break
		}
	}
}

func deleteAccount() {
	account := a.Account{}
	for {
		util.ClearScreen()
		account.Id = util.ParseInputIntWithMessage("Enter Account Id to delete: ")
		err := accountManager.Delete(account)
		if err != nil {
			fmt.Printf("Account was not deleted: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Delete another account") {
			break
		}
	}
}

func listAccount() {
	accounts := &[]a.Account{}
	accountManager.List(accounts)
	fmt.Println("Accounts - List")
	fmt.Println("")
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "Id\tAccountName\tOwnerName\tDateFormat\tReverseSign\tSkipHeader\tLineSeparator\tElementSeparator")
	fmt.Fprintln(writer, "----\t---------\t--------\t--------\t--------\t----------\t--------------\t----------------")
	for _, account := range *accounts {
		lineSep := strings.TrimSuffix(account.LineSep, " ")
		if lineSep == "\n" {
			lineSep = "new-line"
		}
		fmt.Fprintf(writer, "%d\t%s\t%s\t%s\t%t\t%t\t%s\t%s\n", account.Id, account.AccountName, account.OwnerName, account.DateFormat, account.ReverseSign, account.SkipHeader, lineSep, account.ElementSep)
	}
	writer.Flush()
	fmt.Println("")
	fmt.Print("Press 'enter' to continue ")
	util.ParseInput()
}

func getAccount(account *a.Account) {
	for {
		account.Id = util.ParseInputIntWithMessage("Enter Account Id: ")
		err := accountManager.Read(account)
		if err != nil {
			fmt.Printf("Account was not added: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		break
	}
}
