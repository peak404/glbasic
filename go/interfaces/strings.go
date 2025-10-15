package main

import (
	"fmt"
	
	"os"
	"strings"
)


func main(){
	fmt.Println("all args:",os.Args)
	file,err:=os.ReadFile("text1.txt")
	if(err != nil){
		fmt.Println("error",err)
		return
	}
	fn:=[]string{"shen"}
	ln:=[]string{"qi"}
	age:=99
	fn = append(fn, append(ln,fmt.Sprint(age))...)
	

	for _,args:=range os.Args[1:]{
		helper:=strings.Split(args,",")
		fn = append(fn, helper...)
	}
	
	
	
		fmt.Println(fn)
	



	
	

	fmt.Println(strings.Join(fn,","),string(file))
	fmt.Println(len(fn))

	string2 :=[]string{"A","b","c"}
	fmt.Println(string2)
	fmt.Println(strings.Join(string2,","))
	
	fmt.Println(len(string2))
	

	

	
}
	


















	// firstName:=[]string{"shen"}
	// lastName:=[]string{"qi"}
	// age:=18

	// newVar := append(firstName,append(lastName,fmt.Sprint(age))...)
	// fmt.Println(newVar)
	// fmt.Println(strings.Join(newVar,","))

	// fmt.Println("all args",os.Args)
	
	// if len(os.Args)>1{		
	// 	 for i,args:=range os.Args{
	// 		fmt.Printf("Args %d: %s\n",i+1,args)
	// 	 }

	// }else{
	// 	fmt.Println("need provide args")
	// }
	



