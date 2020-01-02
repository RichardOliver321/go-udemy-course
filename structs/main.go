package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zipc  int
}

func main() {
	bob := person{
		firstName: "Bob",
		lastName:  "Bobby",
		contact: contactInfo{
			email: "bob@bobmail.com",
			zipc:  12,
		},
	}

	bob.updateFirstName("booop")
	bob.print()
}

func (pointerToPerson *person) updateFirstName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
