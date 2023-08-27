package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jim.print()
	//jim.firstName = "Jimmy"
	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
