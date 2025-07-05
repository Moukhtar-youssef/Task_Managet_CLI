package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

func filepath() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
		return ""
	}
	return path.Join(cwd, "tasks.json")
}

func ReadFromFile() ([]Task, error) {
	filepath := filepath()

	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		fmt.Println("No task file found. Creating the file ....")
		file, err := os.Create(filepath)
		if err != nil {
			return nil, fmt.Errorf("Error creating file: %v", err)
		}
		os.WriteFile(filepath, []byte("[]"), os.ModeAppend.Perm())

		defer file.Close()

		return []Task{}, nil
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Error opening task file: %v", err)
	}

	defer file.Close()

	tasks := []Task{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, fmt.Errorf("Error Decoding task file: %v", err)
	}

	return tasks, nil
}

func SaveToFile(tasks []Task) error {
	filepath := filepath()

	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("Error saving to task file: %v", err)
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(&tasks)
	if err != nil {
		return fmt.Errorf("Error Encoding taks file: %v", err)
	}

	return nil
}
