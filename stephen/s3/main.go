package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	//err := cards.saveToFile("my_cards")
	//if err != nil {
	//	fmt.Println(err)
	//}
	cards.shuffle()
	cards.print()
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()
	//fDeck := newDeckFromFile("my_cards")
	//fDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
