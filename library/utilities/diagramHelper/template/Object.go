package diagram_template

import (
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/element"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/edge"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/layout"
)

type Object struct {
	Title       string
	Elements    []*diagram_element.Object
	DefaultNode *diagram_node.Styling
	NodeStyling []*diagram_node.Styling
	EdgeStyling *diagram_edge.Styling
	Layout      *diagram_layout.Styling
}

func NewObject(
	Title string,
	Elements []*diagram_element.Object,
	DefaultNode *diagram_node.Styling,
	NodeStyling []*diagram_node.Styling,
	EdgeStyling *diagram_edge.Styling,
	Layout *diagram_layout.Styling,
) (*Object) {

	if Title == "" {
		Title = "Default Cytoscape"
	}
	return &Object{
		Title:       Title,
		Elements:    Elements,
		DefaultNode: DefaultNode,
		NodeStyling: NodeStyling,
		EdgeStyling: EdgeStyling,
		Layout:      Layout,
	}
}
