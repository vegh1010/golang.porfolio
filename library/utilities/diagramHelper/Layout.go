package diagramHelper

const (
	LAYOUT_COSE = "cose"
	LAYOUT_RANDOM = "random"
	LAYOUT_GRID = "grid"
	LAYOUT_CIRCLE = "circle"
	LAYOUT_CONCENTRIC = "concentric"
	LAYOUT_BREADTH_FIRST = "breadthfirst"
)

type Layout struct {
	Name         string
	AvoidOverlap bool
}

func NewLayout(Name string, AvoidOverlap bool) (*Layout) {
	instance := Layout{
		Name:         Name,
		AvoidOverlap: AvoidOverlap,
	}
	return &instance
}
