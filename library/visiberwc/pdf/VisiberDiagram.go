package visiberwc_pdf

import (
	"github.com/jung-kurt/gofpdf"
)

type Diagram struct {
	T        *InvertTriangle
	Birthday *BirthdayTextBox
}

func (self *Diagram) Draw(pdf *gofpdf.Fpdf, data map[string]int64) (err error) {
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
