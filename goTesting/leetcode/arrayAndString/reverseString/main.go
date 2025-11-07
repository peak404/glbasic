package main

import(
	"fmt"
)

func main(){

	fmt.Println(reverse("hello"))
}

func reverse(s string) string{
	r :=[]rune(s)
	i:=0
	j:=len(s)-1

	for i<j{
		temp:=r[i]
		r[i] = r[j]
		r[j]=temp

		i++
		j--
	}
	return string(r)
	


}
