package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// sai := person{"saikrishna","nelavalli"}
	// sai := person{firstName: "saikrishna", lastName: "nelavalli"}
	// var sai person
	// sai.firstName = "saikrishna"
	// sai.lastName = "Nelavalli"
	sai := person{
		firstName: "saikrishna",
		lastName:  "Nelavalli",
		contact: contactInfo{
			email:   "sai@gmail.com",
			zipCode: 524126,
		},
	}
	// saiPointer := &sai
	sai.updateName("sai")
	sai.print()
}

func (pointerToSai *person) updateName(newFirstName string) {
	(*pointerToSai).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
