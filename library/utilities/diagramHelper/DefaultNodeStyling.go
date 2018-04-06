package diagramHelper

import "github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"

//var (
//WORKER       = "worker"
//DATABASE     = "database"
//MICROSERVICE = "microservice"
//COLLECTOR    = "collector"
//QUEUE        = "queue"
//OTHER        = "other"
//STANDALONE   = "stand_alone"
//)

type DefaultNodeStyling struct {
	WORKER       *diagram_node.Styling
	DATABASE     *diagram_node.Styling
	MICROSERVICE *diagram_node.Styling
	COLLECTOR    *diagram_node.Styling
	QUEUE        *diagram_node.Styling
	OTHER        *diagram_node.Styling
	STANDALONE   *diagram_node.Styling
}

//https://cdn.rawgit.com/cytoscape/cytoscape.js/master/dist/cytoscape.min.js
func NewDefaultNodeStyling() (instance *DefaultNodeStyling, err error) {
	instance.WORKER, err = NewNodeStyling("worker", "", "")
	if err != nil {
		return
	}
	instance.DATABASE, err = NewNodeStyling("database", "", "")
	if err != nil {
		return
	}
	instance.MICROSERVICE, err = NewNodeStyling("microservice", "", "")
	if err != nil {
		return
	}
	instance.COLLECTOR, err = NewNodeStyling("collector", "", "")
	if err != nil {
		return
	}
	instance.QUEUE, err = NewNodeStyling("queue", "", "")
	if err != nil {
		return
	}
	instance.OTHER, err = NewNodeStyling("other", "", "")
	if err != nil {
		return
	}
	instance.STANDALONE, err = NewNodeStyling("stand_alone", "", "")
	if err != nil {
		return
	}

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
