package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {

	flag.Parse()

	path := "./"

	if flag.Arg(0) != "" {
		path = flag.Arg(0)
	}

	files, err := ioutil.ReadDir(path)

	if err != nil {
		println(err)
	}

	for _, file := range files {
		fmt.Printf("%s ", file.Name())
	}
	fmt.Println()

}
