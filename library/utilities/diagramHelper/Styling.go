package diagramHelper

import (
	"github.com/pkg/errors"
	"fmt"
)

type Styling struct {
	Selector string
	Style    []StyleOpinion
}

//TODO: NewDefaultNode

func NewNode(name string, opinionStyles ... StyleOpinion) (*Styling, error) {
	var instance Styling
	if name == "" {
		return &instance, errors.New("Node Name Required")
	}

	//map opinions
	mapList := map[string]StyleOpinion{}
	for _, data := range opinionStyles {
		mapList[data.Tag] = data
	}

	//check shape defined
	if value, exist := mapList["shape"]; !exist {
		mapList["shape"] = NodeShape("")
	} else {
		//define default shape
		err := checkNodeShape(fmt.Sprint(value.Value))
		if err != nil {
			return &instance, err
		}
	}

	//check background color defined
	if _, exist := mapList["background-color"]; !exist {
		//define default background color
		mapList["background-color"] = NodeBackgroundColor("")
	}

	var list []StyleOpinion
	for _, data := range mapList {
		list = append(list, data)
	}

	instance = Styling{
		Selector: "node." + name,
		Style:    list,
	}
	return &instance, nil
}

func NewEdgeBezier() (*Styling) {
	instance := Styling{
		Selector: "edge.bezier",
		Style: []StyleOpinion{
			EdgeDistances("node-position"),
			CurveStyle("bezier"),
			TargetArrowShape("triangle"),
			TargetArrowFill("filled"),
			TargetArrowColor("black"),
			ArrowScale(1.5),
			LineColor("black"),
			LineStyle("solid"),
			ControlPointWeight(0.7),
			ControlPointStepSize(40),
		},
	}
	return &instance
}
