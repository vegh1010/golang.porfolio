package documentHelper

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

func NewTreeview(Name string) (*Treeview) {
	Name = strings.Replace(Name, "Document", "", -1)
	return &Treeview{
		Name: Name,
	}
}

type Treeview struct {
	Name     string
	Children []*Treeview
}

func (self *Treeview) AddChild(data *Treeview) {
	self.Children = append(self.Children, data)
}

func (self *Treeview) Print() {
	data := self.Output(0)
	fmt.Print(data)
}

func (self *Treeview) Output(level int) (data string) {
	for i := 0; i < level; i++ {
		data += "| "
	}
	data += "|-" + self.Name + "\n"
	level++
	for i := 0; i < len(self.Children); i++ {
		data += self.Children[i].Output(level)
	}
	return
}

func (self *Treeview) View() (data map[string]interface{}) {
	return map[string]interface{}{
		"label":    self.Name,
		"expanded": true,
		"children": self.ViewChildren(),
	}
}

func (self *Treeview) ViewChildren() (children []interface{}) {
	for _, child := range self.Children {
		children = append(children, map[string]interface{}{
			"label":    child.Name,
			"expanded": true,
			"children": child.ViewChildren(),
		})
	}
	return
}

func (self *Treeview) Generate(folder, project, Title string) (err error) {
	err = self.Create(folder, project)
	if err != nil {
		return
	}

	var output string
	output, err = GetTemplate(Title, []interface{}{self.View()})
	if err != nil {
		return
	}
	err = self.Write(folder, project, output)
	if err != nil {
		return
	}
	return
}

func (self *Treeview) path(folder, project string) (filePath string) {
	filePath = "./"
	if folder != "" {
		filePath += folder + "/"
	}
	if project != "" {
		filePath += project + "_"
	}
	filePath += "treeview.html"

	return
}

//write into file
func (self *Treeview) Write(folder, project, data string) error {
	filePath := self.path(folder, project)
	err := ioutil.WriteFile(filePath, []byte(""), os.ModeAppend)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY, os.ModeAppend)
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

//write into file
func (self *Treeview) Create(folder, project string) (err error) {
	filePath := self.path(folder, project)
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		_, err = os.Create(filePath)
		if err != nil {
			return
		}
	}
	return
}
