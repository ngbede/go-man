package main

import (
	"os"
	"testing"
)

// Go tests always begin with uppercase
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but received %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected deck to have initial card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected deck to have final card of King of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")
	d := newDeck()
	d.saveToFile("_decktesting")

	// read deck from a file
	loadDeck := readFile("_decktesting.txt")
	if len(loadDeck) != 52 {
		t.Errorf("Expected deck length of 52 but received %v", len(loadDeck))
	}
	os.Remove("_decktesting.txt")
}
