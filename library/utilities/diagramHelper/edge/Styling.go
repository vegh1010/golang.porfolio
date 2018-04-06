package diagram_edge

type BezierStyling struct {
	Selector string
	Style    []*Option
}

func NewDefaultBezierStyling() (*BezierStyling) {
	instance := BezierStyling{
		Selector: "edge.bezier",
		Style: []*Option{
			NewEdgeDistances("node-position"),
			NewCurveStyle("bezier"),
			NewTargetArrowShape("triangle"),
			NewTargetArrowFill("filled"),
			NewTargetArrowColor("black"),
			NewArrowScale(1.5),
			NewLineColor("black"),
			NewLineStyle("solid"),
			NewControlPointWeight(0.7),
			NewControlPointStepSize(40),
		},
	}
	return &instance
}

func NewBezierStyling(opinionStyles ... *Option) (*BezierStyling, error) {
	var instance BezierStyling

	//map opinions
	mapList := map[string]*Option{}
	for _, data := range opinionStyles {
		mapList[data.Tag] = data
	}

	var list []*Option
	for _, data := range mapList {
		list = append(list, data)
	}

	instance = BezierStyling{
		Selector: "edge.bezier",
		Style:    list,
	}
	return &instance, nil
}
