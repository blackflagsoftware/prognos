package util

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var (
	clearMap map[string]func()
)

func init() {
	clearMap = make(map[string]func())
	clearMap["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearMap["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clearMap["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ParseInput() string {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func ParseInputWithMessage(msg string) string {
	fmt.Print(msg)
	return ParseInput()
}

func ParseInputStringWithMessageCompare(msg string, compare string) string {
	fmt.Print(msg)
	str := ParseInput()
	if str == "" {
		return compare
	}
	return str
}

func ParseInputIntWithMessage(msg string) int {
	for {
		fmt.Print(msg)
		intStr := ParseInput()
		intInt, err := strconv.Atoi(intStr)
		if err == nil {
			return intInt
		}
		fmt.Println("Not a valid integer, try again")
	}
}

func ParseInputIntWithMessageCompare(msg string, compare int) int {
	for {
		fmt.Print(msg)
		intStr := ParseInput()
		if intStr == "" {
			return compare
		}
		intInt, err := strconv.Atoi(intStr)
		if err == nil {
			return intInt
		}
		fmt.Println("Not a valid integer, try again")
	}
}

func ParseInputFloatWithMessage(msg string) float64 {
	for {
		fmt.Print(msg)
		floatStr := ParseInput()
		floatFloat, err := strconv.ParseFloat(floatStr, 10)
		if err == nil {
			return floatFloat
		}
		fmt.Println("Not a valid float, try again")
	}
}

func ParseInputFloatWithMessageCompare(msg string, compare float64) float64 {
	for {
		fmt.Print(msg)
		floatStr := ParseInput()
		if floatStr == "" {
			return compare
		}
		floatFloat, err := strconv.ParseFloat(floatStr, 10)
		if err == nil {
			return floatFloat
		}
		fmt.Println("Not a valid float, try again")
	}
}

func ParseInputBoolWithMessage(msg string) bool {
	for {
		fmt.Print(msg)
		boolStr := ParseInput()
		if strings.ToLower(boolStr) == "true" {
			return true
		}
		if strings.ToLower(boolStr) == "false" {
			return false
		}
		fmt.Println("true | false, try again")
	}
}

func ParseInputBoolWithMessageCompare(msg string, compare bool) bool {
	for {
		fmt.Print(msg)
		boolStr := ParseInput()
		if boolStr == "" {
			return compare
		}
		if strings.ToLower(boolStr) == "true" {
			return true
		}
		if strings.ToLower(boolStr) == "false" {
			return false
		}
		fmt.Println("true | false, try again")
	}
}

func ClearScreen() {
	clearFunc, ok := clearMap[runtime.GOOS]
	if !ok {
		fmt.Println("\n *** Your platform is not supported to clear the terminal screen ***")
		return
	}
	clearFunc()
}

func BasicPrompt(mainMessage []string, prompts []string, acceptablePrompts []string, exitString string) string {
	for {
		//clearScreen()
		for _, msg := range mainMessage {
			fmt.Println(msg)
		}
		fmt.Println("")
		for _, prompt := range prompts {
			fmt.Println(prompt)
		}
		if exitString != "" {
			// just in case you don't want to show this line
			fmt.Printf("(%s) to exit", exitString)
		}
		fmt.Println("")
		fmt.Println("")
		fmt.Print("Selection Choice: ")
		selection := ParseInput()
		if strings.ToLower(selection) == exitString {
			return exitString
		}
		found := false
		for _, acceptablePrompt := range acceptablePrompts {
			if strings.ToLower(selection) == acceptablePrompt {
				found = true
				break
			}
		}
		if !found {
			fmt.Print("Invalid selection, try again, press 'enter' to continue:")
			ParseInput()
			ClearScreen()
			continue
		}
		return strings.ToLower(selection)
	}
}

func AskYesOrNo(msg string) (answer bool) {
	for {
		msg = fmt.Sprintf("%s (y/n)? ", msg)
		fmt.Print(msg)
		def := ParseInput()
		switch def {
		case "y", "Y":
			answer = true
		case "n", "N":
			answer = false
		default:
			fmt.Println("Invalid value, get it together (y or n)!")
			continue
		}
		break
	}
	return
}
