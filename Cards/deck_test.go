package main

import (
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52")
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card 'Two of Clubs'")
	}
}

func TestSaveAndLoad(t *testing.T) {
	os.Remove("test")

	d1 := newDeck()
	d1.saveToFile("test")
	
	d2 := newDeckFromFile("test")

	if strings.Join(d1,";") != strings.Join(d2,";") {
		t.Errorf("Loaded deck not identical to saved one")
	}

	os.Remove("test")
}
