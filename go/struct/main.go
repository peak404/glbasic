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
	// fmt.Printf("%+v",p)
	// pointerForP:=&p
	p.updateName("dd")
	p.print()
	
} 
func(p1 *person)updateName(newFirstName string){
	(*p1).firstName =  newFirstName
}
func (p1 person) print(){
	fmt.Printf("%+v",p1)
}

