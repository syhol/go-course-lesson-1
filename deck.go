package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	var values = []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}
	var suits = []string{
		"Hearts",
		"Diamonds",
		"Spades",
		"Clubs",
	}

	deck := deck{}
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, value+" of "+suit)
		}
	}
	return deck
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, amount int) (deck, deck) {
	return d[:amount], d[amount:]
}

func (d deck) withCard(card string) deck {
	return append(d, card)
}

func (d deck) saveToFile(filename string) {
	string := d.toString()
	bytes := []byte(string)
	ioutil.WriteFile(filename, bytes, os.FileMode(0644))
}

func newDeckFromFile(filename string) deck {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	string := string(bytes)
	return fromString(string)
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",\n")
}

func fromString(s string) deck {
	return deck(strings.Split(s, ",\n"))
}

func (d deck) shuffle() deck {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		pos := r.Intn(len(d) - 1)
		d[i], d[pos] = d[pos], d[i]
	}
	return d
}
