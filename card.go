// go generate stringer --type=Suit,Rank
package blackjack

import (
	"fmt"
	"sort"
)

type Suit uint8
type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	MinRank = Ace
	MaxRank = King
)

type Card struct {
	Suit
	Rank
}

type Deck []Card

var Suits = [...]Suit{Spade, Diamond, Club, Heart}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Suit.String(), c.Rank.String())
}

func NewDeck(opts ...func([]Card) []Card) []Card {
	var cards Deck
	for _, suit := range Suits {
		for i := MinRank; i <= MaxRank; i++ {
			cards = append(cards, Card{suit, i})
		}
	}
	for _, opt := range opts {
		opt(cards)
	}
	return cards
}

func defaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit) * int(MaxRank) + int(c.Rank)
}