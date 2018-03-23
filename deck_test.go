package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// When
	d := newDeck()

	// Then
	if len(d) != 52 {
		t.Errorf("Should have 52 cards in a deck, but got %v", len(d))
	}
	if d[0] != "Ace of Hearts" {
		t.Errorf("First card in the deck should have been \"Ace of Hearts\", but was \"%v\"", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("First card in the deck should have been \"King of Clubs\", but was \"%v\"", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// Given
	os.Remove("_desktesting")
	deck := newDeck()

	// When
	deck.saveToFile("_desktesting")
	loadedDeck := newDeckFromFile("_desktesting")

	// Then
	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected loaded deck to have %v cards, but actually had %v cards", len(deck), len(loadedDeck))
	}
	os.Remove("_desktesting")
}
