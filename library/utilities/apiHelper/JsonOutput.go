package apiHelper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//format microservice output and some basic functionality for condition output
type APIOutput struct {
	Success    bool        `json:"success"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

type JsonOutput struct {
	APIOutput
	http.ResponseWriter
}

func NewJsonOutput(w http.ResponseWriter) (out JsonOutput) {
	out.ResponseWriter = w
	return
}

func (out *JsonOutput) Print() {
	out.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	out.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(out.ResponseWriter).Encode(out.APIOutput)
}

func (out *JsonOutput) PrintIf(print bool, customMessage interface{}) bool {
	if print {
		out.Success = false
		out.Message = customMessage
		out.Print()
	}
	return print
}

func (out *JsonOutput) PrintError(err error) {
	if err != nil {
		if out.Message == nil || fmt.Sprint(out.Message) == "" {
			out.Message = err.Error()
		}
	}
	out.Print()
}

func (out *JsonOutput) PrintErrorIf(err error, customMessage interface{}) (hasError bool) {
	if err != nil {
		hasError = true
		out.Message = customMessage
		out.PrintError(err)
	}
	return
}

func (out *JsonOutput) DownloadFile(bytes []byte, filename string) {
	if filename == "" {
		filename = "filename"
	}
	out.ResponseWriter.Header().Set("Content-Description", "File Transfer")
	out.ResponseWriter.Header().Set("Content-Disposition", "attachment; filename="+filename)
	out.ResponseWriter.Write(bytes)
}

func (out *JsonOutput) printErrorWithHttpStatusCode(err error, statusCode int) {
	if err != nil {
		// if message not specified use err.Error() as message
		if out.Message == nil || fmt.Sprint(out.Message) == "" {
			out.Message = err.Error()
		}
	}
	out.Message = fmt.Sprint(http.StatusText(statusCode)) + " " + fmt.Sprint(out.Message)
	http.Error(out.ResponseWriter, fmt.Sprint(out.Message), statusCode)
}

func (out *JsonOutput) PrintError400(err error) {
	out.printErrorWithHttpStatusCode(err, 400)
}

func (out *JsonOutput) PrintError500(err error) {
	out.printErrorWithHttpStatusCode(err, 500)
}
