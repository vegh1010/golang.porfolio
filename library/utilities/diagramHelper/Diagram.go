package diagramHelper

import (
	"github.com/vegh1010/golang.porfolio/library/utilities"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/edge"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/layout"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"
	"strings"
	"github.com/pkg/errors"
)

//https://cdn.rawgit.com/cytoscape/cytoscape.js/master/dist/cytoscape.min.js
func NewDiagram(
	folder,
	Filename,
	Title string,
	DefaultNodeStyling *diagram_node.DefaultStyling,
	EdgeStyling *diagram_edge.BezierStyling,
	Layout *diagram_layout.Styling,
) (*Diagram) {

	var filePath = "./"
	if folder != "" {
		filePath += folder + "/"
	}

	return &Diagram{
		FolderPath:         filePath,
		FilePath:           filePath + Filename + ".html",
		Filename:           Filename,
		Title:              Title,
		Nodes:              map[string]*diagram_node.Object{},
		Edges:              map[string]*diagram_edge.Object{},
		DefaultNodeStyling: DefaultNodeStyling,
		NodeStyling:        map[string]*diagram_node.Styling{},
		EdgeStyling:        EdgeStyling,
		Layout:             Layout,
	}
}

type Diagram struct {
	FolderPath string
	FilePath   string
	Filename   string
	Title      string

	Nodes              map[string]*diagram_node.Object
	DefaultNodeStyling *diagram_node.DefaultStyling
	NodeStyling        map[string]*diagram_node.Styling

	Edges       map[string]*diagram_edge.Object
	EdgeStyling *diagram_edge.BezierStyling

	Layout *diagram_layout.Styling
}

func (self *Diagram) AddNode(Name string, Type *diagram_node.Styling) (*diagram_node.Object) {
	var uuid string
	for {
		uuid = utilities.GetUuid().String()
		if _, exist := self.Nodes[uuid]; !exist {
			break
		}
	}
	self.Nodes[uuid] = diagram_node.NewObject(uuid, Name, strings.Replace(Type.Selector, "node.", "", -1))

	return self.Nodes[uuid]
}

func (self *Diagram) AddNodeStyling(name string, opinionStyles ... *diagram_node.Option) (*diagram_node.Styling, error) {
	instance, err := diagram_node.NewStyling(
		name,
		opinionStyles ...,
	)
	if err != nil {
		return instance, err
	}
	if _, exist := self.NodeStyling[instance.Selector]; exist {
		return instance, errors.New("This Node Styling Name: " + name + " Has Been Assigned. Please Use A Different One.")
	}
	self.NodeStyling[instance.Selector] = instance

	return self.NodeStyling[instance.Selector], nil
}

func (self *Diagram) AddEdge(From *diagram_node.Object, To *diagram_node.Object) {
	from := *From
	to := *To

	r := diagram_edge.NewObject(from.ID+"_"+to.ID, from.ID, to.ID)

	self.Edges[r.ID] = r
}
