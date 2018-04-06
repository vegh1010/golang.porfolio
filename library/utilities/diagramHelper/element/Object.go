package diagram_element

type Object struct {
	Data    ObjectData `json:"data"`
	Classes string     `json:"classes"`
}

func NewNode(ID, Name, Classes string) (data *Object) {
	data = &Object{
		Data: ObjectData{
			ID:   &ID,
			Name: &Name,
		},
		Classes: Classes,
	}
	return
}

func NewBezierEdge(ID, Source, Target string) (data *Object) {
	data = &Object{
		Data: ObjectData{
			ID:     &ID,
			Source: &Source,
			Target: &Target,
		},
		Classes: "bezier",
	}
	return
}

type ObjectData struct {
	ID     *string `json:"id"`
	Name   *string `json:"name"`
	Source *string `json:"source"`
	Target *string `json:"target"`
}
