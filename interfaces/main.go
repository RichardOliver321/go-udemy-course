package main

import "fmt"

type chatBot interface {
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

func (englishBot) getGreeting() string {
	return "hello"
}

func (spanishBot) getGreeting() string {
	return "Ola"
}

func printGreeting(b chatBot) {
	fmt.Println(b.getGreeting())
}
