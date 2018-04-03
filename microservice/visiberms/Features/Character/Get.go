package visiberms_character

import "fmt"

func (self *VisiberCharacter) Get(id interface{}) (Result) {
	fmt.Println("VisiberCharacter.Get()")
	var data Result
	for _, character := range self.Characters {
		if fmt.Sprint(id) == character.ID {
			data = Result{
				ID:           character.ID,
				Character:    character.Character,
				Descriptions: character.Descriptions,
				Positive:     character.Positive,
				Negative:     character.Negative,
			}
			return data
		}
	}

	return data
}
