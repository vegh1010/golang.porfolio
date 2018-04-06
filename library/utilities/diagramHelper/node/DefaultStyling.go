package diagram_node

import "fmt"

type DefaultStyling struct {
	Selector string
	Style    []*Option
}

func NewDefaultStyling() (*DefaultStyling) {
	var instance DefaultStyling

	instance = DefaultStyling{
		Selector: "node",
		Style: []*Option{
			NewShape(Polygon),
			NewBackgroundColor("grey"),
			NewLabel(),
			NewPadding(30),
			NewBorderWidth(1),
			NewFontWeight("bold"),
			NewTextVAlign("center"),
			NewTextWrap("wrap"),
		},
	}
	return &instance
}

func NewCustomDefaultStyling(opinionStyles ... *Option) (*DefaultStyling, error) {
	var instance DefaultStyling

	//map opinions
	mapList := map[string]*Option{}
	for _, data := range opinionStyles {
		mapList[data.Tag] = data
	}

	//check shape defined
	if data, exist := mapList["shape"]; !exist {
		mapList["shape"] = NewShape("")
	} else {
		//define default shape
		err := checkShape(fmt.Sprint(data.Value))
		if err != nil {
			return &instance, err
		}
	}

	//check background color defined
	if _, exist := mapList["background-color"]; !exist {
		//define default background color
		mapList["background-color"] = NewBackgroundColor("")
	}

	var list []*Option
	for _, data := range mapList {
		list = append(list, data)
	}

	instance = DefaultStyling{
		Selector: "node",
		Style:    list,
	}
	return &instance, nil
}
