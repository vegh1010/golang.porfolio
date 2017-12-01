package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

func main() {
	xmlFile, err := os.Open("visiber.xml")
	check(err)

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	check(err)

	var data Visiber
	err = xml.Unmarshal(byteValue, &data)
	check(err)

	data.TrimSpace()
	data.Print()

	formatter := NewVisiberFormatter(data)

	vUser1, err := formatter.Calculate("04051")
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
