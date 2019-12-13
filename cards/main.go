package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"   /* variable initialization*/
	// card = "Five of Diamonds" /* variable assign value*/
	card := newCard()
	fmt.Println("card --------------", card)

}

func newCard() string {
	return "Five of Diamonds"
}
