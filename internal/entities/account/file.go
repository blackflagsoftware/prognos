package account

import (
	"encoding/json"
	"fmt"
	"os"
)

func DataRead(acc *Account) error {
	fmt.Println("In Read")
	return nil
}

func DataList(acc *[]Account) error {
	fmt.Println("In List")
	return nil
}

func DataCreate(acc Account) error {
	accs := []Account{}
	file := "../../data/prognos_data/account"
	fileContent, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Unable to open account file: %s for saving", file)
	}
	if err := json.Unmarshal(fileContent, &accs); err != nil {
		return fmt.Errorf("Unable to convert file to account data: %s", err)
	}
	maxId := 0
	for _, accObj := range accs {
		if accObj.Id > maxId {
			maxId = accObj.Id
		}
	}
	acc.Id = maxId + 1
	accs = append(accs, acc)
	fileContent, err = json.MarshalIndent(accs, "", "  ")
	if err != nil {
		return fmt.Errorf("Unable to convert account to file data: %s", err)
	}
	if err := os.WriteFile(file, fileContent, 0644); err != nil {
		return fmt.Errorf("Unable to write account file: %s", err)
	}
	return nil
}

func DataUpdate(acc Account) error {
	accs := []Account{}
	file := "../../data/prognos_data/account"
	fileContent, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Unable to open account file: %s for saving", file)
	}
	if err := json.Unmarshal(fileContent, &accs); err != nil {
		return fmt.Errorf("Unable to convert file to account data: %s", err)
	}
	for i := range accs {
		if accs[i].Id == acc.Id {
			accs[i] = acc
			break
		}
	}
	fileContent, err = json.MarshalIndent(accs, "", "  ")
	if err != nil {
		return fmt.Errorf("Unable to convert account to file data: %s", err)
	}
	if err := os.WriteFile(file, fileContent, 0644); err != nil {
		return fmt.Errorf("Unable to write account file: %s", err)
	}
	return nil
}

func DataDelete(acc Account) error {
	fmt.Println("In Delete")
	return nil
}
