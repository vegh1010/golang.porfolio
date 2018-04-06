package diagram_node

import "errors"

const (
	Ellipse = "ellipse"
	Triangle = "triangle"
	Rectangle = "rectangle"
	Roundrectangle = "roundrectangle"
	BottomRoundRectangle = "bottomroundrectangle"
	CutRectangle = "cutrectangle"
	Barrel = "barrel"
	Rhomboid = "rhomboid"
	Diamond = "diamond"
	Pentagon = "pentagon"
	Hexagon = "hexagon"
	ConcaveHexagon = "concavehexagon"
	Heptagon = "heptagon"
	Octagon = "octagon"
	Star = "star"
	Tag = "tag"
	Vee = "vee"
	Polygon = "polygon"
)

type Option struct {
	Tag   string
	Value interface{}
}

func NewShape(value string) (data *Option) {
	if value == "" {
		value = Polygon
	}
	data = &Option{
		Tag: "shape",
		Value: value,
	}
	return
}

func NewBackgroundColor(value string) (data *Option) {
	if value == "" {
		value = "white"
	}
	data = &Option{
		Tag: "background-color",
		Value: value,
	}
	return
}

func NewHeight(value string) (data *Option) {
	data = &Option{
		Tag: "height",
		Value: value,
	}
	return
}

func NewLabel() (data *Option) {
	data = &Option{
		Tag: "label",
		Value: "data(name)",
	}
	return
}

func NewPadding(value float64) (data *Option) {
	data = &Option{
		Tag: "padding",
		Value: value,
	}
	return
}

func NewBorderWidth(value float64) (data *Option) {
	data = &Option{
		Tag: "border-width",
		Value: value,
	}
	return
}

func NewFontWeight(value string) (data *Option) {
	data = &Option{
		Tag: "font-weight",
		Value: value,
	}
	return
}

func NewTextVAlign(value string) (data *Option) {
	data = &Option{
		Tag: "text-valign",
		Value: value,
	}
	return
}

func NewTextWrap(value string) (data *Option) {
	data = &Option{
		Tag: "text-wrap",
		Value: value,
	}
	return
}

func checkShape(name string) (error) {
	switch name {
	case Ellipse, Triangle, Rectangle, Roundrectangle, BottomRoundRectangle, CutRectangle, Barrel, Rhomboid, Diamond, Pentagon,
		Hexagon, ConcaveHexagon, Heptagon, Octagon, Star, Tag, Vee, Polygon:
		return nil
	}

	return errors.New("Invalid Node Shape")
}