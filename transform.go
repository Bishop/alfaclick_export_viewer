package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func recordsFromItems(data *Database) *Records {
	records := make(Records, len(data.Items))

	for i, item := range data.Items {
		records[i] = *makeRecord(&item.Fields)
	}

	return &records
}

func makeRecord(fields *[]Field) *Record {
	record := Record{}

	for _, field := range *fields {
		switch field.Name {
		case "Дата",
			"Date":
			record.FixDate = parseFixDate(field.Value)
		case "Сумма в валюте счета",
			"Amount in account currency":
			record.AccountAmount = parseAmount(field.Value)
		case "Сумма в валюте операции",
			"Amount in transaction currency":
			record.OperationAmount = parseAmount(field.Value)
		case "Примечание",
			"Note":
			var opDate string

			opDate, record.Place, record.Category, record.Shop = parseNotes(field.Value)

			if opDate != "" {
				record.OpDate = parseOpDate(opDate)
			}
		}
	}

	return &record
}

func parseFixDate(s string) time.Time {
	t, err := time.Parse("02.01.2006", s)

	if err != nil {
		log.Fatal(err)
	}

	return t
}

func parseOpDate(s string) time.Time {
	t, err := time.Parse("20060102", s)

	if err != nil {
		log.Fatal(err)
	}

	return t
}

func parseAmount(s string) Amount {
	amountRegExp := regexp.MustCompile(`(.*?)\s(\w{3}$)`)

	result := amountRegExp.FindStringSubmatch(s)

	return Amount{parseNumber(result[1]), result[2]}
}

func parseNumber(s string) float64 {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", ".", -1)

	amount, err := strconv.ParseFloat(s, 64)

	if err != nil {
		log.Fatal(err)
	}

	return amount
}

func parseNotes(s string) (string, string, string, string) {
	// 20210413 moskva                     Покупка товара / получение услуг    Y.M mif

	notesRegExp := regexp.MustCompile(`(\d{8})\s(.*?)\s{2,}(.*?)\s{2,}(.*)`)

	result := notesRegExp.FindStringSubmatch(s)

	if len(result) == 5 {
		return result[1], result[2], result[3], result[4]
	} else {
		return "", s, "", ""
	}
}
