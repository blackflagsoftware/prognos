package util

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/blackflagsoftware/prognos/config"
)

func createFile(filePathWithName string, defaultObject string) {
	if _, err := os.Stat(filePathWithName); os.IsNotExist(err) {
		err := os.MkdirAll(config.FilePath, 0755)
		if err != nil {
			fmt.Println("error making folder structure", err)
			return
		}
		if err := os.WriteFile(filePathWithName, []byte(defaultObject), 0644); err != nil {
			fmt.Println("error creating storage folder:", err)
		}
	}
}

func OpenFile(name string, obj interface{}) error {
	defaultObject := "{}"
	rf := reflect.ValueOf(obj)
	if rf.Elem().Kind() == reflect.Slice {
		defaultObject = "[]"
	}
	filePathWithName := fmt.Sprintf("%s/%s", config.FilePath, name)
	fmt.Println(defaultObject)
	createFile(filePathWithName, defaultObject)
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
	defaultObject := "{}"
	rf := reflect.ValueOf(obj)
	if rf.Elem().Kind() == reflect.Slice {
		defaultObject = "[]"
	}
	filePathWithName := fmt.Sprintf("%s/%s", config.FilePath, name)
	createFile(filePathWithName, defaultObject)
	content, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to encode file: %s", err)
	}
	if err := os.WriteFile(filePathWithName, content, 0644); err != nil {
		return fmt.Errorf("unable to save file: %s", err)
	}
	return nil
}
