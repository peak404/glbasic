package main

import(
	"fmt"
	"os"
	"net/http"
)

func main(){
	resp,err :=http.Get("http://www.google.com")
	if err !=nil{
		fmt.Println("error: ",err)
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Println(resp.Status)

	

}