package main
import(
	"fmt"
	"os"
)


func main(){
	data,err :=os.ReadFile("text1.txt")
	if err != nil{
		fmt.Println("something wrong",err)
		return
	}
	fmt.Println(string(data))

}