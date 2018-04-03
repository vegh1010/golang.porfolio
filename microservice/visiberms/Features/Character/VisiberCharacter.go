package visiberms_character

import (
	"github.com/vegh1010/golang.porfolio/library/visiberwc"
	"github.com/vegh1010/golang.porfolio/microservice/visiberms/Features/Param"
)

type VisiberCharacter struct {
	visiberms_param.Param
	Characters []visiberwc.CharacterNumber
}

func (self *VisiberCharacter) Init() {
	self.Characters = self.RawData.Characters

	return
}

type Result struct {
	ID           string   `json:"id"`
	Character    string   `json:"character"`
	Descriptions []string `json:"description"`
	Positive     string   `json:"positive"`
	Negative     string   `json:"negative"`
}
