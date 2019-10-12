package main

import "fmt"

// Any types that have getGretting function is the part of the bot type
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGretting(eb)
	printGretting(sb)
}

func printGretting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi There"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}
