package visiberwc_pdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func NewInvertTriangle(Start Point, Height, Width, Length float64) (*InvertTriangle) {
	iTriangle := InvertTriangle{}
	iTriangle.T = NewTriangle(Start, Height, Width)
	iTriangle.FieldNames = []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
	}
	iTriangle.Fields = map[string]TextBox{}

	partitionH := Height / 6
	partitionW := Width / 4

	//segment 1
	iTriangle.segment1(Start, partitionH, partitionW, Length)

	//segment 2
	iTriangle.segment2(Start, partitionH*3, partitionW, Length)

	//segment 3
	iTriangle.segment3(Start, partitionH*4.75, partitionW, Length)

	//segment 4
	iTriangle.segment4(Start, partitionH*7, partitionW, Length)

	//segment 5
	iTriangle.segment5(Start, partitionH*9, partitionW, Length)

	return &iTriangle
}

type InvertTriangle struct {
	T          *Triangle
	FieldNames []string
	Fields     map[string]TextBox
}

func (self *InvertTriangle) segment1(Start Point, partitionH, partitionW, Length float64) {
	self.Fields["A"] = TextBox{
		Start: Point{
			X: Start.X + partitionW - Length,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "A",
	}
	self.Fields["B"] = TextBox{
		Start: Point{
			X: Start.X + partitionW + Length,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "B",
	}
	self.Fields["C"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*3 - Length*2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "C",
	}
	self.Fields["D"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*3,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "D",
	}
}

func (self *InvertTriangle) segment2(Start Point, partitionH, partitionW, Length float64) {
	self.Fields["E"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*2 - Length*2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "E",
	}
	self.Fields["F"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*2 + Length,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "F",
	}
}

func (self *InvertTriangle) segment3(Start Point, partitionH, partitionW, Length float64) {
	self.Fields["G"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*2 - Length/2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "G",
	}

	self.Fields["K"] = TextBox{
		Start: Point{
			X: Start.X + partitionW - Length/2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "K",
	}
	self.Fields["L"] = TextBox{
		Start: Point{
			X: Start.X + partitionW - Length*2 - Length/2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "L",
	}
	self.Fields["M"] = TextBox{
		Start: Point{
			X: Start.X + partitionW - Length*4 - Length/2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "M",
	}

	self.Fields["N"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*3 - Length/2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "N",
	}
	self.Fields["O"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*3 + Length*2 - Length/2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "O",
	}
	self.Fields["P"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*3 + Length*4 - Length/2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "P",
	}
}

func (self *InvertTriangle) segment4(Start Point, partitionH, partitionW, Length float64) {
	self.Fields["H"] = TextBox{
		Start: Point{
			X: Start.X + partitionW - Length,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "H",
	}
	self.Fields["I"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*3,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "I",
	}
}

func (self *InvertTriangle) segment5(Start Point, partitionH, partitionW, Length float64) {
	self.Fields["J"] = TextBox{
		Start: Point{
			X: Start.X + partitionW*2 - Length/2,
			Y: Start.Y + partitionH - Length/2,
		},
		Length: Length,
		Text:   "J",
	}
}

func (self *InvertTriangle) Draw(pdf *gofpdf.Fpdf, data map[string]int64) (err error) {
	err = self.T.Draw(pdf)
	if err != nil {
		return
	}
	for _, value := range self.FieldNames {
		field := self.Fields[value]
		if number, exist := data[value]; exist  {
			field.Text = fmt.Sprint(number)
		}
		field.Draw(pdf)
	}
	return
}
