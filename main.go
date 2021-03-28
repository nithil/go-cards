package main

func main() {
	// cards := deck{"Ace of diamonds", newCard()}
	// cards = append(cards, "six of spades")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards.saveToFile("myCards")

	// cards := newDeckFromFile("myCardss")

	// cards.print()
}

// hand, remainDeck := deal(cards, 5)

// 	fmt.Println("hand")
// 	hand.print()
// 	fmt.Println("")
// 	fmt.Println("remaining deck")
// 	remainDeck.print()

// // function example
// func newCard() string {
// 	return "Five of card"
// }
