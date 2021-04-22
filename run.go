package main

import (
	"encoding/xml"
	"log"
	"os"
	"text/template"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	data := new(Database)

	readXmlData("example.xml", data)

	records := recordsFromItems(data)

	if false {
		printRecords(records, "display.go.tmpl")
	} else {
		createUiTable(records)
	}
}

func createUiTable(records *Records) {
	table := tview.NewTable()
	for i, record := range *records {
		table.SetCell(i, 0, &tview.TableCell{
			Text:            record.FixDateS(),
			Align:           tview.AlignLeft,
			Color:           tcell.ColorLightGray,
			BackgroundColor: tcell.ColorBlack,
		})
		table.SetCell(i, 1, &tview.TableCell{
			Text:            record.OpDateS(),
			Align:           tview.AlignLeft,
			Color:           tcell.ColorGray,
			BackgroundColor: tcell.ColorBlack,
		})

		color1 := tcell.ColorLimeGreen
		color2 := tcell.ColorLightGreen

		if record.AccountAmount.Amount >= 0 {
			color1 = tcell.ColorAquaMarine
			color2 = tcell.ColorAquaMarine
		}

		table.SetCell(i, 2, &tview.TableCell{
			Text:            record.AccountAmount.String(),
			Align:           tview.AlignLeft,
			Color:           color1,
			BackgroundColor: tcell.ColorBlack,
		})
		table.SetCell(i, 3, &tview.TableCell{
			Text:            record.OperationAmountS(),
			Align:           tview.AlignLeft,
			Color:           color2,
			BackgroundColor: tcell.ColorBlack,
		})
		table.SetCell(i, 4, &tview.TableCell{
			Text:            record.Shop,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorLightGray,
			BackgroundColor: tcell.ColorBlack,
		})
		table.SetCell(i, 5, &tview.TableCell{
			Text:            record.Place,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorGray,
			BackgroundColor: tcell.ColorBlack,
		})
	}

	table.SetSelectable(true, false)

	if err := tview.NewApplication().SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
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
