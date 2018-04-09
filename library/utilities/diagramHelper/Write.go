package diagramHelper

import (
	"io/ioutil"
	"os"
)

//write into file
func (self *Diagram) Write(data string) error {
	err := ioutil.WriteFile(self.FilePath, []byte(""), os.ModeAppend)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(self.FilePath, os.O_WRONLY, os.ModeAppend)
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
