package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new custom type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver function which accepts a receiver. By adding we got ability to call this method on any deck; ex: cards.print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// select range on slices to pull out subset of slices
func deal(d deck, handSize int) (deck, deck) {
	// d[:handSize] - start from zero till the index
	// d[handSize:] - start from index till the end
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	sliceOfString := []string(d)
	return strings.Join(sliceOfString, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	sliceOfString := strings.Split(string(byteSlice), ",")
	return deck(sliceOfString)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
