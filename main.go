package main

import "main/deck"

func main() {
	// deck code
	cards := deck.NewDeckFromFile("my_cards")
	cards.Shuffle()
	cards.Print()

	// map definitions
	// colors := map[string]string{
	// 	"red":   "#ff000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }
	// var colors map[string]string
	// colors := make(map[string]string)

}
