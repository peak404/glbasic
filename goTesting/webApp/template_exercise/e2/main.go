package main

import(
	"log"
	"os"
	"text/template"
)
var tpl *template.Template

type hotel struct{
	Name string
	Address string
	City string 
	Zip string 
	Region string
} 

type hotel_CA struct{
	hotels []hotel

}










func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

