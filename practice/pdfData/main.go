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
	"strings"
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

	vUser1, err := GenerateUser(formatter, "Barack Obama", "04081961")
	check(err)
	vUser2, err := GenerateUser(formatter, "Donald Trump", "14071946")
	check(err)

	relation, err := formatter.Compatibility(vUser1, vUser2)
	check(err)

	err = Generate(&relation, vUser1, vUser2)
	check(err)
}

func GenerateUser(formatter *visiberwc.Formatter, name, date string) (visiberwc.User, error) {
	vUser, err := formatter.Calculate(name, date)
	if err != nil {
		return vUser, err
	}
	err = Generate(nil , vUser)
	if err != nil {
		return vUser, err
	}
	return vUser, err
}

func Generate(relation *visiberwc.Relationship, users ... visiberwc.User) (err error) {
	var names []string
	for _, userData := range users {
		names = append(names, userData.Name + "_" + userData.Date)
	}
	var fileName = strings.Replace(strings.Join(names, "_"), " ", "_", -1)
	var filePath string
	filePath, err = utilities.CreateFilePath("output", fmt.Sprint(fileName, ".pdf"))
	if err != nil {
		return
	}

	pdf := gofpdf.New("P", "pt", "A4", "")

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
	err = pdf.OutputFileAndClose(filePath)
	if err != nil {
		return
	}
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
