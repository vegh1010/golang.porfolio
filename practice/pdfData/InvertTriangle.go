package main

import (
	"github.com/jung-kurt/gofpdf"
)

func NewInvertTriangle(Start Point, Height, Width, Length float64) (*InvertTriangle) {
	iTriangle := InvertTriangle{}
	iTriangle.T = NewTriangle(Start, Height, Width)
	iTriangle.FieldNames = []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	iTriangle.Fields = map[string]TextBox{}

	spaces := []float64{0, 0, 1, 0, 1, 0, 0, 0}
	var count float64 = 0
	for i, value := range iTriangle.FieldNames {
		count += spaces[i]
		iTriangle.Fields[value] = TextBox{
			Start: Point{
				X: Start.X + count*Length,
				Y: Start.Y,
			},
			Length: Length,
			Text:   value,
		}
		count += 1
	}

	return &iTriangle
}

type InvertTriangle struct {
	T          *Triangle
	FieldNames []string
	Fields     map[string]TextBox
}

func (self *InvertTriangle) Draw(pdf *gofpdf.Fpdf) (err error) {
	err = self.T.Draw(pdf)
	if err != nil {
		return
	}
	//for _, value := range self.FieldNames {
	//	field := self.Fields[value]
	//	field.Draw(pdf)
	//}
	return
}
