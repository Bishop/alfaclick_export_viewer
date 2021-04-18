package main

import "time"

type Database struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Fields []Field `xml:"field"`
}

type Field struct {
	Name string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}

type Records []Record

type Record struct {
	FixDate time.Time
	OpDate time.Time
	Place string
	Category string
	Shop string
	AccountAmount Amount
	OperationAmount Amount
}

type Amount struct {
	Amount float64
	Currency string
}