package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// handleMessages handles messages
func handleMessages(wi *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "init":
		payload, err = initApp(m.Payload)
		if err != nil {
			payload = err.Error()
			return
		}
		break
	case "generate_report":
		payload, err = GenerateReport(m.Payload)
		if err != nil {
			payload = err.Error()
			return
		}
		break
	}
	return
}
