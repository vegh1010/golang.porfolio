package main

import (
	"fmt"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/edge"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/layout"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper"
	"github.com/Pallinder/go-randomdata"
	"time"
	"math/rand"
)

func GetDDiagram(folder, project, title string) (diagram *diagramHelper.Diagram, err error) {
	fmt.Println("GetDiagram()")

	diagram = diagramHelper.NewDiagram(
		folder,
		project,
		title,
		diagram_node.NewDefaultStyling(),
		diagram_edge.NewDefaultBezierStyling(),
		diagram_layout.NewDefaultStyling(),
	)

	var nodeStylings []*diagram_node.Styling
	nodeStylings, err = GetNodeStylings(diagram)
	if err != nil {
		return
	}

	GenerateNodeObjects(diagram, nodeStylings)

	return
}

func GetNodeStylings(diagram *diagramHelper.Diagram) (nodeStylings []*diagram_node.Styling, err error) {
	fmt.Println("GetNodeStylings()")

	DATABASE := new(diagram_node.Styling)
	MICROSERVICE := new(diagram_node.Styling)
	QUEUE := new(diagram_node.Styling)
	WORKER := new(diagram_node.Styling)
	COLLECTOR := new(diagram_node.Styling)
	OTHER := new(diagram_node.Styling)
	STANDALONE := new(diagram_node.Styling)

	//create node styling
	DATABASE, err = diagram.AddNodeStyling("database", NodeHeightOptions(diagram_node.RoundRectangle, "#F27E31")...)
	if err != nil {
		return
	}
	nodeStylings = append(nodeStylings, DATABASE)

	MICROSERVICE, err = diagram.AddNodeStyling("microservice", NodeOptions(diagram_node.Octagon, "#3473d8")...)
	if err != nil {
		return
	}
	nodeStylings = append(nodeStylings, MICROSERVICE)

	QUEUE, err = diagram.AddNodeStyling("queue", NodeHeightOptions(diagram_node.RoundRectangle, "#f744f4")...)
	if err != nil {
		return
	}
	nodeStylings = append(nodeStylings, QUEUE)

	WORKER, err = diagram.AddNodeStyling("worker", NodeOptions(diagram_node.Ellipse, "green")...)
	if err != nil {
		return
	}
	nodeStylings = append(nodeStylings, WORKER)

	COLLECTOR, err = diagram.AddNodeStyling("collector", NodeOptions(diagram_node.Tag, "#f751e9")...)
	if err != nil {
		return
	}
	nodeStylings = append(nodeStylings, COLLECTOR)

	OTHER, err = diagram.AddNodeStyling("other", NodeHeightOptions(diagram_node.RoundRectangle, "red")...)
	if err != nil {
		return
	}
	nodeStylings = append(nodeStylings, OTHER)

	STANDALONE, err = diagram.AddNodeStyling("stand_alone", NodeOptions(diagram_node.Polygon, "white")...)
	if err != nil {
		return
	}
	nodeStylings = append(nodeStylings, STANDALONE)

	return
}

func GenerateNodeObjects(diagram *diagramHelper.Diagram, nodeStylings []*diagram_node.Styling) {
	fmt.Println("GenerateNodeObjects()")

	nodeObjects := map[string]*diagram_node.Object{}
	var namess [][]string
	for index, styling := range nodeStylings {
		var names []string
		for i := 0; i < random(5, 10); i++ {
			for {
				var name = randomdata.SillyName()
				if _, exist := nodeObjects[name]; !exist {
					nodeObjects[name] = diagram.AddNode(name, styling)
					names = append(names, name)
					break
				}
			}
		}
		if index != 0 {
			nameS := namess[index-1]
			for z := 0; z < random(5, 10); z++ {
				for i := 0; i < len(names); i++ {
					nameS = shuffle(nameS)
					diagram.AddEdge(nodeObjects[nameS[0]], nodeObjects[names[i]])
				}
			}
		}
		namess = append(namess, names)
	}

	//for i := 0; i < 50; i++ {
	//	names = shuffle(names)
	//	node1 := names[0]
	//
	//	names = shuffle(names)
	//	node2 := names[0]
	//
	//	diagram.AddEdge(nodeObjects[node1], nodeObjects[node2])
	//}

	return
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func shuffle(slice []string) []string {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
