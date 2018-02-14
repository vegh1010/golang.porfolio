package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"github.com/vegh1010/golang.porfolio/library/visiberwc"
	"github.com/vegh1010/golang.porfolio/library/visiberwc/pdf"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var RawXML visiberwc.RawXML

func initApp(message json.RawMessage) (interface{}, error) {
	err := xml.Unmarshal([]byte(XMLData), &RawXML)
	if err != nil {
		return "", err
	}
	RawXML.TrimSpace()

	return "Loading Complete", nil
}

type User struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func (self *User) Output() string {
	return fmt.Sprint(strings.Replace(self.Name, " ", "_", -1), "_", self.Date)
}

func GenerateReport(message json.RawMessage) (interface{}, error) {
	// Unmarshal payload
	var data User
	if len(message) > 0 {
		// Unmarshal payload
		if err := json.Unmarshal(message, &data); err != nil {
			return "", err
		}
		formatter := visiberwc.NewFormatter(RawXML)
		vUser1, err := formatter.Calculate(data.Name, data.Date)
		if err != nil {
			return "", err
		}

		myself, err := user.Current()
		if err != nil {
			return "", err
		}

		filePath := CreateFilePath(myself.HomeDir + "/Documents", fmt.Sprint(data.Output(), ".pdf"))
		var file *os.File
		file, err = os.Create(filePath)
		if err != nil {
			return "", err
		}
		defer file.Close()

		pdf := gofpdf.New("P", "pt", "A4", "")
		pdf.AddPage()

		diagram := visiberwc_pdf.Diagram{
			T:        visiberwc_pdf.NewInvertTriangle(visiberwc_pdf.Point{X: 120, Y: 180}, 250, 350, 30),
			Birthday: visiberwc_pdf.NewBirthdayTextBox(visiberwc_pdf.Point{X: 145, Y: 140}, 30),
		}
		err = diagram.Draw(pdf, vUser1.Name, vUser1.Fields)
		if err != nil {
			return "", err
		}

		diagram.Detail(pdf, vUser1)

		err = pdf.OutputFileAndClose(filePath)
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("No Payload Found")
	}

	return fmt.Sprint(data.Output(), ".pdf is generated on your Documents folder"), nil
}

/*
 * Create File Path
 */
func CreateFilePath(pathToDirectory, filename string) (filePath string) {
	fmt.Println("CreateFilePath()", filename)
	// write to file
	if pathToDirectory != "" && filename != "" {
		if _, err := os.Stat(pathToDirectory); os.IsNotExist(err) {
			//create directory if not exists
			os.MkdirAll(pathToDirectory, os.ModePerm)
		}
		filePath = filepath.Join(pathToDirectory, filename)
	}

	return
}
