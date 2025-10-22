package main

import(
	"os"
	"log"
	"text/template"
)
type course struct{
	Number,Name,Units string
}
type semester struct{
	Term string
	Courses []course
}
type year struct{
	Fall,Spring,Summer semester
}


var tpl*template.Template


func init(){
		tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){

	y:=year{
		Fall:semester{
			Term: "Fall Time",
			Courses:[]course{
				course{"cse100","intro to go","4"},
				course{"cse200","ati 200","4"},
				course{"cse300","mta 400","4"},
			},
		},
		Spring:semester{
			Term:"Spring Time",
			Courses:[]course{
				course{"cse666","two","5"},
				course{"cse777","three","5"},
				course{"cse888","four400","5"},
			},
		},
	}






	err :=tpl.Execute(os.Stdout, y)
	if err !=nil{
		log.Fatalln(err)
	}
}
