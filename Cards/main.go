package main

import "fmt"

func main() {
	//cards := []string{"Ace of Hearts", newCards()}
	//cards := deck{"Ace of Diamonds", newCards()}
	//cards = append(cards, "Seven of Spades")
	cards := newDeck()
	fmt.Println(cards[0])

	randomCards := cards.handful(5)
	//randomCards.print()

	handOne, handTwo := cards.deal(22)
	//handOne.print()
	//handTwo.print()

	handOne.saveToFile("one")
	handTwo.saveToFile("two")
	randomCards.saveToFile("three")

	newDeck := newDeckFromFile("newdeck")
	newDeck.print()

	cards.shuffle()
	cards.print()
}

func newCards() string {
	return "Nine of Spades"
}
