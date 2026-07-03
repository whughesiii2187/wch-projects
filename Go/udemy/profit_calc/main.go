package main

import "fmt"

func main() {
	// var revenue, expenses, taxRate float64

	revenue := getUserInput("Revenue: ")
	// fmt.Print("Enter revenue: ")
	// fmt.Scan(&revenue)

	expenses := getUserInput("Expenses: ")
	// fmt.Print("Enter expenses: ")
	// fmt.Scan(&expenses)

	taxRate := getUserInput("Tax Rate: ")
	// fmt.Print("Enter tax rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calculations(revenue, expenses, taxRate)
	fmt.Printf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt - (ebt * taxRate / 100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
