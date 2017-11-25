package paginationHelper

import (
    "net/http"
    "strconv"
)

func GetURLParameters(r *http.Request) (map[string]string) {
    params := make(map[string]string)
    params["page"] = "1"
    params["offset"] = "0"
    params["perPage"] = "20"

    if (r.URL.Query().Get("perPage") != "") {
        params["perPage"] = r.URL.Query().Get("perPage")
    }
    if (r.URL.Query().Get("keyword") != "") {
        params["keyword"] = r.URL.Query().Get("keyword")
    }
    if (r.URL.Query().Get("page") != "" && params["perPage"] != "-1") {
        page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 0)
        limit, _ := strconv.ParseInt(params["perPage"], 10, 0)
        skip := (page - 1) * limit
        params["page"] = r.URL.Query().Get("page")
        params["offset"] = strconv.Itoa(int(skip))
    }

    return params
}