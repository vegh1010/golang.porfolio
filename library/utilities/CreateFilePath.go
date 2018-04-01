package utilities

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
 * Create File Path
 */
func CreateFilePath(pathToDirectory, filename string) (filePath string, err error) {
	fmt.Println("CreateFilePath()", filename)
	// write to file
	if pathToDirectory != "" && filename != "" {
		if _, err = os.Stat(pathToDirectory); os.IsNotExist(err) {
			//create directory if not exists
			err = os.MkdirAll(pathToDirectory, os.ModePerm)
			if err != nil {
				return
			}
		}
		filePath = filepath.Join(pathToDirectory, filename)
	}

	return
}