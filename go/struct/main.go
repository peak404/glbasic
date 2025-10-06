package main
import(
	"fmt"
)

type contactInfo struct{
	email string
	zipCode int

}
type person struct{
	firstName string
	lastName string
	contactInfo
}

func main(){
	// a :=person{"shen","qi"}

	// var p person
	// p.firstName = "dandan"
	// p.lastName = "lin"

	p:=person{
		firstName:"dandan",
		lastName:"lin",
		contactInfo: contactInfo{
			email:"www.ddl.com",zipCode:11111,
		},
	}
	fmt.Printf("%+v",p)
	
	
} 
