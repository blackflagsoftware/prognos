package terminal

import (
	rec "github.com/blackflagsoftware/prognos/internal/terminal/records"
	tra "github.com/blackflagsoftware/prognos/internal/terminal/transactions"
	"github.com/blackflagsoftware/prognos/internal/util"
)

func MainMenu() {
	for {
		util.ClearScreen()
		messages := []string{"** Main Menu **", "Please choose a function"}
		prompts := []string{"(r) Records", "(t) Transactions"}
		acceptablePrompts := []string{"r", "t"}
		exitString := "e"
		selection := util.BasicPrompt(messages, prompts, acceptablePrompts, exitString)
		if selection == "e" {
			break
		}
		switch selection {
		case "r":
			RecordsMenu()
		case "t":
			TransactionMenu()
		}
	}
}

func RecordsMenu() {
	for {
		util.ClearScreen()
		messages := []string{"** Records **", "Please make a selection"}
		prompts := []string{"(a) Account", "(o) AccountColumn", "(c) Category", "(t) Transaction"}
		acceptablePrompts := []string{"a", "o", "c", "t"}
		exitString := "e"
		selection := util.BasicPrompt(messages, prompts, acceptablePrompts, exitString)

		if selection == "e" {
			break
		}
		switch selection {
		case "t":
			rec.TransactionMenu()
		case "a":
			rec.AccountMenu()
		case "o":
			rec.AccountColumnMenu()
		case "c":
			rec.CategoryMenu()
		}
	}
}

func TransactionMenu() {
	for {
		util.ClearScreen()
		messages := []string{"** Transactions **", "Please make a selection"}
		prompts := []string{"(l) Load"}
		acceptablePrompts := []string{"l"}
		exitString := "e"
		selection := util.BasicPrompt(messages, prompts, acceptablePrompts, exitString)

		if selection == "e" {
			break
		}
		switch selection {
		case "l":
			tra.TransactionsLoad()
		}
	}
}
