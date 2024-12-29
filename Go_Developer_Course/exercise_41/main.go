package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	// Create a new person struct
	c := contactInfo{
		email:   "name@domain.com",
		zipCode: 12345,
	}
	p := person{
		firstName: "John",
		lastName:  "Doe",
		contact:   c,
	}
	a := person{"Jane", "Doe", c}

	// Print the person's first name
	fmt.Println(p)
	fmt.Println(p.firstName + " " + p.lastName)
	fmt.Println(a.firstName + " " + a.lastName + " " + a.contact.email + " " + fmt.Sprint(a.contact.zipCode))
	// pPointer := &p
	// pPointer.updateName("Kalyan")
	p.updateName("Kalyan")
	// p.updateName("Jane")
	fmt.Println(p.fullName())
	fmt.Printf("%+v\n", a)
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Doe"
	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)
	fmt.Println(alex.fullName())
}
