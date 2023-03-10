package records

import (
	"fmt"
	"strconv"
	"strings"

	a "github.com/blackflagsoftware/prognos/internal/entities/account"
	"github.com/blackflagsoftware/prognos/internal/util"
)

func AccountMenu() {
	for {
		fmt.Println("** Account Record Menu **")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("(c) Create")
		fmt.Println("(r) Read")
		fmt.Println("(u) Update")
		fmt.Println("(d) Delete")
		fmt.Println("(l) List")
		fmt.Println("(e) Exit")
		fmt.Println("")
		fmt.Print("Your choice: ")
		selection := util.ParseInput()
		if strings.ToLower(selection) == "e" {
			break
		}
		switch strings.ToLower(selection) {
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
		default:
			fmt.Println("Not a valid entry, press 'enter' to continue")
			util.ParseInput()
		}
	}
}

func createAccount() {
	for {
		account := a.Account{}
		fmt.Println("** Account Create **")
		fmt.Println("")
		fmt.Print("Account Name: ")
		account.AccountName = util.ParseInput()
		fmt.Print("Owner Name: ")
		account.OwnerName = util.ParseInput()
		fmt.Print("Date Format: ")
		account.DateFormat = util.ParseInput()
		fmt.Print("Reverse Sign: ")
		revSign := strings.ToLower(util.ParseInput())
		if revSign == "true" || revSign == "t" || revSign == "y" || revSign == "yes" {
			account.ReverseSign = true
		}
		if err := a.DataCreate(account); err != nil {
			fmt.Println("Unable to create Account:", err)
		}
		fmt.Print("Add another? ")
		selection := strings.ToLower(util.ParseInput())
		if selection == "n" || selection == "no" {
			break
		}
	}
}

func readAccount() {
	for {
		fmt.Println("** Account Read **")
		account := a.Account{}
		fmt.Print("Enter Account Id: ")
		selection := util.ParseInput()
		selectionInt, err := strconv.Atoi(selection)
		if err != nil {
			fmt.Println("Not a number")
			continue
		}
		account.Id = selectionInt
		if err := a.DataRead(&account); err != nil {
			fmt.Println("Unable to record:", err)
			continue
		}
		addlText := ""
		if account.Id == 0 {
			addlText = "Record not found"
		}
		fmt.Printf("Account Details: %s\n", addlText)
		fmt.Println("")
		if account.Id != 0 {
			fmt.Printf("Id: ")
			fmt.Println(account.Id)
			fmt.Printf("Account Name: ")
			fmt.Println(account.AccountName)
			fmt.Printf("Owner Name: ")
			fmt.Println(account.OwnerName)
			fmt.Printf("Date Format: ")
			fmt.Println(account.DateFormat)
			fmt.Printf("Reverse Sign: ")
			fmt.Println(account.ReverseSign)
		}
		fmt.Println("")
		fmt.Print("Read another? ")
		selection = strings.ToLower(util.ParseInput())
		if selection == "n" || selection == "no" {
			break
		}
	}
}

func updateAccount() {
	for {
		fmt.Println("** Account Update **")
		account := a.Account{}
		fmt.Print("Enter Account Id: ")
		selection := util.ParseInput()
		selectionInt, err := strconv.Atoi(selection)
		if err != nil {
			fmt.Println("Not a number")
			continue
		}
		account.Id = selectionInt
		if err := a.DataRead(&account); err != nil {
			fmt.Println("Unable to record:", err)
			continue
		}
		fmt.Println("")
		fmt.Printf("Account Name [%s]: ", account.AccountName)
		accountName := util.ParseInput()
		if accountName != "" && accountName != account.AccountName {
			account.AccountName = accountName
		}
		fmt.Printf("Owner Name [%s]: ", account.OwnerName)
		ownerName := util.ParseInput()
		if ownerName != "" && ownerName != account.OwnerName {
			account.OwnerName = ownerName
		}
		fmt.Printf("Date Format [%s]: ", account.DateFormat)
		dateFormat := util.ParseInput()
		if dateFormat != "" && dateFormat != account.DateFormat {
			account.DateFormat = dateFormat
		}
		fmt.Printf("Reverse Sign [%t]: ", account.ReverseSign)
		revSign := strings.ToLower(util.ParseInput())
		var reverseSign bool
		if revSign != "" {
			if revSign == "true" || revSign == "y" || revSign == "yes" {
				reverseSign = true
			}
			if reverseSign != account.ReverseSign {
				account.ReverseSign = reverseSign
			}
		}
		if err := a.DataUpdate(account); err != nil {
			fmt.Println("Unable to update Account:", err)
		}
		fmt.Print("Update another? ")
		selection = strings.ToLower(util.ParseInput())
		if selection == "n" || selection == "no" {
			break
		}
	}
}

func deleteAccount() {
	for {
		fmt.Println("** Account Delete **")
		account := a.Account{}
		fmt.Print("Enter Account Id: ")
		selection := util.ParseInput()
		selectionInt, err := strconv.Atoi(selection)
		if err != nil {
			fmt.Println("Not a number")
			continue
		}
		account.Id = selectionInt
		if err := a.DataDelete(account); err != nil {
			fmt.Println("Unable to delete account:", err)
			continue
		}
		fmt.Print("Delete another? ")
		selection = strings.ToLower(util.ParseInput())
		if selection == "n" || selection == "no" {
			break
		}
	}
}

func listAccount() {
	fmt.Println("** Account List **")
	accounts := []a.Account{}
	if err := a.DataList(&accounts); err != nil {
		fmt.Println("Unable to get records:", err)
		return
	}
	for _, account := range accounts {
		fmt.Println("")
		if account.Id != 0 {
			fmt.Printf("Id: ")
			fmt.Println(account.Id)
			fmt.Printf("Account Name: ")
			fmt.Println(account.AccountName)
			fmt.Printf("Owner Name: ")
			fmt.Println(account.OwnerName)
			fmt.Printf("Date Format: ")
			fmt.Println(account.DateFormat)
			fmt.Printf("Reverse Sign: ")
			fmt.Println(account.ReverseSign)
		}
	}
	fmt.Println("")
	fmt.Print("Press 'enter to continue")
	util.ParseInput()
}
