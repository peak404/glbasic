package main



import(
	"os"
	"io"
	"log"
	"net"
)

func main(){
	li,err :=net.Listen("tcp",":8080")
	if err !=nil{
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err !=li.Accept()
		log.Println(err)
		continue
	}
	go handle(conn)

	}


	func handle(conn,net.Conn){    //net.Conn is type
		scanner :=bufio.NewScanner(conn) 	//Wraps the connection in a Scanner — this reads text input line by line.
		
		for scanner.Scan(){
			ln:=scanner.Text()
			fmt.Printlb(ln)
		}

		defer conn.Close()
		// •	This loop reads each line sent by the client.
		// •	scanner.Scan() returns true if a new line is read successfully.
		// •	scanner.Text() gives you the text of that line.
		// •	fmt.Println(ln) just prints that line on the server console (not to the client).

	}