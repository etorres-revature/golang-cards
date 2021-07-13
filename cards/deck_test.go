package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("That's not enough cards! - Expected deck length of 52 cards, but got %v", len(d))
	}
}
