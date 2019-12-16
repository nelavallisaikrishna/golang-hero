package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Error("Expected deck Length 16, but got %V", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Error("Expected Ace of Spades, but got %V", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Error("Expected Four of Clubs, but got %V", d[len(d)-1])
	}
}
