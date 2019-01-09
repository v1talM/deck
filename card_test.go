package blackjack

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 13 * 4 {
		t.Error("Wrong number of cards")
	}
}