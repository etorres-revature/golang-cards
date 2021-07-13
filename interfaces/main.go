package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}
type russianBot struct{}
type frenchBot struct{}
type italianBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	rb := russianBot{}
	fb := frenchBot{}
	ib := italianBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(rb)
	printGreeting(fb)
	printGreeting(ib)
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

func (rb russianBot) getGreeting() string {
	//cutom logic for generating a greeting in the Russian language
	return "Das vidanya..."
}

func (fb frenchBot) getGreeting() string {
	//cutom logic for generating a greeting in the French language
	return "Bonjour, mademoiselle..."
}

func (ib italianBot) getGreeting() string {
	//cutom logic for generating a greeting in the Italian language
	return "Ciao, bella..."
}
