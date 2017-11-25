package requestHelper

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RequestResult struct {
	ResponseBody []byte
	StatusCode   int
	StatusString string
}

type RequestManager struct {
	requestUrl    *url.URL
	requestMethod string
	contentType   string
	requestHeader map[string]string
	params        map[string]string
	dataBuffer    *bytes.Buffer
}

func (self *RequestManager) ValidateURL(urlString string) (err error) {
	self.requestUrl, err = url.ParseRequestURI(urlString)
	if err != nil {
		return
	}
	return
}

func (self *RequestManager) EncodeBody(data interface{}) (err error) {
	err = json.NewEncoder(self.dataBuffer).Encode(data)
	if err != nil {
		return
	}
	return
}

func (self *RequestManager) doRequest() (result RequestResult, err error) {
	var responseBody []byte
	var statusCode int
	var statusString string

	//add parameters to url
	parameters := url.Values{}
	for key, value := range self.params {
		parameters.Add(key, value)
	}
	self.requestUrl.RawQuery = parameters.Encode()

	// replace any spaces with %20 to avoid failures
	urlString := strings.Replace(self.requestUrl.String(), " ", "%20", -1)

	//new instance request
	newRequest, err := http.NewRequest(self.requestMethod, urlString, self.dataBuffer)
	if err != nil {
		return
	}

	// Set request content types
	newRequest.Header.Set("Content-Type", self.contentType)

	//add request header data
	for key, value := range self.requestHeader {
		newRequest.Header.Set(key, value)
	}

	//create a new http client and execute http request
	newClient := &http.Client{}
	newResponse := new(http.Response)
	newResponse, err = newClient.Do(newRequest)
	if err != nil {
		err = errors.New("Failed to connect: " + err.Error())
		return
	}
	defer newResponse.Body.Close()

	statusCode = newResponse.StatusCode
	statusString = newResponse.Status
	responseBody, err = ioutil.ReadAll(newResponse.Body)
	if err != nil {
		return
	}

	result.StatusCode = statusCode
	result.StatusString = statusString
	result.ResponseBody = responseBody

	return
}
