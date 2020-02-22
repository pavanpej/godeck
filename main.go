package main

import "main/deck"

func main() {
	cards := deck.NewDeckFromFile("my_cards")
	cards.Shuffle()
	cards.Print()
}
