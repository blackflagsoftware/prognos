package records

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	cat "github.com/blackflagsoftware/prognos/internal/entities/category"
	"github.com/blackflagsoftware/prognos/internal/util"
)

func CategoryMenu() {
	for {
		util.ClearScreen()
		messages := []string{"** Categorys **", "Please select your function"}
		prompts := []string{"(c) Create", "(r) Read", "(u) Update", "(d) Delete", "(l) List"}
		acceptablePrompts := []string{"c", "r", "u", "d", "l"}
		exitString := "e"
		selection := util.BasicPrompt(messages, prompts, acceptablePrompts, exitString)

		if selection == "e" {
			break
		}
		switch selection {
		case "c":
			createCategory()
		case "r":
			readCategory()
		case "u":
			updateCategory()
		case "d":
			deleteCategory()
		case "l":
			listCategory()
		}
	}
}

func createCategory() {
	category := cat.Category{}
	for {
		util.ClearScreen()
		fmt.Println("** Category - Create **")
		fmt.Println("* - Required")
		fmt.Println("")
		category.CategoryName = util.ParseInputWithMessage("Category Name*: ")
		err := cat.Create(category)
		if err != nil {
			fmt.Printf("Category was not added: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Add another category") {
			break
		}
	}
}

func readCategory() {
	category := &cat.Category{}
	for {
		util.ClearScreen()
		getCategory(category)
		addlText := ""
		if category.Id == 0 {
			addlText = "Record not found"
		}
		fmt.Printf("Category Details: %s\n", addlText)
		fmt.Println("")
		if category.Id != 0 {
			writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
			fmt.Fprintln(writer, "Id\tCategoryName")
			fmt.Fprintln(writer, "----\t------------")
			fmt.Fprintf(writer, "%d\t%s\n", category.Id, category.CategoryName)
			writer.Flush()
		}
		fmt.Println("")
		if !util.AskYesOrNo("Read another category") {
			break
		}
	}
}

func updateCategory() {
	origCategory := &cat.Category{}
	newCategory := cat.Category{}
	for {
		util.ClearScreen()
		fmt.Println("** Category - Update **")
		fmt.Println("Saved value in [], press enter to keep")
		fmt.Println("* - Required")
		fmt.Println("")
		getCategory(origCategory)
		newCategory.Id = origCategory.Id
		newCategory.CategoryName = util.ParseInputStringWithMessageCompare(fmt.Sprintf("Category Name [%s]*: ", origCategory.CategoryName), origCategory.CategoryName)
		err := cat.Update(newCategory)
		if err != nil {
			fmt.Printf("Category was not updated: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Update another category") {
			break
		}
	}
}

func deleteCategory() {
	category := cat.Category{}
	for {
		util.ClearScreen()
		category.Id = util.ParseInputIntWithMessage("Enter Category Id to delete: ")
		err := cat.Delete(category)
		if err != nil {
			fmt.Printf("Category was not deleted: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		if !util.AskYesOrNo("Delete another category") {
			break
		}
	}
}

func listCategory() {
	categorys := &[]cat.Category{}
	cat.List(categorys)
	fmt.Println("Categorys - List")
	fmt.Println("")
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "Id\tCategoryName")
	fmt.Fprintln(writer, "----\t------------")
	for _, category := range *categorys {
		fmt.Fprintf(writer, "%d\t%s\n", category.Id, category.CategoryName)
	}
	writer.Flush()
	fmt.Println("")
	fmt.Print("Press 'enter' to continue ")
	util.ParseInput()
}

func PrintCategories(categories []cat.Category) {
	line := []string{}
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	for i := range categories {
		if i%4 == 0 {
			fmt.Fprintln(writer, strings.Join(line, "\t"))
			line = []string{}
		}
		line = append(line, fmt.Sprintf("%d - %s", categories[i].Id, categories[i].CategoryName))
	}
	fmt.Fprintln(writer, strings.Join(line, "\t"))
	writer.Flush()
	fmt.Println("")
}

func getCategory(category *cat.Category) {
	for {
		category.Id = util.ParseInputIntWithMessage("Enter Category Id: ")
		err := cat.Read(category)
		if err != nil {
			fmt.Printf("Category was not added: %s\n", err)
			fmt.Print("Press 'enter' to continue")
			util.ParseInput()
			continue
		}
		break
	}
}
