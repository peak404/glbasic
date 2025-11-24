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

// func mySort(nums [] int)[]int{
// 	//newNum :=[]int{}

// 	size:=len(nums)

// 	for i:=0; i<size-1;i++{
// 		minIndex:=i
// 		for j:=1;j<size;j++{
// 			if nums[j]<nums[minIndex]{
// 				minIndex= j
// 			}
// 		}



// 	}
	
	


// }