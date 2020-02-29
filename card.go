package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(name string) error {
	return ioutil.WriteFile(name, []byte(d.toString()), 0666)
}

func deckFromFile(name string) deck {
	byt, err := ioutil.ReadFile(name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return deck(strings.Split(string(byt), ","))
}
