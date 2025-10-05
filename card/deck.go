package main

//create a new type of  'deck'
//which is a slice of strings
import (
	"fmt"
	"strings"
	"os"
	"math/rand"
	"time"
)
type deck []string

func newDeck() deck { 
	cards := deck{}
	cardSuit := []string{"spades", "diamonds", "heart", "club"}
	cardValue := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
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

func (d deck) saveToFile(filename string) error{
	return os.WriteFile(filename,[]byte(d.toString()),0666)
}

func newDeckFromFile(filename string) deck{
	bts,err:=os.ReadFile((filename))
	if err!=nil{
		fmt.Println("error:",err)
		os.Exit(1)
	}

	arr:=strings.Split(string(bts),",")
	// for i:=range arr{
	// 	arr[i] =strings.TrimSpace(arr[i])
	// }	
	return deck(arr)
}

func (d deck)shuffle(){

		source:=rand.NewSource(time.Now().UnixNano())
		r:=rand.New(source)
	    for i:=range d{	
		 
		 newPosition:=r.Intn(len(d)-1)
		 d[i],d[newPosition] = d[newPosition],d[i]

	}

}

