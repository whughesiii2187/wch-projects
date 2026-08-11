package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"white": "#ffffff",
	}
	// var colors map[string]string
	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v:%v\n", color, hex)
	}
}
