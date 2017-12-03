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

	return &box
}

type BirthdayTextBox struct {
	Start  Point
	Length float64

	FieldNames []string
	Fields     map[string]TextBox
}

func (self *BirthdayTextBox) Draw(pdf *gofpdf.Fpdf, data map[string]interface{}) {
	for _, value := range self.FieldNames {
		field := self.Fields[value]
		if number, exist := data[value]; exist  {
			field.Text = fmt.Sprint(number)
		}
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
