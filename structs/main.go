package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// sai := person{"saikrishna","nelavalli"}
	// sai := person{firstName: "saikrishna", lastName: "nelavalli"}
	var sai person
	sai.firstName = "saikrishna"
	sai.lastName = "Nelavalli"
	fmt.Println(sai)
	fmt.Printf("%+v", sai)
}
