package routeHelper

func NewParam(Required bool, Type, Description string) (param map[string]interface{}) {
	param = map[string]interface{}{}
	param["required"] = Required
	param["type"] = Type
	param["description"] = Description
	return
}

func NewParamObject(Required bool, Type string, Object map[string]interface{}, Description string) (param map[string]interface{}) {
	param = map[string]interface{}{}
	param["required"] = Required
	param["type"] = Type
	param["type_object"] = Object
	param["description"] = Description
	return
}

func NewParamPagination() (param map[string]interface{}) {
	param = map[string]interface{}{}
	param["perPage"] = NewParam(true, "integer", "number of rows")
	param["keyword"] = NewParam(false, "string", "Search by keyword")
	param["page"] = NewParam(true, "integer", "current page")
	return
}
