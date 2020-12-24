package main

func main()  {
	cards := newDeck()
	// deck datatype , calling the go receiver
	hand, remainingCards := deal(cards, 5)
	// the two variables returns a datatype of deck, so it has access to the go receivers belonging to type "deck"
	hand.print()
	remainingCards.print()
}


