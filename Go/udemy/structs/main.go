package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Parsons",
		contact: contactInfo{
			email: "jimmyp@gmail.com",
			zip:   46375,
		},
	}
	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
