package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck [] string // we can think about this as the class declaration

// like static methods in js
func newDeck() deck {
	// var cards []string
	cards := deck{} // choose data type

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
// slices are zero indexed

// go receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// expecting two separate values
	// [start - 5] // [5 - end]
 	return d[:handSize], d[handSize:]

}

