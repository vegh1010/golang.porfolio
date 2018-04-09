package diagram_template

const HTML_LAYOUT = `
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
        elements: [{
            "classes": "queue",
            "data": {"id": "f0c53061-dc29-47ac-88ce-d495318dec9b", "name": "RabbitMQ"}
        }, {
            "classes": "database",
            "data": {"id": "436bb690-8c6d-4833-9de7-95ec620e6af7", "name": "CRMDB"}
        }, {
            "classes": "microservice",
            "data": {"id": "af776893-2686-4309-b929-45c0bbc46111", "name": "Monitoring\nMS"}
        }, {
            "classes": "microservice",
            "data": {"id": "ce17d6c1-0cc4-406d-8299-bc420d5253ca", "name": "VoiceMS"}
        }, {
            "classes": "microservice",
            "data": {"id": "9692d2c3-1be3-48cd-bcb4-89cc9d145dcc", "name": "WHMCSMS"}
        }, {
            "classes": "microservice",
            "data": {"id": "71082597-76e6-4edc-b56f-352becbd21e8", "name": "AuthWS"}
        }, {
            "classes": "worker",
            "data": {"id": "a6db059e-93c1-4cfe-8eef-d7342ad0ee2f", "name": "Worker"}
        }, {
            "classes": "database",
            "data": {"id": "0a5e2890-1eef-40ce-8fd2-b3ddad274ea1", "name": "GlobalDB"}
        }, {
            "classes": "microservice",
            "data": {"id": "35bbc6df-ef89-41f5-8ed6-2bc3f3b1e541", "name": "BillWS"}
        }, {
            "classes": "microservice",
            "data": {"id": "19b9db21-878d-4998-bc76-1f2d0e1de834", "name": "WorkerMS"}
        }, {
            "classes": "database",
            "data": {"id": "de705b4d-5cea-40eb-874d-31e1d1140978", "name": "Postgres"}
        }, {
            "classes": "microservice",
            "data": {"id": "51e9ed9f-41ea-49db-9115-968e0d19e970", "name": "CRMWS"}
        }, {
            "classes": "worker",
            "data": {"id": "65f55790-78cb-47b2-b633-c3c2103ed9a9", "name": "Data \nCollator"}
        }, {
            "classes": "microservice",
            "data": {"id": "38da104d-a102-4030-8337-8589df94c42f", "name": "BrandMS"}
        }, {
            "classes": "microservice",
            "data": {"id": "e64111bc-dd97-48f5-8605-96df860dedb8", "name": "AuditWS"}
        }, {
            "classes": "worker",
            "data": {"id": "7cda01a5-a946-4626-8977-e7b1b44a9672", "name": "DR \nWorker"}
        }, {
            "classes": "bezier",
            "data": {
                "id": "51e9ed9f-41ea-49db-9115-968e0d19e970_38da104d-a102-4030-8337-8589df94c42f",
                "source": "51e9ed9f-41ea-49db-9115-968e0d19e970",
                "target": "38da104d-a102-4030-8337-8589df94c42f"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "51e9ed9f-41ea-49db-9115-968e0d19e970_e64111bc-dd97-48f5-8605-96df860dedb8",
                "source": "51e9ed9f-41ea-49db-9115-968e0d19e970",
                "target": "e64111bc-dd97-48f5-8605-96df860dedb8"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "35bbc6df-ef89-41f5-8ed6-2bc3f3b1e541_51e9ed9f-41ea-49db-9115-968e0d19e970",
                "source": "35bbc6df-ef89-41f5-8ed6-2bc3f3b1e541",
                "target": "51e9ed9f-41ea-49db-9115-968e0d19e970"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "71082597-76e6-4edc-b56f-352becbd21e8_51e9ed9f-41ea-49db-9115-968e0d19e970",
                "source": "71082597-76e6-4edc-b56f-352becbd21e8",
                "target": "51e9ed9f-41ea-49db-9115-968e0d19e970"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "f0c53061-dc29-47ac-88ce-d495318dec9b_7cda01a5-a946-4626-8977-e7b1b44a9672",
                "source": "f0c53061-dc29-47ac-88ce-d495318dec9b",
                "target": "7cda01a5-a946-4626-8977-e7b1b44a9672"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "51e9ed9f-41ea-49db-9115-968e0d19e970_0a5e2890-1eef-40ce-8fd2-b3ddad274ea1",
                "source": "51e9ed9f-41ea-49db-9115-968e0d19e970",
                "target": "0a5e2890-1eef-40ce-8fd2-b3ddad274ea1"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "51e9ed9f-41ea-49db-9115-968e0d19e970_436bb690-8c6d-4833-9de7-95ec620e6af7",
                "source": "51e9ed9f-41ea-49db-9115-968e0d19e970",
                "target": "436bb690-8c6d-4833-9de7-95ec620e6af7"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "51e9ed9f-41ea-49db-9115-968e0d19e970_af776893-2686-4309-b929-45c0bbc46111",
                "source": "51e9ed9f-41ea-49db-9115-968e0d19e970",
                "target": "af776893-2686-4309-b929-45c0bbc46111"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "9692d2c3-1be3-48cd-bcb4-89cc9d145dcc_51e9ed9f-41ea-49db-9115-968e0d19e970",
                "source": "9692d2c3-1be3-48cd-bcb4-89cc9d145dcc",
                "target": "51e9ed9f-41ea-49db-9115-968e0d19e970"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "51e9ed9f-41ea-49db-9115-968e0d19e970_f0c53061-dc29-47ac-88ce-d495318dec9b",
                "source": "51e9ed9f-41ea-49db-9115-968e0d19e970",
                "target": "f0c53061-dc29-47ac-88ce-d495318dec9b"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "a6db059e-93c1-4cfe-8eef-d7342ad0ee2f_51e9ed9f-41ea-49db-9115-968e0d19e970",
                "source": "a6db059e-93c1-4cfe-8eef-d7342ad0ee2f",
                "target": "51e9ed9f-41ea-49db-9115-968e0d19e970"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "19b9db21-878d-4998-bc76-1f2d0e1de834_51e9ed9f-41ea-49db-9115-968e0d19e970",
                "source": "19b9db21-878d-4998-bc76-1f2d0e1de834",
                "target": "51e9ed9f-41ea-49db-9115-968e0d19e970"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "0a5e2890-1eef-40ce-8fd2-b3ddad274ea1_de705b4d-5cea-40eb-874d-31e1d1140978",
                "source": "0a5e2890-1eef-40ce-8fd2-b3ddad274ea1",
                "target": "de705b4d-5cea-40eb-874d-31e1d1140978"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "436bb690-8c6d-4833-9de7-95ec620e6af7_de705b4d-5cea-40eb-874d-31e1d1140978",
                "source": "436bb690-8c6d-4833-9de7-95ec620e6af7",
                "target": "de705b4d-5cea-40eb-874d-31e1d1140978"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "ce17d6c1-0cc4-406d-8299-bc420d5253ca_51e9ed9f-41ea-49db-9115-968e0d19e970",
                "source": "ce17d6c1-0cc4-406d-8299-bc420d5253ca",
                "target": "51e9ed9f-41ea-49db-9115-968e0d19e970"
            }
        }, {
            "classes": "bezier",
            "data": {
                "id": "65f55790-78cb-47b2-b633-c3c2103ed9a9_51e9ed9f-41ea-49db-9115-968e0d19e970",
                "source": "65f55790-78cb-47b2-b633-c3c2103ed9a9",
                "target": "51e9ed9f-41ea-49db-9115-968e0d19e970"
            }
        }],
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
