package diagram_edge

import "github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/element"

type Object struct {
	ID     string
	Source string
	Target string
}

func (self *Object) GetElement() (*diagram_element.Object) {

	return diagram_element.NewBezierEdge(self.ID, self.Source, self.Target)
}

func NewObject(ID, Source, Target string) (data *Object) {
	data = &Object{
		ID:     ID,
		Source: Source,
		Target: Target,
	}
	return
}
