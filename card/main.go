package main

import (
	
	"fmt"
)

func main(){
	// var card string = "spades"
	// card:="spades"
	//cards :=[]string{"ace of diamonds", newCard()}
	// cards :=deck{"ace of diamonds", newCard()}
	// cards = append(cards,"six of spades")
	
	// for i, card :=range cards{
	// 	fmt.Println(i,card)
	// }
	cards:=newDeck()
	fmt.Println(cards.toString())

	// hand,remain:=deal(cards,15)
	// hand.print()
	// remain.print()

	// cards.print()

	// i:="hello world"
	// j:=[]byte(i);
	// fmt.Println(string(j))
	

}
// func newCard() string{
// 	return "five of diamonds"
// }