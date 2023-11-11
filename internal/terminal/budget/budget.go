package budget

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	bud "github.com/blackflagsoftware/prognos/internal/entities/budgetallocation"
	cat "github.com/blackflagsoftware/prognos/internal/entities/category"
	tra "github.com/blackflagsoftware/prognos/internal/entities/transaction"
	"github.com/blackflagsoftware/prognos/internal/util"
)

func ReportMenu() {
	for {
		util.ClearScreen()
		messages := []string{"** Report **", "Please make a selection"}
		prompts := []string{"(l) Last Month", "(c) Custom"}
		acceptablePrompts := []string{"l", "c"}
		exitString := "e"
		selection := util.BasicPrompt(messages, prompts, acceptablePrompts, exitString)

		if selection == "e" {
			break
		}
		switch selection {
		case "l":
			LastMonth()
		case "c":
			CustomMonth()
		}
	}
}

func LastMonth() {
	startDate, endDate := util.GetLastMonth(time.Now()) // TODO: maybe take an extra param to be inclusive or not
	endDate = endDate.AddDate(0, 0, 1)                  // make it exclusive, get the whole month
	transactions := []tra.Transaction{}
	if err := tra.TransactionByDate(&transactions, startDate, endDate); err != nil {
		fmt.Println("Unable to get transactions by date", err)
		return
	}
	printBudget(startDate, endDate, transactions)
}

func CustomMonth() {
	// get start and end dates
	var startDate, endDate time.Time
	var err error
	// refactor this; exercise (should be in util)
	for {
		startDateStr := util.ParseInputWithMessage("Custom Start Date (mm-dd-yyyy): ")
		startDate, err = time.Parse("01-02-2006", startDateStr)
		if err != nil {
			fmt.Println("Invalid date format yyyy/mm/dd")
			continue
		}
		break
	}
	for {
		endDateStr := util.ParseInputWithMessage("Custom End Date (mm-dd-yyyy; not inclusive): ")
		endDate, err = time.Parse("01-02-2006", endDateStr)
		if err != nil {
			fmt.Println("Invalid date format yyyy/mm/dd")
			continue
		}
		break
	}
	transactions := []tra.Transaction{}
	if err := tra.TransactionByDate(&transactions, startDate, endDate); err != nil {
		fmt.Println("Unable to get transactions by date", err)
		return
	}
	printBudget(startDate, endDate, transactions)
}

func BudgetAllocation() {
	// show all the categories and link the amounts saved
	// have them enter in a number to alter the amount for the selected
	categories := []cat.Category{}
	if err := cat.List(&categories); err != nil {
		fmt.Println("BudgetAllocation: unable to get categories:", err)
		fmt.Println("Press 'enter' to continue")
		util.ParseInput()
		return
	}

	budgetAllocation := []bud.BudgetAllocation{}
	if err := bud.List(&budgetAllocation); err != nil {
		fmt.Println("BudgetAllocation: unable to get budget allocation:", err)
		fmt.Println("Press 'enter' to continue")
		util.ParseInput()
		return
	}
	for {
		util.ClearScreen()
		fmt.Println("Budget Allocation")
		fmt.Println("")
		totalBudget := 0.0
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
		fmt.Fprintln(writer, "Id\tCategory\tBudget Amount")
		fmt.Fprintln(writer, "--\t----------\t-------")
		for _, category := range categories {
			amount := FindBudgetAllocation(category.Id, budgetAllocation)
			totalBudget += amount
			fmt.Fprintf(writer, "%d\t%s\t%0.2f\n", category.Id, category.CategoryName, amount)
		}
		fmt.Fprintln(writer, "")
		fmt.Fprintf(writer, "Total    \t           \t%0.2f\n", totalBudget)
		writer.Flush()
		fmt.Println("")
		selection := util.ParseInputIntWithMessage("Choose a category to change the budget allocation (0 - exit) ")
		if selection == 0 {
			break
		}
		amount := util.ParseInputFloatWithMessage(fmt.Sprintf("New budget allocation amount for [%s]: ", categories[selection-1].CategoryName))
		if err := bud.Upsert(selection, amount); err != nil {
			fmt.Println("Unable to insert/update budget allocation, please try again")
			break
		}
		// get the list again
		if err := bud.List(&budgetAllocation); err != nil {
			fmt.Println("printBudget: unable to get budget allocation:", err)
			return
		}
	}
}

func printBudget(startDate, endDate time.Time, transactions []tra.Transaction) {
	transactionMap := make(map[int]float64)
	for _, t := range transactions {
		if _, ok := transactionMap[t.CategoryId]; !ok {
			transactionMap[t.CategoryId] = 0
		}
		transactionMap[t.CategoryId] += t.Amount
	}
	// get categories and put into map, for easy lookup
	categories := []cat.Category{}
	if err := cat.List(&categories); err != nil {
		fmt.Println("printBudget: unable to get categories:", err)
		fmt.Println("Press 'enter' to continue")
		util.ParseInput()
		return
	}
	budgetAllocation := []bud.BudgetAllocation{}
	if err := bud.List(&budgetAllocation); err != nil {
		fmt.Println("printBudget: unable to get budget allocation:", err)
		fmt.Println("Press 'enter' to continue")
		util.ParseInput()
		return
	}
	categoryMap := make(map[int]string)
	for _, c := range categories {
		categoryMap[c.Id] = c.CategoryName
	}
	util.ClearScreen()
	fmt.Printf("Report Range: %s - %s\n", startDate.Format("01/02/2006"), endDate.Format("01/02/2006"))
	fmt.Println("")
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(writer, "Category\tAmount\tBudget\tNote")
	fmt.Fprintln(writer, "----------\t---------\t----------\t--------")
	for categoryId, Amount := range transactionMap {
		// TODO: leave for homework
		categoryName := categoryMap[categoryId]
		if !strings.Contains(strings.ToLower(categoryName), "payment") {
			budgetAmount := FindBudgetAllocation(categoryId, budgetAllocation)
			note := ""
			if budgetAmount > 0 {
				if Amount > budgetAmount {
					note = "*** OVERBUDGET ***"
				}
				if Amount == budgetAmount {
					note = "At Budget"
				}
				percent := int(Amount / budgetAmount * 100.0)
				if note == "" && percent > 80 {
					note = "Approaching budget"
				}
			}
			fmt.Fprintf(writer, "%s\t%0.2f\t%0.2f\t%s\n", categoryName, Amount, budgetAmount, note) // original
		}
	}
	writer.Flush()
	fmt.Println("")
	fmt.Println("Press 'enter' to continue")
	util.ParseInput()
}

// TODO: very unefficient, make a map
func FindBudgetAllocation(categoryId int, budAll []bud.BudgetAllocation) float64 {
	for _, bud := range budAll {
		if bud.CategoryId == categoryId {
			return bud.Amount
		}
	}
	return 0.0
}
