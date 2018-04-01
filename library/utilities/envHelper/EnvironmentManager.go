package envHelper

import (
	"fmt"
	"strings"
)

type EnvironmentManager struct {
	Logs []string
}

//if empty value, add to list
func (self *EnvironmentManager) Validate(tag, value string) {
	if value == "" {
		self.Logs = append(self.Logs, fmt.Sprint("environment variable required: ", tag))
	}
}

//panic if list is not empty
func (self *EnvironmentManager) Result() {
	if len(self.Logs) != 0 {
		panic(strings.Join(self.Logs, "\n"))
	}
}

func NewEnvironmentManager() *EnvironmentManager {
	manager := EnvironmentManager{}
	return &manager
}
