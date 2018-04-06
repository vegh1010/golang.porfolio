package visiberms_param

import "github.com/vegh1010/golang.porfolio/library/visiberwc"

type Param struct {
	RawData      visiberwc.RawXML
	Formatter    *visiberwc.Formatter
	Path         string
	OutputFolder string
}

func (self *Param) Init() (err error) {
	self.Path = "./"
	self.OutputFolder = "Output"
	err = self.LoadFile()

	return
}
