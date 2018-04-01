package routeHelper

import (
	"net/http"
	"fmt"
	"github.com/vegh1010/golang.porfolio/library/utilities/apiHelper"
)

type RouteDescription struct {
	Name         string      `json:"name"`
	Method       string      `json:"method"`
	Pattern      string      `json:"pattern"`
	Description  string      `json:"description"`
	AuthRequired bool        `json:"auth_required"`
	URLParams    interface{} `json:"url_params"`
	BodyParams   interface{} `json:"body_params"` //param: int, int32, int64, string, double, float, float32, float64, time, pointer, custom object, array
}

type Handler struct {
	list []RouteDescription
}

func (self *Handler) Init(List []Route) {
	for i := 0; i < len(List); i++ {
		self.list = append(self.list, List[i].Output())
	}
}

func (self *Handler) Help(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler.Help()")
	output := apiHelper.NewJsonOutput(w)

	output.Success = true
	output.Data = self.list
	output.Print()
}
