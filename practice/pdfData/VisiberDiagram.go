package main

import (
	"github.com/jung-kurt/gofpdf"
	"fmt"
	"errors"
)

func NewBirthdayTextBox(Start Point, Length float64) (*BirthdayTextBox) {
	box := BirthdayTextBox{}
	box.Start = Start
	box.Length = Length
	box.FieldNames = []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	box.Fields = map[string]TextBox{}

	spaces := []float64{0, 0, 1, 0, 1, 0, 0, 0}
	var count float64 = 0
	for i, value := range box.FieldNames {
		count += spaces[i]
		box.Fields[value] = TextBox{
			Start: Point{
				X: Start.X + count*Length,
				Y: Start.Y,
			},
			Length: Length,
			Text:   value,
		}
		count += 1
	}
	//box.Fields["A"] = TextBox{Point{Start.X, Start.Y}, "A"}
	//box.Fields["B"] = TextBox{Point{Start.X + float64(len(box.Fields)) * Length, Start.Y}, "B"}

	return &box
}

type BirthdayTextBox struct {
	Start  Point
	Length float64

	FieldNames []string
	Fields     map[string]TextBox
}

func (self *BirthdayTextBox) Draw(pdf *gofpdf.Fpdf) {
	for _, value := range self.FieldNames {
		field := self.Fields[value]
		field.Draw(pdf)
	}
}

func (self *BirthdayTextBox) Insert(Field string, value interface{}) (err error) {
	if data, exist := self.Fields[Field]; !exist {
		err = errors.New("Field " + Field + " Not Found")
		return
	} else {
		data.Text = fmt.Sprint(value)
		self.Fields[Field] = data
	}
	return
}

func NewTriangle(Start Point, Height, Width float64) (*Triangle) {
	triangle := Triangle{}
	triangle.Start = Start
	triangle.Height = Height
	triangle.Width = Width

	return &triangle
}

type Triangle struct {
	Start  Point
	Height float64
	Width  float64
}

func (self *Triangle) Draw(pdf *gofpdf.Fpdf) {
	pdf.MoveTo(self.Start.X, self.Start.Y)
	pdf.LineTo(self.Start.X+self.Width, self.Start.Y)
	pdf.LineTo(self.Width/2+self.Start.X, self.Start.Y+self.Height)
	pdf.LineTo(self.Start.X, self.Start.Y)

	pdf.SetFillColor(200, 200, 200)
	pdf.SetLineWidth(2)
	pdf.DrawPath("DF")
}

type VisiberDiagram struct {
	T        *Triangle
	Birthday *BirthdayTextBox
}

func (self *VisiberDiagram) Draw(pdf *gofpdf.Fpdf) {
	self.T.Draw(pdf)
	self.Birthday.Draw(pdf)

}

type Point struct {
	X float64
	Y float64
}

type TextBox struct {
	Start  Point
	Length float64
	Text   string
}

func (self *TextBox) Draw(pdf *gofpdf.Fpdf) {
	pdf.MoveTo(self.Start.X, self.Start.Y)
	pdf.LineTo(self.Start.X+self.Length, self.Start.Y)
	pdf.LineTo(self.Start.X+self.Length, self.Start.Y+self.Length)
	pdf.LineTo(self.Start.X, self.Start.Y+self.Length)
	pdf.LineTo(self.Start.X, self.Start.Y)

	pdf.SetFillColor(200, 200, 200)
	pdf.SetLineWidth(2)
	pdf.DrawPath("DF")

	pdf.SetFont("Helvetica", "B", 20)
	pdf.Text(self.Start.X + 8, self.Start.Y + 23, self.Text)

}
