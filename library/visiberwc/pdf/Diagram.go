package visiberwc_pdf

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/vegh1010/golang.porfolio/library/visiberwc"
	"strings"
	"sort"
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
	pdf.Ln(5)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 18, element.LBLControl + ": " + element.Control, "", "", false)
	pdf.MultiCell(0, 18, element.LBLProductive + ": " + element.Productive, "", "", false)

	pdf.Ln(15)
	pdf.SetFont("Times", "", 20)
	pdf.MultiCell(0, 18, "Possible Illness", "B", "", false)
	pdf.Ln(5)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 18, "Organs: " + strings.Join(element.Organs, ", "), "", "", false)
	pdf.MultiCell(0, 18, "Ailments: " + strings.Join(element.Ailments, ", "), "", "", false)
	pdf.MultiCell(0, 18, "Symptoms: " + strings.Join(element.Symptoms, ", "), "", "", false)

	pdf.Ln(15)
	pdf.SetFont("Times", "", 20)
	pdf.MultiCell(0, 18, "Traits", "B", "", false)
	pdf.Ln(5)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(0, 18, "Positive: " + character.Positive, "", "", false)
	pdf.MultiCell(0, 18, "Negative: " + character.Negative, "", "", false)

	pdf.Ln(15)
	pdf.SetFont("Times", "", 20)
	pdf.MultiCell(0, 18, "Behavioral Traits", "B", "", false)
	pdf.Ln(5)
	pdf.SetFont("Times", "", 12)
	var traits []string
	for _, field := range data.BehaviourFields {
		traitField := data.Behaviours[field]
		var found bool
		for _, raw := range traits {
			if traitField == raw {
				found = true
			}
		}
		if !found {
			traits = append(traits, traitField)
		}
	}
	sort.Strings(traits)
	for _, field := range traits {
		pdf.MultiCell(0, 18, field + ": " + data.BehavioursData[field].Description, "", "", false)
	}

	//Inside
	//Outside
	//Whole
}

type Point struct {
	X float64
	Y float64
}
