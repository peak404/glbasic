package main

import (
	"fmt"
	
	"net/http"
	"strings"
	"os"
	
)



func main(){

	c :=make(chan string)

	links :=[]string{
		"http://www.nfl.com",
		"http://www.amazon.com",
		"http://www.apple.com",
		"http://nba.com",

	}
	for _,link:=range links{
		
		
		go stringCheck([]string{link},c)
		go checkLink(link,c)
		}


	for i:=0; i<len(links)*2;i++{
		msg:=<-c
		fmt.Println(msg)

		
	}	
	}
	
	

	// stringCheck(links)





func stringCheck (urls []string,c chan string){
	
	if len(urls) ==0 {
		return 
	}
	for _,url:=range urls{
		if (strings.HasPrefix(url,"http://") ||strings.HasPrefix(url,"https://"))&&strings.HasSuffix(url,".com"){
			// fmt.Println(url)
			c <- url
			
		}else{
			// fmt.Println("error, must start with \"https:// or http\" or not endWith .com or contains b")
			c<-"error, must start with \"https:// or http\" or not endWith .com or contains b"
			os.Exit(1)
		}

	}
	


	
	}
func checkLink(link string,c chan string){

	
	resp,err :=http.Get(link)
	if err!=nil{
		// fmt.Println("error",err)
		c<-link+"might be down"
		return
	}
	c<-link
	fmt.Println(resp.Status)
	
}