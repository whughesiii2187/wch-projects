package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000
	var depositAmount float64
	var withdrawlAmount float64

	fmt.Println("Welcome to Go Bank!\nWhat do you want to do?")
	fmt.Println("1. Check Balance\n2. Deposit Money\n3. Withdraw Money\n4. Exit")

	var userInput int
	fmt.Print("Your choice: ")
	_, _ = fmt.Scan(&userInput)

	if userInput == 1 {
		fmt.Println("Your current balance:", accountBalance)
	} else if userInput == 2 {
		fmt.Print("Your deposit amount: ")
		_, err := fmt.Scan(&depositAmount)
		if err == nil && depositAmount <= 0 {
			fmt.Println("Invalid amount! Must be greater than 0.")
			return
		}
		accountBalance += depositAmount
		fmt.Println("Balance Updated! New amount:", accountBalance)
	} else if userInput == 3 {
		fmt.Print("Amount to withdraw: ")
		_, err := fmt.Scan(&withdrawlAmount)

		if err == nil && withdrawlAmount <= 0 {
			fmt.Println("Invalid amount! Must be greater than 0.")
			return
		}

		if withdrawlAmount > accountBalance {
			fmt.Println("Insufficient funds! Choose a smaller amount to withdraw")
			return
		}

		accountBalance -= withdrawlAmount
		fmt.Println("Balance Updated! New amount:", accountBalance)

	} else {
		fmt.Println("Good bye!")
		// break
	}

	// switch userInput {
	// case 1:
	// 	fmt.Printf("Your Current Balance is: $%.2f\n", accountBalance)
	// case 2:
	// 	fmt.Print("How much money to add: ")
	// 	_, _ = fmt.Scan(&depositAmount)
	// 	fmt.Printf("Your new balance is: $%.2f\n", accountBalance+depositAmount)
	// case 3:
	// 	fmt.Print("How much money to remove: ")
	// 	_, _ = fmt.Scan(&withdrawlAmount)
	// 	fmt.Printf("Your new balance is $%.2f\n", accountBalance-withdrawlAmount)
	// case 4:
	// 	fmt.Println("Good Bye!")
	// }
}
