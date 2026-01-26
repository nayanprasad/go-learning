package main

import "fmt"

func main() {
	fmt.Println("cards game")

	deck := newDeck()

	// deck.print()

	deck.suffle()

	// deck.print()

	hand, remaining := deal(deck, 3)

	println("hand")

	hand.print()

	println("remaining card")

	remaining.print()
}
