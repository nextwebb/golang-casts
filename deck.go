package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck [] string // we can think about this as the class declaration

// like static methods in js
func newDeck() deck {
	// var cards []string
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// go receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

