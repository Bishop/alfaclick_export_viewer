package main

import (
	"encoding/xml"
	"log"
	"os"
	"text/template"
)

func main() {
	data := new(Database)

	readXmlData("example.xml", data)

	records := recordsFromItems(data)

	printRecords(records, "display.go.tmpl")
}

func readXmlData(filename string, data interface{}) {
	content, err := os.ReadFile(filename)

	fatalOnError(err)

	err = xml.Unmarshal(content, data)

	fatalOnError(err)

	return
}

func printRecords(records *Records, templateName string) {
	tmpl, err := template.New(templateName).ParseFiles(templateName)

	fatalOnError(err)

	err = tmpl.Execute(os.Stdout, records)

	fatalOnError(err)
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
