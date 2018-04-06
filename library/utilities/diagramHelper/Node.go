package diagramHelper

import "github.com/pkg/errors"

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

func NodeShape(value string) (data StyleOpinion) {
	if value == "" {
		value = Polygon
	}
	data = StyleOpinion{
		Tag: "shape",
		Value: value,
	}
	return
}

func NodeBackgroundColor(value string) (data StyleOpinion) {
	if value == "" {
		value = "white"
	}
	data = StyleOpinion{
		Tag: "background-color",
		Value: value,
	}
	return
}

func checkNodeShape(name string) (error) {
	switch name {
	case Ellipse, Triangle, Rectangle, Roundrectangle, BottomRoundRectangle, CutRectangle, Barrel, Rhomboid, Diamond, Pentagon,
		Hexagon, ConcaveHexagon, Heptagon, Octagon, Star, Tag, Vee, Polygon:
		return nil
	}

	return errors.New("Invalid Node Shape")
}