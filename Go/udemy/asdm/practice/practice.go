package practice

import "fmt"

type Product struct {
	id    string
	desc  string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string{"coding", "gaming", "drumming"}
	fmt.Println(hobbies)
	// 2)
	fmt.Println(hobbies[0])
	newHobbies := hobbies[1:]
	fmt.Println(newHobbies)
	// 3)
	anotherHobbies := hobbies[0:2]
	fmt.Println(anotherHobbies)
	// 4)
	anotherHobbies = anotherHobbies[1:3]
	fmt.Println(anotherHobbies)
	// 5)
	goals := []string{"learn_go", "master_go"}
	fmt.Println(goals)
	// 6)
	goals[1] = "really_master_go"
	goals = append(goals, "automation")
	fmt.Println(goals)

	// 7)
	products := []Product{
		{
			"001",
			"First item",
			12.99,
		},
		{
			"002",
			"Second item",
			129.99,
		},
	}
	fmt.Println(products)

	products = append(products, Product{"003", "Third item", 50.00})
	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
