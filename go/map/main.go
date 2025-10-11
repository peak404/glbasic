package main

import "fmt"

func main(){
	colors := map[string]string{
		"red":"#ff0000",
		"green":"#f3434",
		"yellow":"#f34fff",
	}
	 
	// var colors map[string]string

	// colors :=make(map[int]string)

	// colors[3]="#fff3fds"
	// delete(colors,3)




	
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(color map[string]string){
	for i,v:=range color{
		fmt.Printf("%s  <-> %s\n",i,v) 
	}
}