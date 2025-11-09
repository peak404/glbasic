package main


import(
	"fmt"
	"sort"
)


func main(){
	nums :=[]int{-4,-1,0,3,10}
	fmt.Println(sortSquare(nums))
	

}

func sortSquare(nums []int)[]int{
	newNums :=[]int{}
	for _,v := range  nums{
		// newNums+=v*2 slice cant do +=
		newNums = append(newNums,v*v)

	}
	sort.Ints(newNums)
	return newNums
}