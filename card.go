// go generate stringer --type=Suit,Rank
package blackjack

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func NewDeck(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range Suits {
		for rank := MinRank; rank  <= MaxRank; rank ++ {
			cards = append(cards, Card{suit, rank })
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func defaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Shuffle(cards []Card) []Card {
	var ret = make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

func HaveJokers(number int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < number; i++ {
			j := Card{
				Rank: Rank(i),
				Suit: Joker,
			}
			cards = append(cards, j)
		}
		return cards
	}
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func Greater(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) > absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit) * int(MaxRank) + int(c.Rank)
}