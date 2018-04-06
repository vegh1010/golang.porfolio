package main

import (
	"fmt"
	"github.com/vegh1010/golang.porfolio/library/utilities"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/edge"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/layout"
)

func main() {
	fmt.Println("Generating Treeview")
	var folder = "document"
	var project = "testing"

	err := utilities.CreateFolder(folder)
	check(err)

	diagram, err := GetDiagram(folder, project)
	check(err)

	fmt.Println("Generate()")
	err = diagram.Generate()
	check(err)
	fmt.Println("Diagram File Generated")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDiagram(folder, project string) (diagram *diagramHelper.Diagram, err error) {
	fmt.Println("GetDiagram()")

	//var Nodes *diagramHelper.DefaultNodeStyling
	//Nodes, err = diagramHelper.NewDefaultNodeStyling()
	//if err != nil {
	//	return
	//}

	fmt.Println("NewDiagram()")
	diagram = diagramHelper.NewDiagram(
		folder,
		project,
		"Testing Diagram",
		diagram_node.NewDefaultStyling(),
		diagram_edge.NewDefaultBezierStyling(),
		diagram_layout.NewDefaultStyling(),
	)

	DATABASE := new(diagram_node.Styling)
	MICROSERVICE := new(diagram_node.Styling)
	QUEUE := new(diagram_node.Styling)
	WORKER := new(diagram_node.Styling)

	DATABASE, err = diagram.AddNodeStyling("database", NodeHeightOptions(diagram_node.RoundRectangle, "#F27E31")...)
	if err != nil {
		return
	}
	MICROSERVICE, err = diagram.AddNodeStyling("microservice", NodeOptions(diagram_node.Octagon, "#3473d8")...)
	if err != nil {
		return
	}
	QUEUE, err = diagram.AddNodeStyling("queue", NodeHeightOptions(diagram_node.RoundRectangle, "#f744f4")...)
	if err != nil {
		return
	}
	WORKER, err = diagram.AddNodeStyling("worker", NodeOptions(diagram_node.Ellipse, "green")...)
	if err != nil {
		return
	}

	Postgres := diagram.AddNode("Postgres", DATABASE)
	GlobalDB := diagram.AddNode("GlobalDB", DATABASE)
	CRMDB := diagram.AddNode("CRMDB", DATABASE)

	CRMWS := diagram.AddNode("CRMWS", MICROSERVICE)
	BillingWS := diagram.AddNode("BillingWS", MICROSERVICE)
	MonitoringMS := diagram.AddNode("Monitoring\nMS", MICROSERVICE)
	BrandMS := diagram.AddNode("BrandMS", MICROSERVICE)
	VoiceMS := diagram.AddNode("VoiceMS", MICROSERVICE)
	TelstraMS := diagram.AddNode("TelstraMS", MICROSERVICE)
	WHMCSMS := diagram.AddNode("WHMCSMS", MICROSERVICE)
	AuthWS := diagram.AddNode("AuthWS", MICROSERVICE)
	AuditWS := diagram.AddNode("AuditWS", MICROSERVICE)

	DRRabbitMQ := diagram.AddNode("DR \nRabbitMQ", QUEUE)

	DRCRMWorker := diagram.AddNode("DR CRM \nWorker", WORKER)
	TelstraWorker := diagram.AddNode("Telstra \nWorker", WORKER)
	DataCollator := diagram.AddNode("Data \nCollator", WORKER)

	diagram.AddEdge(GlobalDB, Postgres)
	diagram.AddEdge(CRMDB, Postgres)
	diagram.AddEdge(CRMWS, GlobalDB)
	diagram.AddEdge(CRMWS, CRMDB)
	diagram.AddEdge(CRMWS, MonitoringMS)
	diagram.AddEdge(CRMWS, BrandMS)
	diagram.AddEdge(CRMWS, AuditWS)
	diagram.AddEdge(BillingWS, CRMWS)
	diagram.AddEdge(VoiceMS, CRMWS)
	diagram.AddEdge(TelstraMS, CRMWS)
	diagram.AddEdge(WHMCSMS, CRMWS)
	diagram.AddEdge(AuthWS, CRMWS)
	diagram.AddEdge(CRMWS, DRRabbitMQ)
	diagram.AddEdge(DRRabbitMQ, DRCRMWorker)
	diagram.AddEdge(TelstraWorker, CRMWS)
	diagram.AddEdge(DataCollator, CRMWS)

	return
}

func NodeOptions(shape, background string) (list []*diagram_node.Option) {
	list = append(
		list,
		diagram_node.NewShape(shape),
		diagram_node.NewBackgroundColor(background),
	)
	return
}

func NodeHeightOptions(shape, background string) (list []*diagram_node.Option) {
	list = NodeOptions(shape, background)
	list = append(
		list,
		diagram_node.NewHeight("1px"),
	)
	return
}
