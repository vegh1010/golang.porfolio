package main

import (
	"github.com/vegh1010/golang.porfolio/library/utilities"
	"github.com/vegh1010/golang.porfolio/library/visiberwc/pdf"
	"os"
	"github.com/jung-kurt/gofpdf"
	"github.com/vegh1010/golang.porfolio/library/visiberwc"
	"io/ioutil"
	"encoding/xml"
	"fmt"
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

	formatter := visiberwc.NewFormatter(data)

	vUser1, err := formatter.Calculate("Barack Obama", "04081961")
	check(err)
	vUser2, err := formatter.Calculate("Donald Trump", "14071946")
	check(err)

	relation, err := formatter.Compatibility(vUser1, vUser2)
	check(err)

	filePath, err := utilities.CreateFilePath("output", fmt.Sprint(vUser1.Name, "_", vUser1.Date, "_", vUser2.Name, "_", vUser2.Date, ".pdf"))
	check(err)
	var file *os.File
	file, err = os.Create(filePath)
	check(err)
	defer file.Close()

	pdf := gofpdf.New("P", "pt", "A4", "")

	err = Generate(pdf, &relation, vUser1, vUser2)
	check(err)

	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		panic(err)
	}
}

func Generate(pdf *gofpdf.Fpdf, relation *visiberwc.Relationship, users ... visiberwc.User) (err error) {
	for index, userData := range users {
		pdf.AddPage()
		diagram1 := visiberwc_pdf.Diagram{
			T:        visiberwc_pdf.NewInvertTriangle(visiberwc_pdf.Point{X: 120, Y: 180}, 250, 350, 30),
			Birthday: visiberwc_pdf.NewBirthdayTextBox(visiberwc_pdf.Point{X: 145, Y: 140}, 30),
		}
		err = diagram1.Draw(pdf, userData.Name, userData.Fields)
		if err != nil {
			return
		}
		diagram1.Detail(pdf, userData)

		if relation != nil && (index == len(users)-1) {
			diagram1.Relation(pdf, *relation)
		}
	}
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
