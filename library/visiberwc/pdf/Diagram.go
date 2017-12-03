package visiberwc_pdf

import (
	"github.com/jung-kurt/gofpdf"
)

type Diagram struct {
	T        *InvertTriangle
	Birthday *BirthdayTextBox
}

func (self *Diagram) Draw(pdf *gofpdf.Fpdf, name string, data map[string]int64) (err error) {
	if name == "" {
		name = "N/A"
	}
	pdf.MoveTo(self.Birthday.Start.X - self.Birthday.Length * 2.5, self.Birthday.Start.Y - self.Birthday.Length * 1.2)
	pdf.SetFont("Helvetica", "B", 20)
	pdf.MultiCell(0, 18, name, "", "", false)
	pdf.MultiCell(0, 1, "", "", "", true)

	err = self.T.Draw(pdf, data)
	if err != nil {
		return
	}
	self.Birthday.Draw(pdf, data)

	return
}

type Point struct {
	X float64
	Y float64
}
