package visiberms_param

import (
	"os"
	"io/ioutil"
	"github.com/vegh1010/golang.porfolio/library/visiberwc"
	"encoding/xml"
)

func (self *Param) LoadFile() (error) {
	xmlFile, err := os.Open("visiber.xml")
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(byteValue, &self.RawData)
	if err != nil {
		return err
	}
	self.RawData.TrimSpace()
	self.Formatter = visiberwc.NewFormatter(self.RawData)

	return nil
}
