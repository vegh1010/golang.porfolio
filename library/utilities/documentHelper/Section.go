package documentHelper

import "fmt"

func NewSection(Title, Paragraph string, RawBody *string, Data []ReportData, Note string) (data Section) {
	data.Title = Title
	data.Paragraph = Paragraph
	data.RawBody = RawBody
	data.Data = Data
	data.Note = Note
	return
}

type Section struct {
	Title     string
	Paragraph string
	RawBody   *string
	Data      []ReportData
	Note      string
}

func (self *Section) Output(tag int) (data string) {
	if len(self.Title) != 0 {
		data += ` <a href="#table_content" style="color:black; text-decoration:none">`
		data += `<h2`
		if tag != 0 {
			data += ` id="section_` + fmt.Sprint(tag) + `"`
		}
		data += `>`
		if tag != 0 {
			data += fmt.Sprint(tag, ". ")
		}
		data += self.Title + `</h2></a>`
	}
	if len(self.Paragraph) != 0 {
		data += `<p>` + self.Paragraph + `</p>`
	}
	if self.RawBody != nil {
		data += *self.RawBody
	} else {
		steps := self.Data
		if len(steps) != 0 {
			data += `<ol>`
			for i := 0; i < len(steps); i++ {
				data += steps[i].Output()
			}
			data += `</ol>`
		}
		if len(self.Note) != 0 {
			data += `<p>Note: ` + self.Note + `</p>`
		}
	}
	return
}
