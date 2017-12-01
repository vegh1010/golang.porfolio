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
	diagram.Birthday.Insert("A", 1)
	err = diagram.Draw(pdf)
	if err != nil {
		panic(err)
	}

	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		panic(err)
	}
}
