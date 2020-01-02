package main

import "fmt"

func main() {
	cards := readDeckFromFile("latestDeck")
	fmt.Println(cards.toString())
}
