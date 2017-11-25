package paginationHelper

import (
	"strconv"
)

func Format(params map[string]string, total int) (Pagination, error) {
	pagination := Pagination{}
	page, err := strconv.ParseInt(params["page"], 10, 0)
	if err != nil {
		return pagination, err
	}
	pageSize, err := strconv.ParseInt(params["perPage"], 10, 0)
	if err != nil {
		return pagination, err
	}

	pagination.Page = int(page)
	pagination.PageSize = int(pageSize)
	pagination.Total = total
	if value, exist := params["keyword"]; exist {
		pagination.Keyword = &value
	}
	return pagination, nil
}
