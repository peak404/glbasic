package main


import(
	"fmt"
	"sort"
)

func main(){

	nums:=[]int{-4,-1,0,3,10}
	fmt.Println(sortSquare(nums))

}

func sortSquare(num [] int)[]int{
	newNum :=[]int{}

	for _,v :=range num{
		newNum= append(newNum,v*v)
	}
	sort.Ints(newNum)
	return newNum
}

