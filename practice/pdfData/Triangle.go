package main

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/paulmach/go.geo"
	"errors"
)

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

func (self *Triangle) Draw(pdf *gofpdf.Fpdf) (err error) {
	pdf.MoveTo(self.Start.X, self.Start.Y)
	pdf.LineTo(self.Start.X+self.Width, self.Start.Y)
	pdf.LineTo(self.Width/2+self.Start.X, self.Start.Y+self.Height)
	pdf.LineTo(self.Start.X, self.Start.Y)

	partition := self.Height/3

	//vertical line
	pdf.MoveTo(self.Width/2+self.Start.X, self.Start.Y)
	pdf.LineTo(self.Width/2+self.Start.X, self.Start.Y+partition*2)

	err = self.DrawIntersection(pdf, partition)
	if err != nil {
		return
	}
	err = self.DrawIntersection(pdf, partition*2)
	if err != nil {
		return
	}

	pdf.MoveTo(self.Start.X, self.Start.Y+partition*3)
	pdf.LineTo(self.Start.X+self.Width, self.Start.Y+partition*3)

	pdf.MoveTo(self.Start.X, self.Start.Y+partition*4)
	pdf.LineTo(self.Start.X+self.Width, self.Start.Y+partition*4)

	pdf.SetFillColor(250, 250, 250)
	pdf.SetLineWidth(2)
	pdf.DrawPath("DF")

	return
}

func (self *Triangle) DrawIntersection(pdf *gofpdf.Fpdf, partition float64) (err error) {
	path := geo.NewPath()
	path.Push(geo.NewPoint(self.Start.X, self.Start.Y))
	path.Push(geo.NewPoint(self.Start.X+self.Width, self.Start.Y))
	path.Push(geo.NewPoint(self.Width/2+self.Start.X, self.Start.Y+self.Height))
	path.Push(geo.NewPoint(self.Start.X, self.Start.Y))

	line := geo.NewLine(geo.NewPoint(self.Start.X, self.Start.Y+partition), geo.NewPoint(self.Start.X+self.Width, self.Start.Y+partition))
	points, _ := path.Intersection(line)
	if len(points) != 2 {
		err = errors.New("no 2 point intersected")
	}

	pdf.MoveTo(points[0].X(), points[0].Y())
	pdf.LineTo(points[1].X(), points[1].Y())

	return
}
