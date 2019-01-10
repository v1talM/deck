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

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}

// 默认排序
func defaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// 根据指定方式排序
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// 打乱顺序
func Shuffle(cards []Card) []Card {
	var ret = make([]Card, len(cards))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// 有n张joker
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

// 过滤掉指定牌
func Filter(fn func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !fn(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

// 从小到大排序
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

// 从大到小排序
func Greater(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) > absRank(cards[j])
	}
}

func absRank(c Card) int {
	return int(c.Suit) * int(MaxRank) + int(c.Rank)
}