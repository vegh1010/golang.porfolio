package paginationHelper

func SQLPagination(params map[string]string) (query string) {
    if params["perPage"] != "-1" {
        query += " LIMIT " + params["perPage"] + " "
        query += " OFFSET " + params["offset"] + " "
    }
    return
}