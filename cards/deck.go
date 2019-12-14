package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type deck with array of strings

type deck []string //custom type create

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Cluds"}
	cardValues := []string{"Ace", "two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Receiver creating
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, inhand int) (deck, deck) {
	return d[:inhand], d[inhand:]
}

/*Converting []string to string*/
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/*Save to Local drive*/

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
