package main


import(
	"fmt"
)

func main(){

	fmt.Println(isSubsequence("abc","asgrc"))


}


func isSubsequence(s string, t string) bool{
	left,right :=0,0
	for left<len(s) && right<len(t){
		if s[left] == t[right]{
			left++
			right++
		}else{
			right++
		}
	}
	if left == len(s){
			return true
		}
	return false

	


}

