package visiberwc_pdf

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/vegh1010/golang.porfolio/library/visiberwc"
	"fmt"
	"strings"
)

type Diagram struct {
	T        *InvertTriangle
	Birthday *BirthdayTextBox
}

func (self *Diagram) Draw(pdf *gofpdf.Fpdf, name string, data map[string]int64) (err error) {
	if name == "" {
		name = "N/A"
	}

	pdf.Ln(20)
	pdf.SetFont("Times", "B", 20)
	pdf.MultiCell(0, 18, "Visiber Number", "", "C", false)
	pdf.Ln(self.Birthday.Start.Y - self.Birthday.Length * 3.5)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 18, "Name: " + name, "B", "", false)

	err = self.T.Draw(pdf, data)
	if err != nil {
		return
	}
	self.Birthday.Draw(pdf, data)

	return
}

func (self *Diagram) Detail(pdf *gofpdf.Fpdf, data visiberwc.User) {
	pdf.AddPage()
	pdf.Ln(self.Birthday.Start.Y - self.Birthday.Length * 3.5)
	pdf.SetFont("Times", "B", 20)
	pdf.MultiCell(0, 18, "Character Number Profile", "B", "", false)

	character := data.CharacterData
	element := data.ElementData

	pdf.Ln(25)
	pdf.SetFont("Times", "", 40)
	pdf.MultiCell(0, 18, character.ID + " " + character.Character, "", "", false)

	pdf.Ln(15)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 18, strings.Join(character.Descriptions, "\n\n"), "", "", false)

	pdf.Ln(15)
	pdf.SetFont("Times", "", 20)
	pdf.MultiCell(0, 18, element.Type, "B", "", false)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 18, element.LBLControl + ": " + element.Control, "", "", false)
	pdf.MultiCell(0, 18, element.LBLProductive + ": " + element.Productive, "", "", false)

	pdf.Ln(15)
	pdf.SetFont("Times", "", 20)
	pdf.MultiCell(0, 18, "Traits", "B", "", false)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 18, "Positive: " + character.Positive, "", "", false)
	pdf.MultiCell(0, 18, "Negative: " + character.Negative, "", "", false)

	//Behavioral Traits
	//Inside
	//Outside
	//Whole
}

type Point struct {
	X float64
	Y float64
}
