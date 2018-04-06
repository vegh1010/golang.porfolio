package diagramHelper

type StyleOpinion struct {
	Tag   string
	Value interface{}
}

func CurveStyle(value string) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "curve-style",
		Value: value,
	}
	return
}

func TargetArrowShape(value string) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "target-arrow-shape",
		Value: value,
	}
	return
}

func LineColor(value string) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "line-color",
		Value: value,
	}
	return
}

func LineStyle(value string) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "line-style",
		Value: value,
	}
	return
}

func TargetArrowFill(value string) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "target-arrow-fill",
		Value: value,
	}
	return
}

func TargetArrowColor(value string) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "target-arrow-color",
		Value: value,
	}
	return
}

func ControlPointWeight(value float64) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "control-point-weight",
		Value: value,
	}
	return
}

func EdgeDistances(value string) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "edge-distances",
		Value: value,
	}
	return
}

func ArrowScale(value float64) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "arrow-scale",
		Value: value,
	}
	return
}

func ControlPointStepSize(value float64) (data StyleOpinion) {
	data = StyleOpinion{
		Tag: "control-point-step-size",
		Value: value,
	}
	return
}
