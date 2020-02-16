package main

import (
	"fmt"
  "io/ioutil"
  "os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
  return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
  bytes, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }
  s := strings.Split(string(bytes), ",")
  return deck(s)
}