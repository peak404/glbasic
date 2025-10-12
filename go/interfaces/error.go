// package main
// import (
// 	"fmt"
	
// )

// type argError struct{
// 	arg int
// 	msg string
// }

// func(err *argError) Error() string{
// 	return fmt.Sprintf("%d ->%s",err.arg,err.msg)
// }
// func f1(num int)(int, error){
// 	if num<0{
// 		return -1,&argError{num,"cant be 0 or negitive number"}
// 	}
// 	return 3 *num, nil
// }


// func main(){
// 	result,err :=f1(-10)
// 	fmt.Println(result,err)
// }