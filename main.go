package main

import (
	"log"
	"path/filepath"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	scanProjectDir(dir)
}

func scanProjectDir(dirName string) []string{
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
	}

	var fileList [] string
	for _, file := range files {
		fileList = append(fileList, dirName + string(os.PathSeparator) + file.Name())
		fmt.Println("append file:  ", dirName + string(os.PathSeparator) + file.Name())
		if file.IsDir() {
			retList := scanProjectDir(dirName + string(os.PathSeparator) + file.Name())
			fileList = append(fileList, retList...)
		}
	}
	
	return fileList
}

func formatPrintInvokeInfo(codeFileName string) {

}

func invokeMainFunc(codeFileName string) {

}
