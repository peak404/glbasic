package main

import(
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseFiles("7-data-structure/tpl.gohtml"))
}

type sage struct{
	Name string
	word string
}

func main(){
	// file :=[]string{"a", "b","c","d","e","f"}

	file2 :=map[string]string{
		"A":"a",
		"B":"b",
		"c":"c",
	}

	// budda:=sage{
	// 	Name: "budda",
	// 	word: "peace forever",
	//}
	err :=tpl.Execute(os.Stdout,file2)
	if err !=nil{
		log.Fatalln(err)
	}


}
