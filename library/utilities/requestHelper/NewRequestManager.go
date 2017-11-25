package requestHelper

import (
	"net/url"
	"bytes"
	"errors"
)

func NewRequestManager(URL string) (*RequestManager, error) {
	var err error

	manager := RequestManager{}
	manager.requestUrl, err = url.ParseRequestURI(URL)
	if err != nil {
		return &manager, errors.New("Invalid URL supplied. URL: " + URL + ". Error: " + err.Error())
	}
	manager.requestHeader = map[string]string{}
	manager.params = map[string]string{}
	manager.dataBuffer = new(bytes.Buffer)

	return &manager, nil
}
