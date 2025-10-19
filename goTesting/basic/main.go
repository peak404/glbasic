package main
import (
	"fmt"
	"time"
	"math/rand"

	
)

func main(){
	//print emoji
	fmt.Println("😂")
	a1,b1,c1,d1,e1 :=0,1,2,3,4

	//binary hex
	fmt.Printf("%v \t %b \t %X\n",a1,a1,a1)
	fmt.Printf("%v \t %b \t %X\n",b1,b1,b1)
	fmt.Printf("%v \t %b \t %X\n",c1,c1,c1)
	fmt.Printf("%v \t %b \t %X\n",d1,d1,d1)
	fmt.Printf("%v \t %b \t %X\n",e1,e1,e1)

	fmt.Println(time.Now())
	fmt.Println(rand.Intn(10))

	//var y int
	a:=40

	if i:=2*rand.Intn(a); i>=a{
		fmt.Printf("%d is greater than or equal to %d\n",i,a)
	}else{
		fmt.Printf("%d is less than %d\n",i,a)
	}

	fmt.Println(colorZone(5))





	ch1 := make(chan int)
	ch2 := make(chan int)

	// get an int64, convert it to type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type duration's underlying type is int64
	// time.Sleep takes type duration
	// time.Millisecond is a constant in the time package
	// https://pkg.go.dev/time#pkg-constants
	d1 := time.Duration(rand.Int63n(250))

	fmt.Println(d1)
	d2 := time.Duration(rand.Int63n(250))
	// fmt.Printf("%v \t %T\n", d1, d1)
	// time.Sleep(d1 * time.Millisecond)
	// fmt.Println("Done")

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	// https://go.dev/ref/spec#Select_statements
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

}
var colors=map[int]string{
	1:"red",
	2:"green",
	3:"blue",
}


func colorZone(offset int)string{
	if color,ok :=colors[offset]; ok{
		return color
	}
	return "error"

	
}

