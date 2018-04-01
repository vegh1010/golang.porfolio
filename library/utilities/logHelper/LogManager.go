package logHelper

import (
	"fmt"
	"os"
	"github.com/vegh1010/golang.porfolio/library/utilities/debug"
)

type LogManager struct {
	debugMode bool
	devMode   bool
}

//check environment variable DevMode to enable dev mode
func (self *LogManager) DevMode() {
	mode := os.Getenv("DevMode")
	self.devMode = true
	if mode == "true" {
		self.devMode = true
	}
}

//enable debug mode
func (self *LogManager) EnableDebug() {
	self.debugMode = true
}

//disable debug mode
func (self *LogManager) DisableDebug() {
	self.debugMode = false
}

//print error log
func (self *LogManager) ErrorLog(messages ...interface{}) {
	list := []interface{}{}
	for _, message := range messages {
		list = append(list, message)
	}
	list = append([]interface{}{"[Error]"}, list...)
	fmt.Println(list...)
	debug.Stack()
}

//print stack trace if debug mode enabled
func (self *LogManager) StackLog() {
	if self.debugMode {
		debug.Stack()
	}
}

//print log
func (self *LogManager) NormalLog(messages ...interface{}) {
	list := []interface{}{}
	for _, message := range messages {
		list = append(list, message)
	}
	list = append([]interface{}{"[Normal]"}, list...)
	fmt.Println(list...)
}

//print log if debug mode enabled
func (self *LogManager) DebugLog(messages ...interface{}) {
	if self.debugMode {
		list := []interface{}{}
		for _, message := range messages {
			list = append(list, message)
		}
		list = append([]interface{}{"[Debug]"}, list...)
		fmt.Println(list...)
	}
}

//print log if dev mode enabled
func (self *LogManager) DevDebugLog(messages ...interface{}) {
	if self.devMode {
		list := []interface{}{}
		for _, message := range messages {
			list = append(list, message)
		}
		list = append([]interface{}{"[Dev Debug]"}, list...)
		fmt.Println(list...)
	}
}
