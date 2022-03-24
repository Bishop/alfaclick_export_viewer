package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func recordsFromItems(data *Database) *Records {
	records := make(Records, len(data.Items))

	for i, item := range data.Items {
		records[i] = makeRecord(&item.Fields)
	}

	return &records
}

func makeRecord(fields *[]Field) Record {
	record := Record{}

	for _, field := range *fields {
		switch field.Name {
		case "Дата",
			"Date":
			record.AccDate = parseFixDate(field.Value)
		case "Сумма в валюте счета",
			"Amount in account currency":
			record.AccAmount = parseAmount(field.Value)
		case "Сумма в валюте операции",
			"Amount in transaction currency":
			record.TxAmount = parseAmount(field.Value)
		case "Примечание",
			"Note":
			var opDate string

			opDate, record.Place, record.Category, record.Shop = parseNotes(field.Value)

			if opDate != "" {
				record.TxDate = parseOpDate(opDate)
			}
		}
	}

	return record
}

func parseFixDate(s string) time.Time {
	t, err := time.Parse("02.01.2006", s)

	fatalOnError(err)

	return t
}

func parseOpDate(s string) time.Time {
	t, err := time.Parse("20060102", s)

	fatalOnError(err)

	return t
}

func parseAmount(s string) Amount {
	amountRegExp := regexp.MustCompile(`(.*?)\s(\w{3}$)`)

	result := amountRegExp.FindStringSubmatch(s)

	return Amount{parseNumber(result[1]), result[2]}
}

func parseNumber(s string) float64 {
	normalization := map[string]string{" ": "", " ": "", ",": "."}

	for old, replace := range normalization {
		s = strings.ReplaceAll(s, old, replace)
	}

	amount, err := strconv.ParseFloat(s, 64)

	fatalOnError(err)

	return amount
}

func parseNotes(s string) (string, string, string, string) {
	notesRegExp := regexp.MustCompile(`(\d{8})\s(.*?)\s{2,}(.*?)\s{2,}(.*)`)

	result := notesRegExp.FindStringSubmatch(s)

	if len(result) == 5 {
		return result[1], result[2], result[3], result[4]
	} else {
		return "", s, "", ""
	}
}

func sortRecordsByDate(records []Record) {
	sort.Slice(records, func(i, j int) bool {
		return records[j].Date().Before(records[i].Date())
	})
}
