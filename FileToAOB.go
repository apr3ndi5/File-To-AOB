package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

/*FileName file name to export*/
var FileName string = os.Args[1]

var file string = GetRealtivePath() + `\` + FileName
var path string = GetRealtivePath() + `\file.go`

func main() {
	FileBuffer := Read(file)
	CreateFile()
	WriteFile(FileBuffer)
}

/*GetRealtivePath rola que tal*/
func GetRealtivePath() (path string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(ex)
	return
}

/*Read file buffer*/
func Read(file string) (FileBuffer []byte) {
	FileBuffer, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return
}

/*CreateFile create file*/
func CreateFile() {

	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("[+] done creating file", path)
}

/*WriteFile write file*/
func WriteFile(p []byte) {

	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)

	mystring := fmt.Sprintf("%d", p)

	if isError(err) {
		return
	}
	defer file.Close()

	c := strings.ReplaceAll(mystring, " ", ", ")
	c = strings.Replace(c, "]", "}", 1)
	c = strings.Replace(c, "[", "var file := []bytes{", 1)

	// write into file
	_, err = file.WriteString(c)
	if isError(err) {
		return
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("[+] done writing to file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
