package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	recursivelyOpt = flag.Bool("R", false, "Recursively list")
)

func main() {

	flag.Parse()

	path := "./"
	if flag.Arg(0) != "" {
		path = flag.Arg(0)
	}

	if !isExist(path) {
		fmt.Println("No such file or directory")
		return
	}
	if !isDir(path) {
		fmt.Println(path)
		return
	}

	outputFileList(path)

}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isDir(path string) bool {
	Info, _ := os.Stat(path)
	return Info.IsDir()
}

func outputFileList(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		fmt.Printf("%s ", file.Name())
	}
	fmt.Print("\n\n")

	if !*recursivelyOpt {
		return
	}

	for _, file := range files {
		newpath := path + "/" + file.Name()
		if isDir(newpath) {
			fmt.Println(newpath)
			outputFileList(newpath)
		}
	}
}
