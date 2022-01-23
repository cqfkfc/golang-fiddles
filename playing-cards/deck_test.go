package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if (len(d) != 52) {
		t.Errorf("one deck should have 52 cards")
	}
}