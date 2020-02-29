package main

func main() {

	card := newDeck()
	card.saveToFile("sai.txt")
	cc := deckFromFile("sai.txt")
	cc.print()
}
