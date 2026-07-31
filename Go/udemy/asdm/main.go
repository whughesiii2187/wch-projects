package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["ruby"] = 4.8
	courseRatings["java"] = 4.7

	fmt.Println(courseRatings)
}
