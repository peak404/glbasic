package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Record struct{
	Date time.Time
	Values []float64
}

func main(){
	http.HandleFunc("/",foo)
	http.ListenAndServe(":8080",nil)
}

func foo(res http.ResponseWriter, req *http.Request){
	records:=prs("table.csv")
	tpl,err:=template.ParseFiles("hw.gohtml")
	if err !=nil{
		log.Fatalln(err)
	}

	err =tpl.Execute(res,records)
	if err!=nil{
		log.Fatalln(err)
	}
}

func prs(files string )[]Record{
	src,err :=os.Open(files)
	if err !=nil{
		log.Fatalln(err)
	}
	defer src.Close()
	
	readData:=csv.NewReader(src)
	rows, err:= readData.ReadAll()
	if err != nil{
		log.Fatalln(err)
	}
	
	records :=make([]Record, 0, len(rows))

	for i, row :=range rows{
		if i ==0{
			continue
		}
		date,_ :=time.Parse("2006-01-02",row[0])

		values :=make([]float64,0,len(row)-1)
		for j :=1; j<len(row);j++{
			data,_:=strconv.ParseFloat(row[j],64)
			values=append(values,data)
		}
		


		records = append(records, Record{Date:date,Values:values,})
	}
	return records
	
	
}





















// type Record struct {
// 	Date time.Time
// 	Open float64
// }

// func main() {
// 	http.HandleFunc("/", foo)
// 	http.ListenAndServe(":8080", nil)

// }

// func foo(res http.ResponseWriter, req *http.Request) {  
// 	records := prs("table.csv") //函数 处理records

// 	tpl, err := template.ParseFiles("hw.gohtml") //解析面板
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	err = tpl.Execute(res, records) //把records传入tpl面板
// 	if err != nil {
// 		log.Fatalln(err)

// 	}

// }

// func prs(filePath string) []Record {
// 	src, err := os.Open(filePath) //open table cvs file
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer src.Close()

// 	rdr := csv.NewReader(src) //create reader and read
// 	rows, err := rdr.ReadAll()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	records := make([]Record, 0, len(rows)) //create records slice to save the data

// 	for i, row := range rows {
// 		if i == 0 {
// 			continue
// 		}
// 		date, _ := time.Parse("2006-01-02", row[0])
// 		open, _ := strconv.ParseFloat(row[1], 64)

// 		records = append(records, Record{Date: date, Open: open})

// 	}
// 	return records
// }
