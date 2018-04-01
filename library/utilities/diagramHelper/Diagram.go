package diagramHelper

import (
	"os"
	"io/ioutil"
	"github.com/vegh1010/golang.porfolio/library/utilities"
)

const (
	WORKER = "worker"
	DATABASE = "database"
	MICROSERVICE = "microservice"
	COLLECTOR = "collector"
	QUEUE = "queue"
	OTHER = "other"
	STANDALONE = "stand_alone"
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
		Elements: map[string]Element{},
		Edges: map[string]Edge{},
	}
}

type Diagram struct {
	FilePath string
	Filename string
	Title    string
	Elements map[string]Element
	Edges    map[string]Edge
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
	elements := []interface{}{}
	for _, element := range self.Elements {
		node := map[string]interface{}{}
		data := map[string]string{}
		data["id"] = element.ID
		data["name"] = element.Name
		node["data"] = data
		node["classes"] = element.Type
		elements = append(elements, node)
	}
	for _, edge := range self.Edges {
		node := map[string]interface{}{}
		data := map[string]string{}
		data["id"] = edge.ID
		data["source"] = edge.Source
		data["target"] = edge.Target
		node["data"] = data
		node["classes"] = "bezier"
		elements = append(elements, node)
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

func (self *Diagram) AddWorker(Name string) (Element) {
	return self.Add(Name, WORKER)
}

func (self *Diagram) AddDatabase(Name string) (Element) {
	return self.Add(Name, DATABASE)
}

func (self *Diagram) AddMicroservice(Name string) (Element) {
	return self.Add(Name, MICROSERVICE)
}

func (self *Diagram) AddCollector(Name string) (Element) {
	return self.Add(Name, COLLECTOR)
}

func (self *Diagram) AddQueue(Name string) (Element) {
	return self.Add(Name, QUEUE)
}

func (self *Diagram) AddOther(Name string) (Element) {
	return self.Add(Name, OTHER)
}

func (self *Diagram) AddStandAlone(Name string) (Element) {
	return self.Add(Name, STANDALONE)
}

func (self *Diagram) Add(Name, Type string) (Element) {
	var uuid string
	for {
		uuid = utilities.GetUuid().String()
		if _, exist := self.Elements[uuid]; !exist {
			break
		}
	}
	self.Elements[uuid] = Element{
		ID:   uuid,
		Name: Name,
		Type: Type,
	}
	return self.Elements[uuid]
}

func (self *Diagram) AddEdge(From Element, To Element) {
	r := Edge{
		ID:     From.ID + "_" + To.ID,
		Source: From.ID,
		Target: To.ID,
	}
	self.Edges[r.ID] = r
}
