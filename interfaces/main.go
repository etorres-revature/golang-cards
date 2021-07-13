package main

import "fmt"

type bot interface {
	getGreeting() string
}
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

func (eb englishBot) getGreeting() string {
	//cutom logic for generating a greeting in the English language
	return "Hi, there..."
}

func (sb spanishBot) getGreeting() string {
	//cutom logic for generating a greeting in the Spanish language
	return "Hola, ese..."
}
