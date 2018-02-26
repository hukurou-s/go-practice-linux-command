package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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

	files, err := ioutil.ReadDir(path)

	if err != nil {
		println(err)
		return
	}

	for _, file := range files {
		fmt.Printf("%s ", file.Name())
	}

	fmt.Println()

}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isDir(path string) bool {
	Info, _ := os.Stat(path)
	return Info.IsDir()
}
