package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	//cutom logic for generating a greeting in the English language
	return "Hi, there..."
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (sb spanishBot) getGreeting() string {
	//cutom logic for generating a greeting in the Spanish language
	return "Hola, ese..."
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
