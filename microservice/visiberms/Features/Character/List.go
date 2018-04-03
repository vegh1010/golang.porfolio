package visiberms_character

import "fmt"

func (self *VisiberCharacter) List() (results []Result) {
	fmt.Println("VisiberCharacter.List()")
	for _, character := range self.Characters {
		data := Result{
			ID:           character.ID,
			Character:    character.Character,
			Descriptions: character.Descriptions,
			Positive:     character.Positive,
			Negative:     character.Negative,
		}
		results = append(results, data)
	}

	return
}
