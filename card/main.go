package main

// import "fmt"

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

	deal(cards,2)
	// cards.print()

}
// func newCard() string{
// 	return "five of diamonds"
// }