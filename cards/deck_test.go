package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 32 {
		t.Errorf("Expected deck length of 32 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Seven of Diamonds" {
		t.Errorf("Expected Ace of Spades but got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndReadDeckFromFile(t *testing.T) {

	d := newDeck()

	fileName := "_decktesting"

	d.saveToFile(fileName)

	loadedDeck := readDeckFromFile(fileName)

	if len(loadedDeck) != 32 {
		t.Errorf("Expected length of 32 but got %v", len(loadedDeck))
	}

}
