package main

import (

"testing"
"os"
)

func TestNewDeck(t *testing.T){
	d:= newDeck()
	// if len(d)!=16{
	// 	t.Errorf("Expected deck length of 16, but got %v", len(d))
	// }
	
	if d[0]!="ace of spades"{
		t.Errorf("not ace of spades, but got %v",d[0])
	}
	
	if d[len(d)-1]!="four of club"{
		t.Errorf("not four of club, but got %v",d[len(d)-1])
	}


}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck:=newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck :=newDeckFromFile("_decktesting")

	if len(loadedDeck)!=16{
		t.Errorf("expected 16 cards,got %v",len(loadedDeck))
	}
	os.Remove("_decktesting")
}






