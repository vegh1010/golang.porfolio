package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
 * Create File Path
 */
func CreateFilePath(pathToDirectory, filename string) (filePath string) {
	fmt.Println("CreateFilePath()", filename)
	// write to file
	if pathToDirectory != "" && filename != "" {
		if _, err := os.Stat(pathToDirectory); os.IsNotExist(err) {
			//create directory if not exists
			os.MkdirAll(pathToDirectory, os.ModePerm)
		}
		filePath = filepath.Join(pathToDirectory, filename)
	}

	return
}