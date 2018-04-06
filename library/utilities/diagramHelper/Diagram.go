package diagramHelper

import (
	"os"
	"io/ioutil"
	"github.com/vegh1010/golang.porfolio/library/utilities"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/edge"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/layout"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/element"
)

//https://cdn.rawgit.com/cytoscape/cytoscape.js/master/dist/cytoscape.min.js
func NewDiagram(folder, Filename, Title string) (*Diagram) {
	var filePath = "./"
	if folder != "" {
		filePath += folder + "/"
	}
	filePath += Filename + ".html"

	return &Diagram{
		FilePath: filePath,
		Filename: Filename,
		Title:    Title,
		Nodes:    map[string]*diagram_node.Object{},
		Edges:    map[string]*diagram_edge.Object{},
	}
}

type Diagram struct {
	FilePath string
	Filename string
	Title    string

	Nodes              map[string]*diagram_node.Object
	DefaultNodeStyling *diagram_node.Styling
	NodeStyling        []*diagram_node.Styling

	Edges       map[string]*diagram_edge.Object
	EdgeStyling *diagram_edge.Styling

	Layout *diagram_layout.Styling
}

//write into file
func (self *Diagram) Create() (err error) {
	if _, err = os.Stat(self.FilePath); os.IsNotExist(err) {
		_, err = os.Create(self.FilePath)
		if err != nil {
			return
		}
	}
	return
}

//write into file
func (self *Diagram) Write(data string) error {
	err := ioutil.WriteFile(self.FilePath, []byte(""), os.ModeAppend)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(self.FilePath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

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

	var output string
	output, err = GetTemplate(self.Title, elements)
	if err != nil {
		return
	}
	err = self.Write(output)
	if err != nil {
		return
	}
	return
}

func (self *Diagram) Add(Name string, Type *diagram_node.Styling) (*diagram_node.Object) {
	var uuid string
	for {
		uuid = utilities.GetUuid().String()
		if _, exist := self.Nodes[uuid]; !exist {
			break
		}
	}
	self.Nodes[uuid] = diagram_node.NewObject(uuid, Name, Type.Selector)

	return self.Nodes[uuid]
}

func (self *Diagram) AddEdge(From diagram_node.Object, To diagram_node.Object) {
	r := diagram_edge.NewObject(From.ID+"_"+To.ID, From.ID, To.ID)

	self.Edges[r.ID] = r
}
