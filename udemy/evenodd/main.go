package main

import "fmt"

func main() {
	var ints []int

	for i := range 11 {
		ints = append(ints, i)
	}

	for _, l := range ints {
		if l%2 == 0 {
			fmt.Printf("%v is even\n", l)
		} else {
			fmt.Printf("%v is odd\n", l)
		}
	}
}
