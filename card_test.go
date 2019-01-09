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

func TestDefaultSort(t *testing.T)  {
	cards := NewDeck(defaultSort)
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("Expected the first card is Ace of Spade, received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := NewDeck(Sort(Less))
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("Expected the first card is Ace of Spade, received:", cards[0])
	}

	cards = NewDeck(Sort(Greater))
	expected = Card{Rank: King, Suit: Heart}
	if cards[0] != expected {
		t.Error("Expected the first card is King of Heart, received:", cards[0])
	}
}