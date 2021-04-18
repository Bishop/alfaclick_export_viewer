package main

import (
	"encoding/xml"
	"log"
	"os"
	"text/template"
)

func main() {
	println("hello")

	data := new(Database)
	readFile("example.xml", data)

	records := recordsFromItems(data)

	//fmt.Println(records)

	tmpl, _ := template.New("display.go.tmpl").ParseFiles("display.go.tmpl")
	_ = tmpl.Execute(os.Stdout, records)
}

func readFile(filename string, data interface{}) {
	content, _ := os.ReadFile(filename)

	err := xml.Unmarshal(content, data)

	if err != nil {
		log.Fatal(err)
	}

	return
}
