package main

func main() {
	deck := newDeck()
	_, deck = deal(deck, 3)
	// fmt.Println("Hand")
	// hand.print()
	// fmt.Println("Deck")
	// deck.print()
	deck.saveToFile("foo.txt")
	newDeck := newDeckFromFile("foo.txt")
	newDeck.shuffle().print()
}
