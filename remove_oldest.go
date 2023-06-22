package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	directory := "/home/sfserver/lgsm/backup/"

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	if len(files) == 0 {
		fmt.Println("No files to remove in the directory")
		os.Exit(0)
	}

	var oldestFile os.FileInfo
	for _, fileInfo := range files {
		if oldestFile == nil || fileInfo.ModTime().Before(oldestFile.ModTime()) {
			oldestFile = fileInfo
		}
	}

	oldestPath := directory + "/" + oldestFile.Name()
	err = os.Remove(oldestPath)
	if err != nil {
		fmt.Println("Error removing file:", err)
		os.Exit(1)
	}

	fmt.Println("Removed oldest file:", oldestPath)
}
