package main

import (
	"os"
	"testing"
)

// t.Errorf("Expected deck length of 16, but got", len(d))
// to
// t.Errorf("Expected deck length of 16, but got %v", len(d))

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card to be Two of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Clubs" {
		t.Errorf("Expected last card to be Ace of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected deck length of %v, but got %v", len(deck), len(loadedDeck))
	}

	os.Remove("_decktesting")
}
