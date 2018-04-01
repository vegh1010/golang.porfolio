package paginationHelper

//format pagination output for list of items
type Pagination struct {
	Page     int     `json:"page"`     // current page number
	PageSize int     `json:"pageSize"` // number of items per page
	Total    int     `json:"total"`    // total number of items returned in the query, without pagination
	Keyword  *string `json:"keyword"`
}
