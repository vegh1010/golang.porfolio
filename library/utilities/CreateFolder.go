package utilities

import (
	"fmt"
	"os"
)

/*
 * Create File Path
 */
func CreateFolder(pathToDirectory string) (err error) {
	fmt.Println("CreateFolder()")
	// write to file
	if pathToDirectory != "" {
		if _, err = os.Stat(pathToDirectory); os.IsNotExist(err) {
			//create directory if not exists
			err = os.MkdirAll(pathToDirectory, os.ModePerm)
			if err != nil {
				return
			}
		}
	}

	return
}