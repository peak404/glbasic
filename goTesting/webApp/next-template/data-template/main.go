package main

import(
	"os"
	"log"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout,"index.gohtml",100)
	if err != nil{
		log.Fatalln(err)
	}

}