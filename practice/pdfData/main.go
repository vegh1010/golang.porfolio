package main

import (
	"github.com/vegh1010/golang.porfolio/library/utilities"
	"github.com/vegh1010/golang.porfolio/library/visiberwc/pdf"
	"os"
	"github.com/jung-kurt/gofpdf"
	"github.com/vegh1010/golang.porfolio/library/visiberwc"
	"io/ioutil"
	"encoding/xml"
)

func main() {
	xmlFile, err := os.Open("visiber.xml")
	check(err)

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	check(err)

	var data visiberwc.RawXML
	err = xml.Unmarshal(byteValue, &data)
	check(err)

	data.TrimSpace()
	data.Print()

	formatter := visiberwc.NewFormatter(data)

	vUser1, err := formatter.Calculate("14021989")
	check(err)
	//vUser1.Print()

	filePath := utilities.CreateFilePath("output", "test.pdf")
	var file *os.File
	file, err = os.Create(filePath)
	check(err)
	defer file.Close()

	pdf := gofpdf.New("P", "pt", "A4", "")
	pdf.AddPage()

	diagram := visiberwc_pdf.Diagram{
		T:        visiberwc_pdf.NewInvertTriangle(visiberwc_pdf.Point{120, 180}, 250, 350, 30),
		Birthday: visiberwc_pdf.NewBirthdayTextBox(visiberwc_pdf.Point{145, 140}, 30),
	}
	err = diagram.Draw(pdf, "Val", vUser1.Fields)
	if err != nil {
		panic(err)
	}

	diagram.Detail(pdf, vUser1)

	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		panic(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
