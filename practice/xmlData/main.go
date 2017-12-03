package main

import (
	"encoding/xml"
	"github.com/vegh1010/golang.porfolio/library/visiberwc"
	"io/ioutil"
	"os"
)

func main() {
	xmlFile, err := os.Open("visiber.xml")
	check(err)

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	check(err)

	var data visiberwc.RawXML
	err = xml.Unmarshal(byteValue, &data)
	check(err)

	data.TrimSpace()
	data.Print()

	formatter := visiberwc.NewFormatter(data)

	vUser1, err := formatter.Calculate("14021989")
	check(err)
	vUser1.Print()

	//vUser2, err := formatter.Calculate("18021990")
	//check(err)
	//vUser2.Print()
	//
	//relation, err := formatter.Compatibility(vUser1, vUser2)
	//check(err)
	//relation.Print()
}
