package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	cards := newDeck()
	// cards := newDeckFromFile("my_cards")

	cards.shuffleDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	cards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
