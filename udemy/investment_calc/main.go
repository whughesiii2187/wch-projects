package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Enter an investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	futureRealValue := calculateFutureRealValue(futureValue, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) float64 {
	return investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
}

func calculateFutureRealValue(futureValue, years float64) float64 {
	return futureValue / math.Pow(1+inflationRate/100, float64(years))
}
