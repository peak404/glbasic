// package testing1

// import (
// 	"fmt"
// 	"net/http"

// )



//  func main(){
// 	links:=[]string{
// 		"http://google.com",
// 		"http://facebook.com",
// 		"http://stackoverflow.com",
// 		"http://golang.org",
// 		"http://amazon.com",

// 	}
	 
// 	c:=make(chan string)


// 	for _,link:=range links{
// 		 go checkLink(link,c)
// 	}

// 	// for i:= 0;i<len(links);i++{
// 	// 	fmt.Println(<-c)
		
// 	// }
// 	for{
// 		go checkLink(<-c,c)
// 	}
//  }

//  func checkLink(link string,c chan string){
// 	_, err := http.Get(link)
// 	if err!=nil{
// 		c <- link + " might be down!"
// 		return
// 	}
// 	fmt.Println(link+" is up")
// 	c<-link
// 	// fmt.Println(resp.Status)
//  }