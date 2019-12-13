package main

import "fmt"

// Create a new type deck with array of strings

type deck []string //custom type create

// Receiver creating
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
