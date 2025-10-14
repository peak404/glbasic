// package main

// import(
// 	"fmt"
// 	"os"
// 	"net/http"
// 	"io"
// )














// type shape interface{
// 	getArea() float64
// }

// type triangle struct{
// 	height,base float64
// }

// type square struct{
// 	 sideLength float64
// }
//  func ( t triangle) getArea() float64{

// 	return 0.5*t.base*t.height

//  }

//  func (s square) getArea() float64{
// 	return s.sideLength*s.sideLength
//  }


// func printArea(s shape){
// 	fmt.Println(s.getArea())

// }



// type logWriter struct{

// }
// func main(){










// 	t:=triangle{base:10,height:10}
// 	s:=square{sideLength: 10}
// 	printArea(t)
// 	printArea(s)









// 	resp,err :=http.Get("http://www.google.com")
// 	if err !=nil{ 
// 		fmt.Println("error: ",err)
// 		os.Exit(1)
// 	}
	
// 	// fmt.Println(resp.Status)

// 	// bs:=make([]byte,99999)
// 	// resp.Body.Read(bs)
// 	// fmt.Println(string(bs))

// 	lw:=logWriter{}
// 	// io.Copy(os.Stdout,resp.Body)
// 	io.Copy(lw,resp.Body)

// }
// func (logWriter) Write(bs []byte) (int,error){
// 	fmt.Println(string(bs))
// 	fmt.Println("->>>>>>>>>>>",len(bs))

	
// 	return len(string(bs)), nil

	
// }


// package main

// import "fmt"

// type rect struct {
// 	width, height int
// }

// // ✅ Value receiver — works on a COPY
// func (r rect) doubleValueReceiver() {
// 	r.width *= 2
// 	r.height *= 2
// 	fmt.Println("Inside doubleValueReceiver:", r)
// }

// // ✅ Pointer receiver — works on the ORIGINAL
// func (r *rect) doublePointerReceiver() {
// 	r.width *= 2
// 	r.height *= 2
// 	fmt.Println("Inside doublePointerReceiver:", *r)
// }

// func main() {
// 	r1 := rect{10, 5}

// 	fmt.Println("Before value receiver:", r1)
// 	r1.doubleValueReceiver()
// 	fmt.Println("After value receiver:", r1) // ❌ unchanged (copy modified)

// 	fmt.Println("\nBefore pointer receiver:", r1)
// 	r1.doublePointerReceiver()
// 	fmt.Println("After pointer receiver:", r1) // ✅ changed (original modified)
// }


// Before value receiver: {10 5}
// Inside doubleValueReceiver: {20 10}
// After value receiver: {10 5}          // unchanged — worked on a COPY

// Before pointer receiver: {10 5}
// Inside doublePointerReceiver: {20 10}
// After pointer receiver: {20 10}       // changed — worked on the ORIGINAL