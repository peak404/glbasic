package main

import(
	"fmt"
	"io"
	"log"
	"net"
)

func main(){
	li, err :=net.Listen("tcp",":8080")

	if err != nil{
		log.Fatalln(err)
	}
	defer li.Close()


	for{
		conn, err :=li.Accept()
		if err!=nil{
			log.Println(err)
			continue
		}

		io.WriteString(conn,"\nhello fromm TCP server\n")
		fmt.Fprintln(conn,"hows your date")
		fmt.Fprintf(conn,"very very good")
		conn.Close()

	}

}
	


