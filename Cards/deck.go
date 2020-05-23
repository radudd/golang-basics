package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
// 'extends' a slice of strings
type deck []string

func (d deck) toString() string {
	//Conversion from deck to []string(ie []string(d)) works
	//as deck is already a slice of strings
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(filename string) error {
	//First convert the deck to string d.toString()
	//then convert the string to byte slice
	content := []byte(d.toString())
	return ioutil.WriteFile(filename, content, 0644)
}

func (d deck) handful(no int) deck {
	newDeck := deck{}
	for i := 1; i < no; i++ {
		index := rand.Intn(len(d) - 1)
		newDeck = append(newDeck, d[index])
	}
	return newDeck
}

func (d deck) shuffle() {

	for i := range d {
		//generate radom seed
		//check:
		//https://golang.org/pkg/math/rand/#NewSource
		//https://golang.org/pkg/time/#Time.UnixNano
		t := time.Now() //return a Time type
		tn := t.UnixNano()
		source := rand.NewSource(tn)
		random := rand.New(source)

		// we have a random number seed, proceed with substitution
		newPosition := random.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	fmt.Println("Deck:")
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(99)
	}
	s := strings.Split(string(bs), "\n")
	return deck(s)
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{
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
		"Juve",
		"Queen",
		"King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}