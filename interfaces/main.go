package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// very custom logic for genereting an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// very custom logic for genereting an spanish greeting
	return "Hola!"
}
