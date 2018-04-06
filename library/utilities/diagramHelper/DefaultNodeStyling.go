package diagramHelper

import (
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"
)

//https://cdn.rawgit.com/cytoscape/cytoscape.js/master/dist/cytoscape.min.js
func NewDefaultNodeStyling() (instance *DefaultNodeStyling, err error) {
	instance = new(DefaultNodeStyling)
	instance.MICROSERVICE, err = NewNodeStyling("microservice", diagram_node.Octagon, "#3473d8")
	if err != nil {
		return
	}
	instance.WORKER, err = NewNodeStyling("worker", diagram_node.Ellipse, "green")
	if err != nil {
		return
	}
	instance.DATABASE, err = NewNodeStylingHeight("database", diagram_node.RoundRectangle, "#F27E31")
	if err != nil {
		return
	}
	instance.COLLECTOR, err = NewNodeStyling("collector", diagram_node.Tag, "#f751e9")
	if err != nil {
		return
	}
	instance.QUEUE, err = NewNodeStylingHeight("queue", diagram_node.RoundRectangle, "#f744f4")
	if err != nil {
		return
	}
	instance.OTHER, err = NewNodeStylingHeight("other", diagram_node.RoundRectangle, "red")
	if err != nil {
		return
	}
	instance.STANDALONE, err = NewNodeStyling("stand_alone", diagram_node.Polygon, "#ef2f2f")
	if err != nil {
		return
	}

	return
}

type DefaultNodeStyling struct {
	WORKER       *diagram_node.Styling
	DATABASE     *diagram_node.Styling
	MICROSERVICE *diagram_node.Styling
	COLLECTOR    *diagram_node.Styling
	QUEUE        *diagram_node.Styling
	OTHER        *diagram_node.Styling
	STANDALONE   *diagram_node.Styling
}

func (self *DefaultNodeStyling) List() (list []*diagram_node.Styling) {
	list = append(list, self.WORKER)
	list = append(list, self.DATABASE)
	list = append(list, self.MICROSERVICE)
	list = append(list, self.COLLECTOR)
	list = append(list, self.QUEUE)
	list = append(list, self.OTHER)
	list = append(list, self.STANDALONE)

	return
}

func NewNodeStyling(name, shape, background string) (*diagram_node.Styling, error) {
	instance, err := diagram_node.NewStyling(
		name,
		diagram_node.NewShape(shape),
		diagram_node.NewBackgroundColor(background),
	)
	if err != nil {
		return instance, err
	}
	return instance, nil
}

func NewNodeStylingHeight(name, shape, background string) (*diagram_node.Styling, error) {
	instance, err := diagram_node.NewStyling(
		name,
		diagram_node.NewShape(shape),
		diagram_node.NewBackgroundColor(background),
		diagram_node.NewHeight("1px"),
	)
	if err != nil {
		return instance, err
	}
	return instance, nil
}
