package diagramHelper

import "os"

//create into file
func (self *Diagram) Create() (err error) {
	if _, err = os.Stat(self.FilePath); os.IsNotExist(err) {
		_, err = os.Create(self.FilePath)
		if err != nil {
			return
		}
	}
	return
}
