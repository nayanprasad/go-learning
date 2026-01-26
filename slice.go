package main

import "fmt"

func main() {

	cards := []string{"a", "b", "c"}

	fmt.Println(cards)

	for i, card := range cards {
		println(i, card)
	}

	for _, card := range cards {
		println(card)
	}
}
