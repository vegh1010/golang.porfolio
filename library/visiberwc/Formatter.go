package visiberwc

import "fmt"

type Formatter struct {
	RawCharacters    map[string]CharacterNumber
	RawElements      map[string]Element
	RawTraits        map[string]Trait
	RawRelationships map[string]Relationship
	RawGroups        map[string]Group
}

func (self *Formatter) Calculate(date string) (User, error) {
	fmt.Println("Formatter.Calculate() - " + date)

	var vUser User
	err := vUser.Parse(date)
	if err != nil {
		return vUser, err
	}

	vUser.CharacterData = self.RawCharacters[fmt.Sprint(vUser.Character)]
	vUser.ElementData = self.RawElements[fmt.Sprint(vUser.Character)]
	vUser.GroupData = self.RawGroups[vUser.Group]

	for _, data := range vUser.Behaviours {
		vUser.BehavioursData = append(vUser.BehavioursData, self.RawTraits[data])
	}
	for _, data := range vUser.Insides {
		vUser.InsidesData = append(vUser.InsidesData, self.RawCharacters[data])
	}
	for _, data := range vUser.Outsides {
		vUser.OutsidesData = append(vUser.OutsidesData, self.RawCharacters[data])
	}

	return vUser, nil
}

func (self *Formatter) Compatibility(vUser1, vUser2 User) (Relationship, error) {
	fmt.Println("Formatter.Compatibility()")

	var data Relationship
	character, err := reduce(fmt.Sprint(vUser1.Character + vUser2.Character))
	if err != nil {
		return data, err
	}
	data = self.RawRelationships[fmt.Sprint(character)]

	return data, nil
}

func NewFormatter(vData RawXML) *Formatter {
	user := Formatter{}

	user.RawCharacters = map[string]CharacterNumber{}
	user.RawElements = map[string]Element{}
	user.RawTraits = map[string]Trait{}
	user.RawRelationships = map[string]Relationship{}
	user.RawGroups = map[string]Group{}

	for _, data := range vData.Characters {
		user.RawCharacters[data.ID] = data
	}

	for _, data := range vData.Elements {
		for _, number := range data.Numbers {
			user.RawElements[number] = data
		}
	}

	for _, data := range vData.Traits {
		user.RawTraits[data.Group] = data
	}

	for _, data := range vData.Relationships {
		user.RawRelationships[data.ID] = data
	}

	for _, data := range vData.Groups {
		user.RawGroups[data.ID] = data
	}

	return &user
}
