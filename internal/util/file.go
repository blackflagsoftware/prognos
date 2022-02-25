package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blackflagsoftware/prognos/config"
)

func createFile(filePathWithName string) {
	if _, err := os.Stat(filePathWithName); os.IsNotExist(err) {
		if err := os.WriteFile(filePathWithName, []byte("[]"), 0644); err != nil {
			fmt.Println("error creating storage folder:", err)
		}
	}
}

func OpenFile(name string, obj interface{}) error {
	filePathWithName := fmt.Sprintf("%s/%s", config.FilePath, name)
	createFile(filePathWithName)
	content, err := os.ReadFile(filePathWithName)
	if err != nil {
		return fmt.Errorf("unable to open file: %s", err)
	}
	if err := json.Unmarshal(content, obj); err != nil {
		return fmt.Errorf("unable to decode file: %s", err)
	}
	return nil
}

func SaveFile(name string, obj interface{}) error {
	filePathWithName := fmt.Sprintf("%s/%s", config.FilePath, name)
	createFile(filePathWithName)
	content, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("unable to encode file: %s", err)
	}
	if err := os.WriteFile(filePathWithName, content, 0644); err != nil {
		return fmt.Errorf("unable to save file: %s", err)
	}
	return nil
}
