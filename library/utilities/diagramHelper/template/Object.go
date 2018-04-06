package diagram_template

import (
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/element"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/edge"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/layout"
	"encoding/json"
	"strings"
)

type Object struct {
	Title       string
	Elements    []*diagram_element.Object
	DefaultNode *diagram_node.DefaultStyling
	NodeStyling []*diagram_node.Styling
	EdgeStyling *diagram_edge.BezierStyling
	Layout      *diagram_layout.Styling
}

func NewObject(
	Title string,
	Elements []*diagram_element.Object,
	DefaultNode *diagram_node.DefaultStyling,
	NodeStyling []*diagram_node.Styling,
	EdgeStyling *diagram_edge.BezierStyling,
	Layout *diagram_layout.Styling,
) (*Object) {

	if Title == "" {
		Title = "Default Cytoscape"
	}
	if DefaultNode == nil {
		DefaultNode = diagram_node.NewDefaultStyling()
	}
	if EdgeStyling == nil {
		EdgeStyling = diagram_edge.NewDefaultBezierStyling()
	}
	if Layout == nil {
		Layout = diagram_layout.NewDefaultStyling()
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

func (self *Object) GetDefaultNode() (result map[string]interface{}) {
	data := *self.DefaultNode

	result = map[string]interface{}{}
	result["selector"] = data.Selector

	style := map[string]interface{}{}
	for _, option := range data.Style {
		style[option.Tag] = option.Value
	}
	result["style"] = style

	return
}

func (self *Object) GetNodeStyling() (results []map[string]interface{}) {
	for _, raw := range self.NodeStyling {
		data := *raw

		result := map[string]interface{}{}
		result["selector"] = data.Selector

		style := map[string]interface{}{}
		for _, option := range data.Style {
			style[option.Tag] = option.Value
		}
		result["style"] = style

		results = append(results, result)
	}

	return
}

func (self *Object) GetEdgeStyling() (result map[string]interface{}) {
	data := *self.EdgeStyling

	result = map[string]interface{}{}
	result["selector"] = data.Selector

	style := map[string]interface{}{}
	for _, option := range data.Style {
		style[option.Tag] = option.Value
	}
	result["style"] = style

	return
}

func (self *Object) GetLayout() (result map[string]interface{}) {
	data := *self.Layout

	result = map[string]interface{}{}
	result["name"] = data.Name
	result["avoidOverlap"] = data.AvoidOverlap

	return
}

func (self *Object) Output() (string, error) {
	output := Template

	byteElements, err := json.Marshal(self.Elements)
	if err != nil {
		return output, err
	}

	var style []map[string]interface{}
	style = append(style, self.GetDefaultNode())
	style = append(style, self.GetNodeStyling()...)
	style = append(style, self.GetEdgeStyling())

	byteStyle, err := json.Marshal(style)
	if err != nil {
		return output, err
	}

	byteLayout, err := json.Marshal(self.GetLayout())
	if err != nil {
		return output, err
	}

	output = strings.Replace(output, `%title%`, self.Title, -1)
	output = strings.Replace(output, `%elements%`, string(byteElements), -1)
	output = strings.Replace(output, `%style%`, string(byteStyle), -1)
	output = strings.Replace(output, `%layout%`, string(byteLayout), -1)

	return output, nil
}
