package requestHelper

const APPLICATION_ENCODED = "application/x-www-form-urlencoded"
const APPLICATION_JSON = "application/json"

func GET(urlString string, requestHeader, params map[string]string) (result RequestResult, err error) {
	var manager *RequestManager
	manager, err = NewRequestManager(urlString)
	if err != nil {
		return
	}
	manager.requestMethod = "GET"
	manager.contentType = APPLICATION_ENCODED
	manager.requestHeader = requestHeader
	manager.params = params
	result, err = manager.doRequest()

	return
}

func POST(urlString string, requestHeader map[string]string, data interface{}) (result RequestResult, err error) {
	var manager *RequestManager
	manager, err = NewRequestManager(urlString)
	if err != nil {
		return
	}
	err = manager.EncodeBody(data)
	if err != nil {
		return
	}
	manager.requestMethod = "POST"
	manager.contentType = APPLICATION_JSON
	manager.requestHeader = requestHeader
	result, err = manager.doRequest()

	return
}

func PUT(urlString string, requestHeader map[string]string, data interface{}) (result RequestResult, err error) {
	var manager *RequestManager
	manager, err = NewRequestManager(urlString)
	if err != nil {
		return
	}
	err = manager.EncodeBody(data)
	if err != nil {
		return
	}
	manager.requestMethod = "PUT"
	manager.contentType = APPLICATION_JSON
	manager.requestHeader = requestHeader
	result, err = manager.doRequest()

	return
}

func PATCH(urlString string, requestHeader map[string]string, data interface{}) (result RequestResult, err error) {
	var manager *RequestManager
	manager, err = NewRequestManager(urlString)
	if err != nil {
		return
	}
	err = manager.EncodeBody(data)
	if err != nil {
		return
	}
	manager.requestMethod = "PATCH"
	manager.contentType = APPLICATION_JSON
	manager.requestHeader = requestHeader
	result, err = manager.doRequest()

	return
}

func DELETE(urlString string, requestHeader map[string]string) (result RequestResult, err error) {
	var manager *RequestManager
	manager, err = NewRequestManager(urlString)
	if err != nil {
		return
	}
	manager.requestMethod = "DELETE"
	manager.contentType = APPLICATION_JSON
	manager.requestHeader = requestHeader
	result, err = manager.doRequest()

	return
}
