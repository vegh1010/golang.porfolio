package diagram_edge

type Styling struct {
	Selector string
	Style    []*Option
}

func NewDefaultBezierStyling() (*Styling) {
	instance := Styling{
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

//TODO: NewCustomDefaultBezierStyling
