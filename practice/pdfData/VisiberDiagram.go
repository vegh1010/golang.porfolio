package main

import (
	"github.com/jung-kurt/gofpdf"
)

type VisiberDiagram struct {
	T        *InvertTriangle
	Birthday *BirthdayTextBox
}

func (self *VisiberDiagram) Draw(pdf *gofpdf.Fpdf) (err error) {
	err = self.T.Draw(pdf)
	if err != nil {
		return
	}
	self.Birthday.Draw(pdf)

	return
}

type Point struct {
	X float64
	Y float64
}
