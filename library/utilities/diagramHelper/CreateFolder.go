package diagramHelper

import "os"

func (self *Diagram) CreateFolder() (err error) {
	// write to file
	if self.FolderPath != "" {
		if _, err = os.Stat(self.FolderPath); os.IsNotExist(err) {
			//create directory if not exists
			err = os.MkdirAll(self.FolderPath, os.ModePerm)
			if err != nil {
				return
			}
		}
	}

	return
}