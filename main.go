package main

func main() {

	card := newDeck()
	a, b := deal(card, 5)
	a.print()
	b.print()

}
