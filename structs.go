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
	AccDate   time.Time
	TxDate    time.Time
	Place     string
	Category  string
	Shop      string
	AccAmount Amount
	TxAmount  Amount
}

type Amount struct {
	Amount   float64
	Currency string
}

func (r *Record) OneCurrency() bool {
	return r.AccAmount.Currency == r.TxAmount.Currency
}

func (r *Record) TxAmountS() string {
	if r.OneCurrency() {
		return ""
	} else {
		return r.TxAmount.String()
	}
}

func (r *Record) AccDateS() string {
	return r.AccDate.Format("2006-01-02")
}

func (r *Record) TxDateS() string {
	if r.TxDate.IsZero() {
		return ""
	} else {
		return r.TxDate.Format("2006-01-02")
	}
}

func (r *Amount) String() string {
	return fmt.Sprintf("% .2f %s", r.Amount, r.Currency)
}
