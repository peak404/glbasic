package main

import (
	// "bytes"
	// "fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
func init(){	
	tpl = template.Must(template.ParseGlob("4-parse/*"))

}

func main() {

	err :=tpl.Execute(os.Stdout,nil)
	if err !=nil{
		log.Fatalln(err)
	}

	err=tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",nil)
		if err !=nil{
			log.Fatalln(err)
		}
	
	err=tpl.ExecuteTemplate(os.Stdout,"tpl2.gohtml",`无语了`)
		if err !=nil{
			log.Fatalln(err)
		}
	
}
	// tpl, err := template.ParseFiles("tpl.gohtml")
	// if err != nil {
	// 	log.Fatalln("error", err)
	// }
	// err = tpl.Execute(os.Stdout, nil)   //cant modify
	// if err != nil {
	// 	log.Fatalln("error", err)
	// }

	
	// fmt.Println("\n--------------------------------")

	// var buf bytes.Buffer       //can be modify,
	// err = tpl.Execute(&buf, nil)
	// fmt.Println(buf.String())


