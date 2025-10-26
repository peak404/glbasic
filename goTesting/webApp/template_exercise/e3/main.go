package main

import(
	"os"
	"log"
	"text/template"
)

type mealInfo struct{
	Name string
	Price float64
}


type mealType struct{
	Type string
	MealInfo []mealInfo

}


var tmp *template.Template
func init(){
	tmp = template.Must(template.ParseFiles("tpl2.gohtml"))
}


func main(){
meals := []mealType{
	{
		Type :"breakfast",
		MealInfo:[]mealInfo{
			{"banana bread", 3.99},
			{"egg muffin", 2.45},
		},
	},
	{
		Type: "lunch",
		MealInfo :[]mealInfo{
			{"chicken rice",10.99},
			{"beef noodesoup",12.99},
		},
	},
	{
		Type: "dinner",
		MealInfo:[]mealInfo{
			{"supreme pizza",34.99},
			{"four season",22.99},
		},
	},
}
err:=tmp.Execute(os.Stdout,meals)
if err!=nil{
	log.Fatalln(err)
}




}


