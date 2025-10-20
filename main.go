package main

import (
	"log"
	"os"
	"string"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct{
	Name string
	Word strings
}

type car struct{
	Manufacturer string
	Model strings
	Door int
}
	
	
var fm = template.FuncMap{
	"uc":strings.ToUpper,
	"ft":firstThree,

}

func init(){
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string{
	s=strings.TrimSpace(s)
	if len(s)>=3{
		s=s[:3]
	}
	return s
}



func main()
{
	b:=sage{
		Name: "budda"
		Word: "peace forever"
	}

	s:sage{
		Name : "shen"
		word : "happy happy happy"
	}

	f:car{
		Manufacturer:"ford"
		Model:"fx350"
		Door:"4"
	}

	t:car{
		Manufacturer:"toyota"
		Model:"avalon666"
	}

	sages:=[]sage{b,s}
	cars:=[]car{f,t}

	datas:=struct{
		wisdom []sage
		transport []car
	}{
		sages
		cars
	}

	err :=tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml",datas)
	if err!=nil{
		log.Fatalln(err)
	}

}
