// Lesson_17
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var pathName string = "C:\\Reports"
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for t := range ticker.C {
			fmt.Printf("Time : %s\n", t)
			getFileList(pathName)
		}
	}()

	var enter string
	fmt.Println("Press Enter for Exit")
	fmt.Scanln(&enter)
	ticker.Stop()
}

func getFileList(pathName string) {
	fmt.Println("___", pathName, "___")
	filepath.Walk(pathName,
		func(path string, fileInfo os.FileInfo, err error) error {
			if !fileInfo.IsDir() {
				fmt.Printf("\t%s\t%d bytes\n", fileInfo.Name(), fileInfo.Size())
			}
			return nil
		})
	fmt.Println("____________________________________")
}
