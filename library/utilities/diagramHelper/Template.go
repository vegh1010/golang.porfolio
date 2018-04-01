package diagramHelper

import (
	"strings"
	"encoding/json"
)

func GetTemplate(Title string, Elements []interface{}) (string, error) {
	output := Template
	byteElements, err := json.Marshal(Elements)
	if err != nil {
		return output, err
	}
	output = strings.Replace(output, `%title%`, Title, -1)
	output = strings.Replace(output, `%elements%`, string(byteElements), -1)

	return output, nil
}

const Template = `
<!doctype html>
<html>
<head>
    <title>%title%</title>
    <script src="https://cdn.rawgit.com/cytoscape/cytoscape.js/master/dist/cytoscape.min.js"></script>
</head>
<style>
    #cy {
        width: 100%;
        height: 100%;
        position: absolute;
        top: 0px;
        left: 0px;
    }
</style>
<body>
<div id="cy"></div>
<script>
    var cy = cytoscape({
        container: document.getElementById('cy'),
        elements: %elements%,
        style: [
            {
                selector: 'node',
                style: {
                    shape: 'polygon',
                    'background-color': 'grey',
                    label: 'data(name)',
                    padding: 30,
                    'border-width': '1',
                    'font-weight': 'bold',
                    'text-valign': 'center',
            		'text-wrap': 'wrap'
                }
            },
            {
                selector: 'node.microservice',
                style: {
                    shape: 'octagon',
                    'background-color': '#3473d8'
                }
            },
            {
                selector: 'node.worker',
                style: {
                    shape: 'ellipse',
                    'background-color': 'green'
                }
            },
            {
                selector: 'node.database',
                style: {
                    shape: 'roundrectangle',
					height: '1px',
                    'background-color': '#F27E31'
                }
            },
            {
                selector: 'node.collector',
                style: {
                    shape: 'tag',
                    'background-color': '#f751e9'
                }
            },
            {
                selector: 'node.queue',
                style: {
                    shape: 'roundrectangle',
					height: '1px',
                    'background-color': '#f744f4'
                }
            },
            {
                selector: 'node.other',
                style: {
                    shape: 'roundrectangle',
					height: '1px',
                    'background-color': 'red'
                }
            },
            {
                selector: 'node.stand_alone',
                style: {
                    shape: 'polygon',
                    'background-color': '#ef2f2f'
                }
            },
            {
                selector: "edge.bezier",
                style: {
                    'curve-style': 'bezier',
                    'target-arrow-shape': 'triangle',
                    'line-color': 'black',
                    'line-style': 'solid',
                    'target-arrow-fill': 'filled',
                    'target-arrow-color': 'black',
                    'control-point-weight': 0.7,
                    'edge-distances': 'node-position',
                    'arrow-scale': 1.5,
                    "control-point-step-size": 40
                }
            }],
        layout: {
            name: 'cose',
            avoidOverlap: true
        }
    });
</script>
</body>
</html>`
