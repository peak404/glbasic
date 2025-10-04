package main

//create a new type of  'deck'
//which is a slice of strings
import (
	"fmt"
	"strings"
)
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuit := []string{"spades", "diamonds", "heart", "club"}
	cardValue := []string{"ace", "two", "three", "four "}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" -of- "+suit)
		}
	}

	return cards

}

func (d deck) print() {
	for i, v := range d {
		fmt.Println(i, v)
	}
}

func deal(d deck, handSize int) (deck, deck) { //type of deck, handsize of int, return two type both are deck
	return d[:handSize], d[handSize:]
}

func (d deck)toString() string {
	return strings.Join([]string(d),",")
	
} 