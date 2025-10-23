// package main



// import(
// 	"os"
// 	"log"
// 	"text/template"
// )

// var tpl *template.Template
// func init(){
// 	tpl = template.Must(template.ParseFiles("tpl2.gohtml"))
// }

// type data struct{
// 	Name string
// 	Address string
// 	Region string
// }

// type hotels []data 

// func main(){

// 	h:=hotels{
// 		data{
// 			Name: "South hotel",
// 			Address:"S ave",
// 			Region: "South",
// 		},
// 		data{
// 			Name: "East hotel",
// 			Address: "E ave",
// 			Region: "East",
// 		},
// 	}

// 	err:=tpl.Execute(os.Stdout,h)
// 	if err!=nil{
// 		log.Fatalln(err)
// 	}

// }