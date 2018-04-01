package documentHelper

func FormatPagination() (data ReportData) {
	data = NewReportData("Format Pagination Structure with params and total rows")
	data.AddStep(NewReportData("function " + Underline(Bold("paginationHelper.Format")) + "(params map[string]string, total int) (Pagination, error)"))
	data.AddStep(NewReportData("Process page and perPage value in params"))
	data.AddStep(NewReportData("Parse to Pagination Object"))
	data.AddStep(NewReportData("Return 2 results: pagination data and error"))

	return
}
