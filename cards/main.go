package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"   /* variable initialization*/
	// card = "Five of Diamonds" /* variable assign value*/
	// card := newCard()
	cards := newDeck()
	fmt.Println("cards --------------", cards)
	cards.print()

}

