package main

//user input reader scanner

import(
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	li,err :=net.Listen("tcp",":8080")
	if err !=nil{
		log.Fatalln(err)
	}
	defer li.Close()

	for {                       
		conn, err :=li.Accept()   //accepct return new connection obj conn
		if err !=nil{
		log.Println(err)
		continue
		}
		go handle(conn)          //conn is type net.Conn
	}


}


	func handle(conn net.Conn){    //net.Conn is type
		err :=conn.SetDeadline(time.Now().Add(10 *time.Second))  //set time when it will be done for 10 sec
		if err !=nil{
			log.Fatalln("conn timeout")
		}
		scanner :=bufio.NewScanner(conn) 	//Wraps the connection in a Scanner — this reads text input line by line.
		
		for scanner.Scan(){
			ln:=scanner.Text()
			fmt.Println(ln)
		}

		defer conn.Close()
		// •	This loop reads each line sent by the client.
		// •	scanner.Scan() returns true if a new line is read successfully.
		// •	scanner.Text() gives you the text of that line.
		// •	fmt.Println(ln) just prints that line on the server console (not to the client).

		//we never get here
		//how doe reader know when it is done
		//we set the conn.SetDeadline

		fmt.Println("code got here")
	}