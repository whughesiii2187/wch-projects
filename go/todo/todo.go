package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const todoFile string = "todo.json"

func main() {
	// Check for existence of a file first, create if needed
	_, err := os.Stat(todoFile)
	if err != nil {
		fmt.Println("File does not exist, creating...")
		_, err := os.Create(todoFile)
		if err != nil {
			fmt.Printf("error with creating a file %v", err)
			return
		}
	}

	openTheFile, err := os.Open(todoFile)
	if err != nil {
		fmt.Printf("error opening file %s: %v", todoFile, err)
		return
	}

	fileInfo, err := io.ReadAll(openTheFile)
	if err != nil {
		fmt.Printf("error parsing data: %v", err)
		return
	}

	defer openTheFile.Close()

	var todoData []string
	if len(fileInfo) > 0 {
		err = json.Unmarshal(fileInfo, &todoData)
		if err != nil {
			fmt.Printf("error importing json: %v", err)
			return
		}
	}

	// Ask for input: A new todo item to and
	fmt.Println("Enter your todo item")
	todoInput := bufio.NewScanner(os.Stdin)
	todoInput.Scan()
	err = todoInput.Err()
	if err != nil {
		fmt.Printf("error reading todo: %v", err)
	}

	// append new item to end of json data
	todoData = append(todoData, todoInput.Text())

	// write data to file
	writeData, err := json.Marshal(todoData)
	if err != nil {
		fmt.Printf("error converting to json: %v", err)
	}

	err = os.WriteFile(todoFile, writeData, 0664)
	if err != nil {
		fmt.Printf("unable to write file: %v", err)
	}
	return
}
