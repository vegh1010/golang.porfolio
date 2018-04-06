package diagram_layout

const (
	LAYOUT_COSE = "cose"
	LAYOUT_RANDOM = "random"
	LAYOUT_GRID = "grid"
	LAYOUT_CIRCLE = "circle"
	LAYOUT_CONCENTRIC = "concentric"
	LAYOUT_BREADTH_FIRST = "breadthfirst"
)

type Styling struct {
	Name         string
	AvoidOverlap bool
}

func NewStyling(Name string, AvoidOverlap bool) (*Styling) {
	instance := Styling{
		Name:         Name,
		AvoidOverlap: AvoidOverlap,
	}
	return &instance
}

func NewDefaultStyling() (*Styling) {
	instance := Styling{
		Name:         LAYOUT_COSE,
		AvoidOverlap: true,
	}
	return &instance
}
