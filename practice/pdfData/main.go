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
		T:        NewTriangle(Point{120, 80}, 420, 350),
		Birthday: NewBirthdayTextBox(Point{145, 40}, 30),
	}
	diagram.Birthday.Insert("A", 1)
	diagram.Draw(pdf)

	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		panic(err)
	}
}
