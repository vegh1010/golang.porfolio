package visiberwc_pdf

import "github.com/jung-kurt/gofpdf"

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

	pdf.SetFillColor(250, 250, 250)
	pdf.SetLineWidth(2)
	pdf.DrawPath("DF")

	pdf.SetFont("Times", "B", 20)
	pdf.Text(self.Start.X+8, self.Start.Y+23, self.Text)
}
