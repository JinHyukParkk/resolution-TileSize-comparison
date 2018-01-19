package xlsx

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func CreateXlsx() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	// var cell *xlsx.Cell
	var attr = [5]string{"", "Origin", "90%", "80%", "70%"}
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("result")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	for i := 0; i < 5; i++ {
		row.AddCell()
		row.Cells[i].Value = attr[i]
	}
	err = file.Save("CompareResult.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
