package diagram_edge

type Option struct {
	Tag   string
	Value interface{}
}

func NewCurveStyle(value string) (data *Option) {
	data = &Option{
		Tag: "curve-style",
		Value: value,
	}
	return
}

func NewTargetArrowShape(value string) (data *Option) {
	data = &Option{
		Tag: "target-arrow-shape",
		Value: value,
	}
	return
}

func NewLineColor(value string) (data *Option) {
	data = &Option{
		Tag: "line-color",
		Value: value,
	}
	return
}

func NewLineStyle(value string) (data *Option) {
	data = &Option{
		Tag: "line-style",
		Value: value,
	}
	return
}

func NewTargetArrowFill(value string)  (data *Option) {
	data = &Option{
		Tag: "target-arrow-fill",
		Value: value,
	}
	return
}

func NewTargetArrowColor(value string) (data *Option) {
	data = &Option{
		Tag: "target-arrow-color",
		Value: value,
	}
	return
}

func NewControlPointWeight(value float64) (data *Option) {
	data = &Option{
		Tag: "control-point-weight",
		Value: value,
	}
	return
}

func NewEdgeDistances(value string) (data *Option) {
	data = &Option{
		Tag: "edge-distances",
		Value: value,
	}
	return
}

func NewArrowScale(value float64) (data *Option) {
	data = &Option{
		Tag: "arrow-scale",
		Value: value,
	}
	return
}

func NewControlPointStepSize(value float64) (data *Option) {
	data = &Option{
		Tag: "control-point-step-size",
		Value: value,
	}
	return
}
