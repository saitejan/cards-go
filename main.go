package main

func main() {

	card := newDeck()

	card.shuffle()
	card.print()

	// card.saveToFile("sai.txt")
	// cc := deckFromFile("sai.txt")
	// cc.print()
}
