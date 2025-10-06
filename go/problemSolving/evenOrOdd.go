package main



import (
    "fmt"
    "math/rand"
    "time"
)


func randomNum(op string) float64{
	source :=rand.NewSource(time.Now().Unix())
	r:=rand.New(source)
	x:=r.Intn(1000)+1
	y:=r.Intn(1000)+1

	var result float64
	switch op{
	case "+":
		result = float64(x+y)
	case "-":
		result = float64(x-y)
	case "*":
		result = float64(x*y)
	case "/":
		if y!=0{
		result = float64(x/y)
		}else{
			fmt.Println("zero is not allowed")
			return 0
		}
	default:
		fmt.Println("must be op")
		return 0
	}
	fmt.Println(result)
	return result
}
    func evenOrOdd(n float64){
		num:=int(n)
		if num%2 == 0{
			fmt.Println("even")
		}else{
			fmt.Println("odd")
		}
		
	}


func main(){
	result:= randomNum("+")
	evenOrOdd(result)
}


















