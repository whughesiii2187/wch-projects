package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type outputter interface {
	Save() error
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo content: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	value := bufio.NewReader(os.Stdin)
	text, err := value.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSpace(text)

	return text
}

func saveData(data outputter) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}
	fmt.Println("Saving the note successful")
	return nil
}

func outputData(data outputter) error {
	data.Display()
	return saveData(data)
}
