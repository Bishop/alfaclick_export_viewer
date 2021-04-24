package main

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	data := new(Database)

	fileName, err := filepath.Abs(getFileName())

	fatalOnError(err)

	readXmlData(fileName, data)

	records := recordsFromItems(data)

	if false {
		printRecords(records, "display.go.tmpl")
	} else {
		createUiTable(fileName, records)
	}
}

func getFileName() (fileName string) {
	if len(os.Args) == 2 {
		fileName = os.Args[1]
	} else {
		files, err := filepath.Glob("./*.xml")

		fatalOnError(err)

		if len(files) > 0 {
			fileName = files[0]
		} else {
			println("No filename was given")
			os.Exit(1)
		}
	}

	return
}

func createUiTable(title string, records *Records) {
	table := tview.NewTable()
	for i, record := range *records {
		table.SetCell(i, 0, &tview.TableCell{
			Text:            record.AccDateS(),
			Align:           tview.AlignLeft,
			Color:           tcell.ColorLightGray,
			BackgroundColor: tcell.ColorBlack,
		})
		table.SetCell(i, 1, &tview.TableCell{
			Text:            record.TxDateS(),
			Align:           tview.AlignLeft,
			Color:           tcell.ColorGray,
			BackgroundColor: tcell.ColorBlack,
		})

		color1 := tcell.ColorLimeGreen
		color2 := tcell.ColorLightGreen

		if record.AccAmount.Amount >= 0 {
			color1 = tcell.ColorAquaMarine
			color2 = tcell.ColorAquaMarine
		}

		table.SetCell(i, 2, &tview.TableCell{
			Text:            record.AccAmount.String(),
			Align:           tview.AlignLeft,
			Color:           color1,
			BackgroundColor: tcell.ColorBlack,
		})
		table.SetCell(i, 3, &tview.TableCell{
			Text:            record.TxAmountS(),
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
	table.SetTitle(title).SetBorder(true).SetBorderColor(tcell.ColorGray)

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
		panic(err)
	}
}
