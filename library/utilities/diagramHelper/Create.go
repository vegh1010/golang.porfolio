package diagramHelper

import (
	"os"
)

//create into file
func (self *Diagram) Create() (err error) {
	err = self.CreateFolder()
	if err != nil {
		return
	}
	if _, err = os.Stat(self.FilePath); os.IsNotExist(err) {
		_, err = os.Create(self.FilePath)
		if err != nil {
			return
		}
	}
	return
}
