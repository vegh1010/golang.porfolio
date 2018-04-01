package documentHelper

func NewReportData(Text string) (data ReportData) {
	data.Text = Text
	return
}

func NewIfErrorReportData(isTransaction bool) (data ReportData) {
	data.Text = "If there's error,"
	if isTransaction {
		data.Text += " Rollback Transaction and"
	}
	data.Text += " Output Failed, Message: Error"
	return
}

func NewTransactionReportData() (init ReportData, commit ReportData) {
	init = NewReportData("Initialize Database Transaction")
	commit = NewReportData("Commit Database Transaction")
	return
}

func NewBodyReportData() (data ReportData) {
	data.Text = "Check Body"
	data.AddStep(NewReportData("If Body is nil return 400 error"))
	data.AddStep(NewReportData("Unmarshal Body to struct"))
	data.AddStep(NewIfErrorReportData(false))
	data.AddStep(NewReportData("Validate data in struct"))
	data.AddStep(NewReportData("If validation fails, Output Failed, Message: invalid fields"))
	return
}

type ReportData struct {
	Text     string
	Steps    []ReportData
}

func (self *ReportData) AddStep(step ReportData) {
	self.Steps = append(self.Steps, step)
}

func (self *ReportData) Output() (data string) {
	nestedTag := `ol`
	data += `<li>` + self.Text
	steps := self.Steps
	if len(steps) != 0 {
		data += `<` + nestedTag + `>`
		for i := 0; i < len(steps); i++ {
			data += steps[i].Output()
		}
		data += `</` + nestedTag + `>`
	}
	data += `</li>`
	return
}
