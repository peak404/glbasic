// package x

// import(
// 	"log"
// 	"os"
// 	"text/template"
// )
// var tpl *template.Template

// type Hotel struct{
// 	Name string
// 	Address string
// 	City string 
// 	Zip string 
// 	Region string
// } 

// type hotel_CA struct{
// 	Hotels []Hotel

// }


// func init(){
// 	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
// }

// func main(){
// 	hotelData:=hotel_CA{
// 		Hotels:[]Hotel{
// 			{Name: "South Hotel", Address: "123 South Ave", City: "South City", Zip: "S123", Region: "Southern"},
// 			{Name: "Central Hotel", Address: "123 Central Ave", City: "Central City", Zip: "C123", Region: "Central"},
// 			{Name: "North Hotel", Address: "123 North Ave", City: "North City", Zip: "N123", Region: "Northern"},
// 		},

// 	}

// 	err:= tpl.Execute(os.Stdout,hotelData)
// 	if err != nil{
// 		log.Fatalln(err)
// 	}
	

// }
