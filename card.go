// go generate stringer --type=Suit,Rank
package blackjack

import "fmt"

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

func NewDeck() []Card {
	var cards Deck
	for _, suit := range Suits {
		for i := MinRank; i <= MaxRank; i++ {
			cards = append(cards, Card{suit, i})
		}
	}
	return cards
}