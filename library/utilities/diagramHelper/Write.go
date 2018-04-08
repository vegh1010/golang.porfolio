package diagramHelper

import (
	"os"
)

//write into file
func (self *Diagram) Write(data string) error {
	f, err := os.OpenFile(self.FilePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
