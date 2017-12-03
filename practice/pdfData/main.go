package main

import (
	"os"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	filePath := CreateFilePath("../pdfData", "test.pdf")
	var file *os.File
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pdf := gofpdf.New("P", "pt", "A4", "")
	pdf.AddPage()

	diagram := VisiberDiagram{
		T:        NewInvertTriangle(Point{120, 180}, 250, 350, 30),
		Birthday: NewBirthdayTextBox(Point{145, 140}, 30),
	}
	data := map[string]interface{}{
		"A": 1,
		"B": 8,
		"C": 0,
		"D": 2,
		"E": 1,
		"F": 9,
		"G": 9,
		"H": 1,
		"I": 1,
		"J": 1,
		"K": 1,
		"L": 1,
		"M": 1,
		"N": 1,
		"O": 1,
		"P": 1,
	}
	err = diagram.Draw(pdf, data)
	if err != nil {
		panic(err)
	}

	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		panic(err)
	}
}
