package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"   /* variable initialization*/
	// card = "Five of Diamonds" /* variable assign value*/
	// card := newCard()
	// cards := newDeck()
	// calling function with multiple return values
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	// call := "Hi! How are you?"
	// //converting to bytes
	// fmt.Println([]byte(call))

	// cards := newDeckFromFile("my_deck")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
