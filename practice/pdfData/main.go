package main

import (
	"github.com/vegh1010/golang.porfolio/library/utilities"
	"github.com/vegh1010/golang.porfolio/library/visiberwc/pdf"
	"os"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	filePath := utilities.CreateFilePath("output", "test.pdf")
	var file *os.File
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pdf := gofpdf.New("P", "pt", "A4", "")
	pdf.AddPage()

	diagram := visiberwc_pdf.Diagram{
		T:        visiberwc_pdf.NewInvertTriangle(visiberwc_pdf.Point{120, 180}, 250, 350, 30),
		Birthday: visiberwc_pdf.NewBirthdayTextBox(visiberwc_pdf.Point{145, 140}, 30),
	}
	data := map[string]int64{
		//"A": 1,
		//"B": 8,
		//"C": 0,
		//"D": 2,
		//"E": 1,
		//"F": 9,
		//"G": 9,
		//"H": 0,
		//"I": 1,
		//"J": 1,
		//"K": 1,
		//"L": 1,
		//"M": 1,
		//"N": 1,
		//"O": 1,
		//"P": 1,
	}
	err = diagram.Draw(pdf, "Val", data)
	if err != nil {
		panic(err)
	}

	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		panic(err)
	}
}
