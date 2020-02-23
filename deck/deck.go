package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck is a slice of strings
type Deck []string

// NewDeck creates a new deck
func NewDeck() Deck {
	cards := Deck{}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print prints the entire deck
func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal deals a deck into 2 subdivided decks
func Deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// ToString returns string representation of the deck
func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}

// SaveToFile saves deck to a file
func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

// NewDeckFromFile gets a new deck from file
func NewDeckFromFile(filename string) Deck {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bytes), ",")
	return Deck(s)
}

// Shuffle shuffles the deck
func (d Deck) Shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range d {
		newPos := r.Intn(len(d) - 1)

		d[i], d[newPos] = d[newPos], d[i]
	}
}
