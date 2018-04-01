package routeHelper

import (
	"io/ioutil"
	"os"
	"encoding/json"
)

type DocumentGenerator struct {
	FILEPath string
	Routes   Routes
}

func (self *DocumentGenerator) Init(routes Routes, folder, filename string) (err error) {
	filePath := "./"
	if folder != "" {
		filePath += "/" + folder
	}
	self.FILEPath = filePath + "/" + filename + ".json"
	self.Routes = routes
	err = self.Create()
	if err != nil {
		return
	}
	err = ioutil.WriteFile(self.FILEPath, []byte(""), os.ModeAppend)
	if err != nil {
		return
	}
	err = self.Write()
	if err != nil {
		return
	}
	return
}

func (self *DocumentGenerator) Create() (error) {
	if _, err := os.Stat(self.FILEPath); os.IsNotExist(err) {
		_, err := os.Create(self.FILEPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (self *DocumentGenerator) Write() error {
	var list []RouteDescription
	for i := 0; i < len(self.Routes.list); i++ {
		list = append(list, self.Routes.list[i].Output())
	}
	f, err := os.OpenFile(self.FILEPath, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	out, err := json.Marshal(list)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(out) + "\n")
	if err != nil {
		return err
	}
	return nil
}
