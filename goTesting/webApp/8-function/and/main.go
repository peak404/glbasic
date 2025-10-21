package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
type user struct{
	Name string
	Word string
	Admin bool
}

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){
	u1 := user{
		Name:"dan",
		Word:"im fat",
		Admin: false,
	}
	u2 := user{
		Name:"shen",
		Word: "play for win",
		Admin: true,
	}
	u3 :=user{
		Name: "kk",
		Word: "zen",
		Admin: true,
	}

	users :=[]user{u1,u2,u3}

	err:=tpl.Execute(os.Stdout,users)
	if err !=nil{
		log.Fatalln(err)
	}



}

