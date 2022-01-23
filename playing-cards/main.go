package main

import "fmt"

// to run playing-cards, type go run main.go deck.go

func main() {
	cards:= newDeck()
	cards.print()

	fmt.Println("dealing 5 cards now")
	dealedCards, remainingCards := deal(cards, 5)
	dealedCards.print()
	remainingCards.print()

	fmt.Println("dealing 3 cards now")
	dealedCards, remainingCards = deal(remainingCards, 3)
	dealedCards.print()
	remainingCards.print()

	remainingCards.writeToFile("savedDeck")

	newlyLoadedDeck := readDeckFromFile("savedDeck")
	fmt.Println("before shuffle")
	newlyLoadedDeck.print()

	newlyLoadedDeck.shuffle()
	fmt.Println("after shuffle")
	newlyLoadedDeck.print()
}

