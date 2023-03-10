package records

import (
	"fmt"
	"time"

	tra "github.com/blackflagsoftware/prognos/internal/entities/transaction"
	"github.com/blackflagsoftware/prognos/internal/util"
)

func TransactionMenu() {
	for {
		util.ClearScreen()
		messages := []string{"** Transactions **", "Please select your function"}
		prompts := []string{"(c) Create", "(r) Read", "(u) Update", "(d) Delete", "(l) List"}
		acceptablePrompts := []string{"c", "r", "u", "d", "l", "da"}
		exitString := "e"
		selection := util.BasicPrompt(messages, prompts, acceptablePrompts, exitString)

		if selection == "e" {
			break
		}
		switch selection {
		case "c":
			createTransaction()
		case "r":
			readTransaction()
		case "u":
			updateTransaction()
		case "d":
			deleteTransaction()
		case "l":
			listTransaction()
		case "da":
			deleteAll()
		}
	}
}

func createTransaction() {
	transaction := tra.Transaction{}
	for {
		util.ClearScreen()
		fmt.Println("** Transaction - Create **")
		fmt.Println("* - Required")
		fmt.Println("")
		transaction.AccountId = util.ParseInputIntWithMessage("Account ID*: ")
		transaction.CategoryId = util.ParseInputIntWithMessage("Category ID*: ")
		for {
			fmt.Print("Transaction Date (mm-dd-yyyy)*: ")
			txnDateStr := util.ParseInput()
			var err error
			transaction.TxnDate, err = time.Parse("01-02-2006", txnDateStr)
			if err == nil {
				break
			}
			fmt.Println("Not a valid date (mm-dd-yyyy)")
		}
		transaction.Amount = util.ParseInputFloatWithMessage("Amount*: ")
		fmt.Print("Description: ")
		transaction.Description = util.ParseInput()
		err := tra.Create(transaction)
		if err != nil {
			fmt.Printf("Transaction was not added: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Add another transaction") {
			break
		}
	}
}

func readTransaction() {
	transaction := &tra.Transaction{}
	for {
		util.ClearScreen()
		getTransaction(transaction)
		addlText := ""
		if transaction.Id == 0 {
			addlText = "Record not found"
		}
		fmt.Printf("Transaction Details: %s\n", addlText)
		if transaction.Id != 0 {
			printTransactionDetail(*transaction)
		}

		if !util.AskYesOrNo("Read another transaction") {
			break
		}
	}
}

func updateTransaction() {
	origTransaction := &tra.Transaction{}
	newTransaction := tra.Transaction{}
	for {
		util.ClearScreen()
		fmt.Println("** Transaction - Update **")
		fmt.Println("Saved value in [], press enter to keep")
		fmt.Println("* - Required")
		fmt.Println("")
		getTransaction(origTransaction)
		newTransaction.Id = origTransaction.Id
		newTransaction.AccountId = util.ParseInputIntWithMessageCompare(fmt.Sprintf("Account ID [%d]*: ", origTransaction.AccountId), origTransaction.AccountId)
		newTransaction.CategoryId = util.ParseInputIntWithMessageCompare(fmt.Sprintf("Category ID [%d]*: ", origTransaction.CategoryId), origTransaction.CategoryId)
		for {
			fmt.Printf("Transaction Date (mm-dd-yyyy) [%s]*: ", origTransaction.TxnDate.Format("01-02-2006"))
			txnDateStr := util.ParseInput()
			if txnDateStr == "" {
				newTransaction.TxnDate = origTransaction.TxnDate
				break
			}
			var err error
			newTransaction.TxnDate, err = time.Parse("01-02-2006", txnDateStr)
			if err == nil {
				break
			}
			fmt.Println("Not a valid date (mm-dd-yyyy)")
		}
		newTransaction.Amount = util.ParseInputFloatWithMessageCompare(fmt.Sprintf("Amount [%0.2f]*: ", origTransaction.Amount), origTransaction.Amount)
		fmt.Printf("Description [%s]: ", origTransaction.Description)
		newTransaction.Description = origTransaction.Description
		description := util.ParseInput()
		if description != "" {
			newTransaction.Description = description
		}
		err := tra.Update(newTransaction)
		if err != nil {
			fmt.Printf("Transaction was not updated: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Update another transaction") {
			break
		}
	}
}

func deleteTransaction() {
	transaction := tra.Transaction{}
	for {
		util.ClearScreen()
		transaction.Id = util.ParseInputIntWithMessage("Enter Transaction Id to delete: ")
		err := tra.Delete(transaction)
		if err != nil {
			fmt.Printf("Transaction was not deleted: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Delete another transaction") {
			break
		}
	}
}

func listTransaction() {
	transactions := &[]tra.Transaction{}
	tra.List(transactions)
	fmt.Println("Transactions - List")
	for _, transaction := range *transactions {
		printTransactionDetail(transaction)
	}
	fmt.Println("")
	fmt.Print("Press 'enter' to continue ")
	util.ParseInput()
}

func deleteAll() {
	if util.AskYesOrNo("Delete All?") {
		tra.DeleteAll()
	}
}

func getTransaction(transaction *tra.Transaction) {
	for {
		transaction.Id = util.ParseInputIntWithMessage("Enter Transaction Id: ")
		err := tra.Read(transaction)
		if err != nil {
			fmt.Printf("Transaction was not added: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		break
	}
}

func printTransactionDetail(transaction tra.Transaction) {
	fmt.Printf("Id: %d\n", transaction.Id)
	fmt.Printf("Account Id: %d\n", transaction.AccountId)
	fmt.Printf("Category Id: %d\n", transaction.CategoryId)
	fmt.Printf("Transaction Date: %s\n", transaction.TxnDate.Format("01-02-2006"))
	fmt.Printf("Amount: %0.2f\n", transaction.Amount)
	fmt.Printf("Description: %s\n", transaction.Description)
	fmt.Println("")
}
