package pkg

import (
	"github.com/xuri/excelize/v2"
)

func Exec() {
	file1, err := excelize.OpenFile("./files/file1.xlsx")

	if err != nil {
		panic("Not load first file")
	}

	file2, err := excelize.OpenFile("./files/file2.xlsx")

	if err != nil {
		panic("Not load two file")
	}

	Rows1, err := file1.GetRows("Лист1")

	if err != nil {
		panic("Incorrect file1 rows")
	}

	Rows2, err := file2.GetRows("Лист1")

	if err != nil {
		panic("Incorrect file2 rows")
	}

	Parser(Rows1, Rows2)
}
