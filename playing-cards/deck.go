package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

// receiver function
func (d deck) print() {
	for index, card:=range(d) {
		fmt.Println(index,":",card)
	}
}

// dont need input, so dont need receiver
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts","Clubs","Spades","Diamonds",}
	cardNumbers := []string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King",}
	for _,cardSuit := range(cardSuits) {
		for _,cardNumber := range(cardNumbers) {
			cards = append(cards, cardNumber+" of "+cardSuit)
		}
	}
	return cards
}

func deal(d deck, cardCount int) (deck, deck) {
	cutOffPoint := len(d)-cardCount
	dealedCards := d[cutOffPoint:]
	remainingCards := d[:cutOffPoint]
	return dealedCards, remainingCards

}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) writeToFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0777)
}

func readDeckFromFile(filename string) deck {
	fileInByte, err:= ioutil.ReadFile(filename)

	if (err != nil) {
		fmt.Println("Error: ", err)
		os.Exit((1))
	}
	fileInString := string(fileInByte)
	return strings.Split(fileInString, ",")
}

func (d deck) shuffle() deck {
	for i := range d {
		newPosition := rand.Intn(len(d)-1)
		saveThisCard := d[newPosition]
		d[newPosition] = d[i]
		d[i] = saveThisCard
	}

	return d
}