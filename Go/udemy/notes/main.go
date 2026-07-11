package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

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
