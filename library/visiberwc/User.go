package visiberwc

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	CharacterFields []string
	BehaviourFields []string
	InsideFields    []string
	OutsideFields   []string

	Fields        map[string]int64
	Character     int64
	CharacterData CharacterNumber

	Group     string
	GroupData Group

	Behaviours     map[string]string
	BehavioursData []Trait

	Insides     []string
	InsidesData []CharacterNumber

	Outsides     []string
	OutsidesData []CharacterNumber
}

func (self *User) Init() {
	self.CharacterFields = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X"}
	self.BehaviourFields = []string{"I-J", "J-K", "K-L", "I-M", "J-M", "K-N", "L-N", "M-N", "M-O", "N-O", "P-Q", "U-V", "W-X"}
	self.InsideFields = []string{"I", "J", "K", "L", "M", "N", "O"}
	self.OutsideFields = []string{"S", "U", "V", "W", "X", "T", "P", "Q", "R"}

	self.Fields = map[string]int64{}
	for _, field := range self.CharacterFields {
		self.Fields[field] = 0
	}
	self.Character = 0
	self.CharacterData = CharacterNumber{}

	self.Group = ""
	self.GroupData = Group{}

	self.Behaviours = map[string]string{}
	self.BehavioursData = []Trait{}

	self.Insides = []string{}
	self.InsidesData = []CharacterNumber{}

	self.Outsides = []string{}
	self.OutsidesData = []CharacterNumber{}
}

func (self *User) Parse(date string) (err error) {
	self.Init()

	var layout = "02012006"
	_, err = time.Parse(layout, date)
	if err != nil {
		return
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

	self.Character = self.Fields["O"]
	self.Group = fmt.Sprint(self.Fields["O"], "-", self.Fields["M"], "-", self.Fields["N"])

	return
}

func (self *User) Print() {
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
	fmt.Println("Character:", self.Character)
	fmt.Println("Group:", self.Group)

	self.CharacterData.Print()
	//self.GroupData.Print()
	//for _, data := range self.BehavioursData {
	//	data.Print()
	//}
	//for _, data := range self.InsidesData {
	//	data.Print()
	//}
	//for _, data := range self.OutsidesData {
	//	data.Print()
	//}
}
