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
	data, err := readData()
	check(err)

	formatter := visiberwc.NewFormatter(data)

	list := [][]string{
		{"Barack Obama", "04081961", "Donald Trump", "14071946"},
	}

	for _, record := range list {
		if len(record) == 4 {
			err := GeneratePair(formatter, record[0], record[1], record[2], record[3])
			check(err)
		} else if len(record) == 2 {
			_, err := GenerateUser(formatter, record[0], record[1])
			check(err)
		}
	}
}

func GeneratePair(formatter *visiberwc.Formatter, name1, date1, name2, date2 string) (error) {
	vUser1, err := GenerateUser(formatter, name1, date1)
	if err != nil {
		return err
	}
	vUser2, err := GenerateUser(formatter, name2, date2)
	if err != nil {
		return err
	}

	relation, err := formatter.Compatibility(vUser1, vUser2)
	if err != nil {
		return err
	}

	err = Generate(&relation, vUser1, vUser2)
	if err != nil {
		return err
	}

	return nil
}

func readData() (visiberwc.RawXML, error) {
	var data visiberwc.RawXML
	xmlFile, err := os.Open("visiber.xml")
	if err != nil {
		return data, err
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return data, err
	}

	err = xml.Unmarshal(byteValue, &data)
	if err != nil {
		return data, err
	}

	data.TrimSpace()

	return data, nil
}

func GenerateUser(formatter *visiberwc.Formatter, name, date string) (visiberwc.User, error) {
	vUser, err := formatter.Calculate(name, date)
	if err != nil {
		return vUser, err
	}
	err = Generate(nil, vUser)
	if err != nil {
		return vUser, err
	}
	return vUser, err
}

func Generate(relation *visiberwc.Relationship, users ... visiberwc.User) (err error) {
	var names []string
	for _, userData := range users {
		names = append(names, userData.Name+"_"+userData.Date)
	}
	var fileName = strings.Replace(strings.Join(names, "_"), " ", "_", -1)
	var filePath string
	filePath = utilities.CreateFilePath("output", fmt.Sprint(fileName, ".pdf"))

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
