package terminal

import (
	"fmt"
	"strings"

	rec "github.com/blackflagsoftware/prognos/internal/terminal/records"
	"github.com/blackflagsoftware/prognos/internal/util"
)

func MainMenu() {
	// TODO:
	// show a menu of options
	// ask for input until the user is done
	// options will go "down" into folders
	// those will have their own menus
	for {
		fmt.Println("** Main Menu **")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("(r) Records")
		fmt.Println("(e) Exit")
		fmt.Println("")
		fmt.Print("Your choice: ")
		selection := util.ParseInput()
		if strings.ToLower(selection) == "e" {
			break
		}
		switch strings.ToLower(selection) {
		case "r":
			RecordsMenu()
		default:
			fmt.Println("Not a valid entry, press 'enter' to continue")
			util.ParseInput()
		}
	}
}

func RecordsMenu() {
	for {
		fmt.Println("** Records **")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("(a) Account")
		fmt.Println("(e) Exit")
		fmt.Println("")
		fmt.Print("Your choice: ")
		selection := util.ParseInput()
		if strings.ToLower(selection) == "e" {
			break
		}
		switch strings.ToLower(selection) {
		case "a":
			rec.AccountMenu()
		default:
			fmt.Println("Not a valid entry, press 'enter' to continue")
			util.ParseInput()
		}
	}
}
