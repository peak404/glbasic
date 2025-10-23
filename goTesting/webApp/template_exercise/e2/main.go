package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl3.gohtml"))
}

type hotelAtt struct {
	Name, Address, Region string
}
type region struct {
	Region string
	Hotels []hotelAtt
}

type Regions []region

func main() {
	h := Regions{
		region{
			Region: "East",
			Hotels: []hotelAtt{
				hotelAtt{
					Name:    "east hotel",
					Address: "east ave",
					Region:  "east",
				},
				hotelAtt{
					Name:    "e hotel",
					Address: "e ave",
					Region:  "east",
				},
			},
		},
		region{
			Region: "West",
			Hotels: []hotelAtt{
				hotelAtt{
					Name:    "west hotel",
					Address: "west ave",
					Region:  "west",
				},
				hotelAtt{
					Name:    "w hotel",
					Address: "w ave",
					Region:  "west",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
