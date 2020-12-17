package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck [] string // we can think about this as the class declaration

// go receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

