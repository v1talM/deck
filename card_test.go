package blackjack

import "fmt"

func ExampleCard()  {
	fmt.Println(Card{Rank: Ace, Suit: Diamond})

	// output:
	// Diamond of Aces
}
