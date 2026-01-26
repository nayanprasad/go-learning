package main

import "fmt"

type bot interface {
	getGreeting() string // a custom logic method
}

// if any of the follwoing type associated with ' getGreeting() string ' then it will be part of bot interfaces

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (sb englishBot) getGreeting() string {
	return "Hi"
}

func (sp spanishBot) getGreeting() string {
	return "Hola"
}
