package diagramHelper

import (
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/element"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/template"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"
)

func (self *Diagram) Generate() (err error) {
	var elements []*diagram_element.Object
	for _, element := range self.Nodes {
		elements = append(elements, element.GetElement())
	}
	for _, edge := range self.Edges {
		elements = append(elements, edge.GetElement())
	}

	err = self.Create()
	if err != nil {
		return
	}

	var NodeStyling []*diagram_node.Styling
	for _, node := range self.NodeStyling {
		NodeStyling = append(NodeStyling, node)
	}
	template := diagram_template.NewObject(
		self.Title,
		elements,
		self.DefaultNodeStyling,
		NodeStyling,
		self.EdgeStyling,
		self.Layout,
	)

	var output string
	output, err = template.Output()
	if err != nil {
		return
	}

	err = self.Write(output)
	if err != nil {
		return
	}
	return
}
