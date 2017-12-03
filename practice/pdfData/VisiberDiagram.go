package main

import (
	"github.com/jung-kurt/gofpdf"
)

type VisiberDiagram struct {
	T        *InvertTriangle
	Birthday *BirthdayTextBox
}

func (self *VisiberDiagram) Draw(pdf *gofpdf.Fpdf, data map[string]interface{}) (err error) {
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
