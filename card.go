package main

import "fmt"

type deck []string

func newDeck() deck {
	dk := deck{}
	cs := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cv := []string{"Aces", "Two", "Three", "Four"}
	for _, s := range cs {
		for _, v := range cv {
			dk = append(dk, v+" of "+s)
		}
	}
	return dk
}

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}
