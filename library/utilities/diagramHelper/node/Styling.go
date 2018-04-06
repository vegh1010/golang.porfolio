package diagram_node

import (
	"errors"
	"fmt"
)

type Styling struct {
	Selector string
	Style    []*Option
}

func NewStyling(name string, opinionStyles ... *Option) (*Styling, error) {
	var instance Styling
	if name == "" {
		return &instance, errors.New("Node Name Required")
	}

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

	instance = Styling{
		Selector: "node." + name,
		Style:    list,
	}
	return &instance, nil
}
