package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"   /* variable initialization*/
	// card = "Five of Diamonds" /* variable assign value*/
	// card := newCard()
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards,"Six of Diamonds")
	fmt.Println("cards --------------", cards)
	for i, card := range cards{
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
