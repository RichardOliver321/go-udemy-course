package main

import "fmt"

func main() {
	cards := readDeckFromFile("latestDeck")
	cards.shuffle()
	fmt.Println(cards.toString())
}
