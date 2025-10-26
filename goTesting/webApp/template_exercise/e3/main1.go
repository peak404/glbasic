// package main1

// import(
// 	"log"
// 	"os"
// 	"text/template"
// )

// type mealInfo struct{
// 	TypeMeal string
// 	Name string
// 	Price float64
// }


// type meal struct{
// 	MealInfo  [] mealInfo
// }

// type Meal []meal


// var tpl *template.Template
// func init(){
// 	tpl = template.Must(template.ParseFiles("tpl1.gohtml"))
// }
// func main(){
// 	meals :=Meal{
// 		meal{
// 			MealInfo:[]mealInfo{
// 				{
// 					TypeMeal: "breakfast",
// 					Name: "egg bacon oatmeal",
// 					Price: 7.99,
// 				},
// 				{
// 					TypeMeal:"lunch",
// 					Name: "turkey rice",
// 					Price: 14.99,
// 				},
// 				{
// 					TypeMeal:"dinner",
// 					Name:"curry chicken meal",
// 					Price: 18.99,
// 				},

// 			},
// 		},
// 	}




// 	err :=tpl.Execute(os.Stdout,meals)
// 	if err != nil{
// 		log.Fatalln(err)
// 	}

// }