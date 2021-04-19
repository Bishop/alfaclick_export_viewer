package main

import (
	"fmt"
	"time"
)

type Database struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Fields []Field `xml:"field"`
}

type Field struct {
	Name  string `xml:"name,attr"`
	Value string `xml:",innerxml"`
}

type Records []Record

type Record struct {
	FixDate         time.Time
	OpDate          time.Time
	Place           string
	Category        string
	Shop            string
	AccountAmount   Amount
	OperationAmount Amount
}

type Amount struct {
	Amount   float64
	Currency string
}

func (r *Record) OneCurrency() bool {
	return r.AccountAmount.Currency == r.OperationAmount.Currency
}

func (r *Record) FixDateS() string {
	return r.FixDate.Format("2006-01-02")
}

func (r *Record) OpDateS() string {
	if r.OpDate.IsZero() {
		return ""
	} else {
		return r.OpDate.Format("2006-01-02")
	}
}

func (r *Amount) String() string {
	return fmt.Sprintf("% .2f %s", r.Amount, r.Currency)
}
