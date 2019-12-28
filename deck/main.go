package main

func main() {
	cards := deck{"Ace of Diamonda", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {

	return "Five of Diamonds"

}
