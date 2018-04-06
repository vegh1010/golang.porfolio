package diagram_template

import (
	"strings"
	"encoding/json"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/element"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/node"
	"github.com/vegh1010/golang.porfolio/library/utilities/diagramHelper/edge"
)

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
        style: %style%,
        layout: %layout%
    });
</script>
</body>
</html>`
