package main



//never use this unless its custom 
import(
	"fmt"
	"io"
	"log"
	"net"
)

func main(){
	li, err :=net.Listen("tcp",":8080")  //发起tcp链接

	if err != nil{
		log.Panic(err)
	}
	defer li.Close()


	for{
		conn, err :=li.Accept()  //accepct return new connection obj conn
		if err!=nil{
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nhello fromm TCP server\n")   //io.WriteString 直接把字符串转为字节并调用 conn.Write。
		fmt.Fprintln(conn, "hows your date")
		fmt.Fprintf(conn, "%v\n","very very good")
		fmt.Println(conn.RemoteAddr())

		conn.Close()

	}

}
	


