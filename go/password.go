package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var password string
	var newCharSet string
	passLen := flag.Int("length", 1234, "number of characters for the password")
	passSymbols := flag.Bool("special", false, "use special characters")
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	symbolSet := "!@#$%&*"

	newCharSet += charSet

	flag.Parse()

	if *passLen < 16 {
		fmt.Printf("You must enter a value of 16 or greater")
		os.Exit(1)
	}
	if *passSymbols {
		newCharSet += symbolSet
	}

	for i := 0; i < *passLen; i++ {
		randChar := rand.Intn(len(newCharSet))
		newChar := newCharSet[randChar]
		if i == 0 || i == *passLen-1 {
			randChar := rand.Intn(len(charSet))
			newChar = charSet[randChar]
		}
		password += string(newChar)
	}

	fmt.Printf("Your new password is: %s\n", password)
}
