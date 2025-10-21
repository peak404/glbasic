package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct{
	Name string
	Word string
}

type car struct{
	Manufacturer string
	Model string
	Door int
}
	
	
var fm = template.FuncMap{
	"uc":strings.ToUpper,
	"ft":firstThree,

}

func init(){
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("8-function/tpl.gohtml"))
}

func firstThree(s string) string{
	s=strings.TrimSpace(s)
	if len(s)>=3{
		s=s[:3]
	}
	return s
}



func main(){
	b:=sage{
		Name: "budda",
		Word: "peace forever",
	}

	s:=sage{
		Name : "shen",
		Word : "happy happy happy",
	}

	f:=car{
		Manufacturer:"ford",
		Model:"fx350",
		Door: 4,
	}

	t:=car{
		Manufacturer:"toyota",
		Model:"avalon666",
	}

	sages:=[]sage{b,s}
	cars:=[]car{f,t}

	datas:=struct{
		Wisdom []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err :=tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",datas)
	if err!=nil{
		log.Fatalln(err)
	}

}
