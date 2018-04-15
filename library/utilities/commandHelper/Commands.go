package commandHelper

import (
	"flag"
	"fmt"
	"github.com/apcera/termtables"
)

//command line library for developers
//use for quick help or other commands for development process
// such as generate list of available api in a microservice, check tools required to execute this code and so on
func NewCommands() (data Commands) {
	data.Init()
	return
}

type Commands struct {
	boolHelp *bool
	list     []CommandUsage
}

func (self *Commands) Init() {
	self.boolHelp = flag.Bool("help", false, "list commands available")
}

func (self *Commands) Add(Name, Usage string, Function func()) {
	var data CommandUsage
	data.Name = Name
	data.Usage = Usage
	data.IsBool = flag.Bool(Name, false, Usage)
	data.Function = Function
	self.list = append(self.list, data)
}

func (self *Commands) Process() (isCommand bool) {
	flag.Parse()

	if *self.boolHelp {
		isCommand = true
		MemoryTable := termtables.CreateTable()
		MemoryTable.AddHeaders(
			"Command",
			"Description",
		)
		for i := 0; i<len(self.list); i++ {
			data := self.list[i]
			MemoryTable.AddRow(
				"-" + data.Name,
				data.Usage,
			)
		}
		fmt.Println(MemoryTable.Render())
	} else {
		for i := 0; i<len(self.list); i++ {
			data := self.list[i]
			if *data.IsBool {
				isCommand = true
				data.Function()
				break
			}
		}
	}
	return
}

type CommandUsage struct {
	Name     string
	Usage    string
	IsBool   *bool
	Function func()
}
