package main

import (
	"fmt"
	"os"
)

// func main(){
	
// 	var s, sep string
// 	for i:=1;i<len(os.Args);i++{
// 		s+=sep+os.Args[i]
// 		sep =" - - "
// 	}
// 	fmt.Println(s)

// }



	

	// var count, k string
	// for i:=1; i<len(os.Args);i++{
	// 	count+=k+os.Args[i]
	// 	k=" "
		
	// }
	// fmt.Println(count)
	
func main(){

		
		s, sep :="", ""
		
		for
		_, arg := range os.Args[1:] {
		s += sep + arg
		sep =
		" "
		}
		fmt.Println(s)
	

}

