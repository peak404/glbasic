package main

import(
	"log"
	"os"
	"text/template"
)

type breakfast struct{

}

type


var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseFiles("tpl1.gohtml"))
}
func main(){

}