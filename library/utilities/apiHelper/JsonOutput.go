package apiHelper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIResult struct {
	Success    bool        `json:"success"`
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

type jsonOutput struct {
	APIResult
	ResponseWriter http.ResponseWriter
}

func NewJsonOutput(w http.ResponseWriter) (out jsonOutput) {
	out.ResponseWriter = w
	return
}

func (out *jsonOutput) Print() {
	out.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	out.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(out.ResponseWriter).Encode(out.APIResult)
}

func (out *jsonOutput) PrintIf(print bool, customMessage interface{}) bool {
	if print {
		out.Success = false
		out.Message = customMessage
		out.Print()
	}
	return print
}

func (out *jsonOutput) PrintError(err error) {
	if err != nil {
		if out.Message == nil || fmt.Sprint(out.Message) == "" {
			out.Message = err.Error()
		}
	}
	out.Print()
}

func (out *jsonOutput) PrintErrorIf(err error, customMessage interface{}) (hasError bool) {
	if err != nil {
		hasError = true
		out.Message = customMessage
		out.PrintError(err)
	}
	return
}

func (out *jsonOutput) DownloadFile(bytes []byte, filename string) {
	if filename == "" {
		filename = "filename"
	}
	out.ResponseWriter.Header().Set("Content-Description", "File Transfer")
	out.ResponseWriter.Header().Set("Content-Disposition", "attachment; filename="+filename)
	out.ResponseWriter.Write(bytes)
}

func (out *jsonOutput) printErrorWithHttpStatusCode(err error, statusCode int) {
	if err != nil {
		// if message not specified use err.Error() as message
		if out.Message == nil || fmt.Sprint(out.Message) == "" {
			out.Message = err.Error()
		}
	}
	out.Message = fmt.Sprint(http.StatusText(statusCode)) + " " + fmt.Sprint(out.Message)
	http.Error(out.ResponseWriter, fmt.Sprint(out.Message), statusCode)
}

func (out *jsonOutput) PrintError400(err error) {
	out.printErrorWithHttpStatusCode(err, 400)
}

func (out *jsonOutput) PrintError500(err error) {
	out.printErrorWithHttpStatusCode(err, 500)
}
