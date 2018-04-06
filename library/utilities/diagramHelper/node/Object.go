package diagram_node

import "github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/element"

type Object struct {
	ID   string
	Name string
	Type string
}

func (self *Object) GetElement() (*diagram_element.Object) {

	return diagram_element.NewNode(self.ID, self.Name, self.Type)
}

func NewObject(ID, Name, Type string) (data *Object) {
	data = &Object{
		ID:   ID,
		Name: Name,
		Type: Type,
	}
	return
}

type ElementData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Element struct {
	Data    ElementData `json:"data"`
	Classes string      `json:"classes"`
}
