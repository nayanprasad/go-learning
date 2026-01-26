package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "diamonds", "heats"}
	cardValues := []string{"ace", "two", "three", "four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, d := range d {
		fmt.Println(i, ": ", d)
	}
}

func (d deck) suffle() {
	rand.Seed(time.Now().Unix())

	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	if handSize > len(d) {
		fmt.Println("Error: invalide handeSize vlaue")
		os.Exit(1)
	}

	return d[:handSize], d[handSize:]
}
