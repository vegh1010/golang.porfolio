package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"strconv"
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

	vUser := NewVisiberUser()
	err = vUser.Calculate("14021989")
	check(err)
	vUser.Print()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type VisiberUser struct {
	CharacterFields []string
	Fields          map[string]int64

	BehaviourFields []string
	Behaviours      map[string]string

	InsideFields []string
	Insides      []string

	OutsideFields []string
	Outsides      []string
}

func (self *VisiberUser) Calculate(date string) (error) {
	fmt.Println("VisiberUser.Calculate() - " + date)
	var layout = "02012006"
	_, err := time.Parse(layout, date)
	if err != nil {
		return err
	}

	self.Fields["A"], _ = reduce(date[0:1])
	self.Fields["B"], _ = reduce(date[1:2])
	self.Fields["C"], _ = reduce(date[2:3])
	self.Fields["D"], _ = reduce(date[3:4])
	self.Fields["E"], _ = reduce(date[4:5])
	self.Fields["F"], _ = reduce(date[5:6])
	self.Fields["G"], _ = reduce(date[6:7])
	self.Fields["H"], _ = reduce(date[7:8])

	self.Fields["I"], _ = reduce(fmt.Sprint(self.Fields["A"] + self.Fields["B"]))
	self.Fields["J"], _ = reduce(fmt.Sprint(self.Fields["C"] + self.Fields["D"]))
	self.Fields["K"], _ = reduce(fmt.Sprint(self.Fields["E"] + self.Fields["F"]))
	self.Fields["L"], _ = reduce(fmt.Sprint(self.Fields["G"] + self.Fields["H"]))
	self.Fields["M"], _ = reduce(fmt.Sprint(self.Fields["I"] + self.Fields["J"]))
	self.Fields["N"], _ = reduce(fmt.Sprint(self.Fields["K"] + self.Fields["L"]))

	self.Fields["O"], _ = reduce(fmt.Sprint(self.Fields["M"] + self.Fields["N"]))
	self.Fields["P"], _ = reduce(fmt.Sprint(self.Fields["N"] + self.Fields["O"]))
	self.Fields["Q"], _ = reduce(fmt.Sprint(self.Fields["M"] + self.Fields["O"]))
	self.Fields["R"], _ = reduce(fmt.Sprint(self.Fields["P"] + self.Fields["Q"]))
	self.Fields["V"], _ = reduce(fmt.Sprint(self.Fields["J"] + self.Fields["M"]))
	self.Fields["U"], _ = reduce(fmt.Sprint(self.Fields["I"] + self.Fields["M"]))
	self.Fields["S"], _ = reduce(fmt.Sprint(self.Fields["U"] + self.Fields["V"]))
	self.Fields["W"], _ = reduce(fmt.Sprint(self.Fields["K"] + self.Fields["N"]))
	self.Fields["X"], _ = reduce(fmt.Sprint(self.Fields["L"] + self.Fields["N"]))
	self.Fields["T"], _ = reduce(fmt.Sprint(self.Fields["W"] + self.Fields["X"]))

	for _, behaviourField := range self.BehaviourFields {
		fields := strings.Split(behaviourField, "-")
		self.Behaviours[behaviourField] = fmt.Sprint(self.Fields[fields[0]], "-", self.Fields[fields[1]])
	}

	all := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var insides, outsides []int64
	for _, behaviourField := range self.InsideFields {
		insides = append(insides, self.Fields[behaviourField])
	}
	for _, behaviourField := range self.OutsideFields {
		outsides = append(outsides, self.Fields[behaviourField])
	}
	self.Insides = difference(all, insides)
	self.Outsides = difference(all, outsides)

	return nil
}

func (self *VisiberUser) Print() {
	var characters []string
	for _, field := range self.CharacterFields {
		characters = append(characters, fmt.Sprint(field+": ", self.Fields[field]))
	}
	fmt.Println("Character Fields:", strings.Join(characters, ", "))

	var behaviours []string
	for _, field := range self.BehaviourFields {
		behaviours = append(behaviours, fmt.Sprint(field+": ", self.Behaviours[field]))
	}
	fmt.Println("Behaviour Fields:", strings.Join(behaviours, ", "))
	fmt.Println("Inside Lacking:", strings.Join(self.Insides, ", "))
	fmt.Println("Outside Lacking:", strings.Join(self.Outsides, ", "))
}

func NewVisiberUser() *VisiberUser {
	user := VisiberUser{}

	user.CharacterFields = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X"}
	user.Fields = map[string]int64{}

	user.BehaviourFields = []string{"I-J", "J-K", "K-L", "I-M", "J-M", "K-N", "L-N", "M-N", "M-O", "N-O", "P-Q", "U-V", "W-X"}
	user.Behaviours = map[string]string{}

	user.InsideFields = []string{"I", "J", "K", "L", "M", "N", "O"}
	user.OutsideFields = []string{"S", "U", "V", "W", "X", "T", "P", "Q", "R"}

	for _, field := range user.CharacterFields {
		user.Fields[field] = 0
	}

	return &user
}

type Visiber struct {
	XMLName       xml.Name          `xml:"visiber"`
	Characters    []CharacterNumber `xml:"cnumber"`
	Elements      []Element         `xml:"element"`
	Traits        []Trait           `xml:"trait"`
	Relationships []Relationship    `xml:"relationship"`
	Groups        []Group           `xml:"group"`
}

func (self *Visiber) TrimSpace() {
	for i := 0; i < len(self.Characters); i++ {
		self.Characters[i].TrimSpace()
	}
	for i := 0; i < len(self.Elements); i++ {
		self.Elements[i].TrimSpace()
	}
	for i := 0; i < len(self.Traits); i++ {
		self.Traits[i].TrimSpace()
	}
	for i := 0; i < len(self.Relationships); i++ {
		self.Relationships[i].TrimSpace()
	}
	for i := 0; i < len(self.Groups); i++ {
		self.Groups[i].TrimSpace()
	}
}

func (self *Visiber) Print() {
	for i := 0; i < len(self.Characters); i++ {
		self.Characters[i].Print()
	}
	for i := 0; i < len(self.Elements); i++ {
		self.Elements[i].Print()
	}
	for i := 0; i < len(self.Traits); i++ {
		self.Traits[i].Print()
	}
	for i := 0; i < len(self.Relationships); i++ {
		self.Relationships[i].Print()
	}
	for i := 0; i < len(self.Groups); i++ {
		self.Groups[i].Print()
	}
}

type CharacterNumber struct {
	XMLName      xml.Name `xml:"cnumber"`
	ID           string   `xml:"id,attr"`
	Character    string   `xml:"character"`
	Descriptions []string `xml:"description"`
	Positive     string   `xml:"positive"`
	Negative     string   `xml:"negative"`
}

func (self *CharacterNumber) TrimSpace() {
	self.ID = strings.TrimSpace(self.ID)
	self.Character = strings.TrimSpace(self.Character)
	for i := 0; i < len(self.Descriptions); i++ {
		self.Descriptions[i] = strings.TrimSpace(self.Descriptions[i])
	}
	self.Positive = strings.TrimSpace(self.Positive)
	self.Negative = strings.TrimSpace(self.Negative)
}

func (self *CharacterNumber) Print() {
	fmt.Println("Character ID: " + self.ID)
	fmt.Println("Character Type: " + self.Character)
	fmt.Println("Description: " + strings.Join(self.Descriptions, " "))
	fmt.Println("Positives: " + self.Positive)
	fmt.Println("Negatives: " + self.Negative)
	fmt.Println("-----------------------------------------------------------")
}

type Element struct {
	XMLName       xml.Name `xml:"element"`
	Type          string   `xml:"id,attr"`
	Numbers       []string `xml:"number"`
	Organs        []string `xml:"organ"`
	Ailments      []string `xml:"ailment"`
	Symptoms      []string `xml:"symptom"`
	LBLProductive string   `xml:"lblproductive"`
	Productive    string   `xml:"productive"`
	LBLControl    string   `xml:"lblcontrol"`
	Control       string   `xml:"control"`
}

func (self *Element) TrimSpace() {
	self.Type = strings.TrimSpace(self.Type)
	for i := 0; i < len(self.Numbers); i++ {
		self.Numbers[i] = strings.TrimSpace(self.Numbers[i])
	}
	for i := 0; i < len(self.Organs); i++ {
		self.Organs[i] = strings.TrimSpace(self.Organs[i])
	}
	for i := 0; i < len(self.Ailments); i++ {
		self.Ailments[i] = strings.TrimSpace(self.Ailments[i])
	}
	for i := 0; i < len(self.Symptoms); i++ {
		self.Symptoms[i] = strings.TrimSpace(self.Symptoms[i])
	}
	self.LBLProductive = strings.TrimSpace(self.LBLProductive)
	self.Productive = strings.TrimSpace(self.Productive)
	self.LBLControl = strings.TrimSpace(self.LBLControl)
	self.Control = strings.TrimSpace(self.Control)
}

func (self *Element) Print() {
	fmt.Println("Element Type: " + self.Type)
	fmt.Println("Number: " + strings.Join(self.Numbers, ", "))
	fmt.Println("Organs: " + strings.Join(self.Organs, ", "))
	fmt.Println("Ailments: " + strings.Join(self.Ailments, ", "))
	fmt.Println("Symptoms: " + strings.Join(self.Symptoms, ", "))
	fmt.Println(self.LBLProductive + ": " + self.Productive)
	fmt.Println(self.LBLControl + ": " + self.Control)
	fmt.Println("-----------------------------------------------------------")
}

type Trait struct {
	XMLName     xml.Name `xml:"trait"`
	Group       string   `xml:"id,attr"`
	Description string   `xml:"tdescription"`
}

func (self *Trait) TrimSpace() {
	self.Group = strings.TrimSpace(self.Group)
	self.Description = strings.TrimSpace(self.Description)
}

func (self *Trait) Print() {
	fmt.Println("Trait Group: " + self.Group)
	fmt.Println("Description: " + self.Description)
	fmt.Println("-----------------------------------------------------------")
}

type Relationship struct {
	XMLName     xml.Name `xml:"relationship"`
	ID          string   `xml:"id,attr"`
	Description string   `xml:"rdetail"`
}

func (self *Relationship) TrimSpace() {
	self.ID = strings.TrimSpace(self.ID)
	self.Description = strings.TrimSpace(self.Description)
}

func (self *Relationship) Print() {
	fmt.Println("Relationship ID: " + self.ID)
	fmt.Println("Description: " + self.Description)
	fmt.Println("-----------------------------------------------------------")
}

type Group struct {
	XMLName     xml.Name `xml:"group"`
	ID          string   `xml:"id,attr"`
	Types       []string `xml:"type"`
	Points      []string `xml:"gpoint"`
	Description string   `xml:"gdescription"`
}

func (self *Group) TrimSpace() {
	self.ID = strings.TrimSpace(self.ID)
	for i := 0; i < len(self.Types); i++ {
		self.Types[i] = strings.TrimSpace(self.Types[i])
	}
	for i := 0; i < len(self.Points); i++ {
		self.Points[i] = strings.TrimSpace(self.Points[i])
	}
	self.Description = strings.TrimSpace(self.Description)
}

func (self *Group) Print() {
	fmt.Println("Group ID: " + self.ID)
	fmt.Println("Type: " + strings.Join(self.Types, " - "))
	fmt.Println("Point: " + strings.Join(self.Points, ". "))
	fmt.Println("Description: " + self.Description)
	fmt.Println("-----------------------------------------------------------")
}

func difference(slice1 []int64, slice2 []int64) ([]string) {
	var diffStr []string
	m := map[string]int{}

	for _, s1Val := range slice1 {
		m[fmt.Sprint(s1Val)] = 1
	}
	for _, s2Val := range slice2 {
		m[fmt.Sprint(s2Val)] = m[fmt.Sprint(s2Val)] + 1
	}

	for mKey, mVal := range m {
		if mVal == 1 {
			diffStr = append(diffStr, mKey)
		}
	}

	return diffStr
}

func reduce(value string) (int64, error) {
	ival, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return ival, err
	}
	if ival > 9 {
		val1 := fmt.Sprint(ival)[0:1]
		val2 := fmt.Sprint(ival)[1:]
		ival1, err := strconv.ParseInt(val1, 10, 0)
		if err != nil {
			return ival, err
		}
		ival2, err := strconv.ParseInt(val2, 10, 0)
		if err != nil {
			return ival, err
		}
		return reduce(fmt.Sprint(ival1 + ival2))
	}
	return ival, nil
}
