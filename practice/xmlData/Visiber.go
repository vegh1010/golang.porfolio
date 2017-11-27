package main

import (
	"fmt"
	"encoding/xml"
	"strings"
)

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
