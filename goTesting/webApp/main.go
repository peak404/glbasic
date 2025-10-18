package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name:=os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	str :=fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
		</head>
	<body>the name is = `+name +`
    
	</body>
	</html>
	
	
	
	
	`)

	

	i, err := os.Create("index.html")
	
	if err != nil {
		log.Fatal("error", err)

	}
	defer i.Close()
	s,err:=io.Copy(i, strings.NewReader(str))
	fmt.Println(s)
	
}
