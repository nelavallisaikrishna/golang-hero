package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"   /* variable initialization*/
	// card = "Five of Diamonds" /* variable assign value*/
	// card := newCard()
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

}
