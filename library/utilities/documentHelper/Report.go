package documentHelper

import (
	"io/ioutil"
	"os"
	"fmt"
)

func NewReport(FilePath, Style, Title, Paragraph string) (data Report) {
	data.FilePath = FilePath
	data.Title = Title
	data.Paragraph = Paragraph
	data.Note = "This document is to be referred with source code."
	data.Font = "courier"
	data.Style = `OL { counter-reset: item }
			      LI { display: block }   
			      LI:before { content: counters(item, ".") ". "; counter-increment: item }`
	data.Style += Style
	return
}

type Report struct {
	FilePath  string
	Style     string
	Font      string
	Title     string
	Paragraph string
	Note      string
	Sections  []Section
}

func (self *Report) AddSection(sections ... Section) {
	self.Sections = append(self.Sections, sections...)
}

func (self *Report) Output() (data string) {
	data += `<html>
			 <head>
			 <style>`
	data += self.Style
	data += `</style>
			 </head>`

	data += `<body style="font-family:` + self.Font + `;">`
	if len(self.Title) != 0 {
		data += `<h1>` + self.Title + `</h1>`
	}
	if len(self.Paragraph) != 0 {
		data += `<p>` + self.Paragraph + `</p>`
	}
	data += `<p style="color:red">` + FontStyle(`*Note: `+self.Note, true, false, true) + `</p>`
	if len(self.Sections) != 0 {
		data += `<h2 id="table_content">Table Of Content:</h2>`
		data += `<ol>`
		for i := 0; i < len(self.Sections); i++ {
			data += `<li><a href="#section_` + fmt.Sprint(i+1) + `" style="color:black; text-decoration:none">` + self.Sections[i].Title + `</a></li>`
		}
		data += `</ol>`
	}

	for i := 0; i < len(self.Sections); i++ {
		data += self.Sections[i].Output(i + 1)
		data += `<br>`
	}
	data += `</body>`
	data += `</html>`
	return
}

//write into file
func (self *Report) Write() error {
	err := ioutil.WriteFile(self.FilePath, []byte(""), os.ModeAppend)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(self.FilePath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(self.Output())
	if err != nil {
		return err
	}
	return nil
}

//write into file
func (self *Report) Create() (err error) {
	if _, err = os.Stat(self.FilePath); os.IsNotExist(err) {
		_, err = os.Create(self.FilePath)
		if err != nil {
			return
		}
	}
	return
}
