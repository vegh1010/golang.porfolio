package documentHelper

import (
	"strings"
	"encoding/json"
)

func GetTemplate(Title string, Data interface{}) (string, error) {
	output := Template
	byteData, err := json.Marshal(Data)
	if err != nil {
		return output, err
	}
	output = strings.Replace(output, `%title%`, Title, -1)
	output = strings.Replace(output, `%data%`, string(byteData), -1)

	return output, nil
}

const Template = `
<!doctype html>
<html>
<head>
    <title>%title%</title>
    <script src="https://cdn.alloyui.com/3.0.1/aui/aui-min.js"></script>
	<link href="https://cdn.alloyui.com/3.0.1/aui-css/css/bootstrap.min.css" rel="stylesheet"></link>
</head>
<body>
<div id="treeview"></div>
<script>
	YUI().use(
	  'aui-tree-view',
	  function(Y) {
		var data = %data%;
	
		new Y.TreeView(
		  {
			boundingBox: '#treeview',
			children: data
		  }
		).render();
	  }
	);
</script>
</body>
</html>`
