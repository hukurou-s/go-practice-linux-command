package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	lineNumberOpt = flag.Bool("n", false, "Number the output lines")
)

func main() {

	flag.Parse()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Could not open file", err)
		return
	}
	defer file.Close()

	text := make([]byte, 0)
	buf := make([]byte, 100)

	for {
		n, err := file.Read(buf[0:])
		if err != nil {
			if err == io.EOF { // deep nest
				break
			}
			fmt.Println("Could not read file", err)
			return
		}
		text = append(text, buf[:n]...)
	}

	textArray := strings.Split(string(text), "\n")

	if *lineNumberOpt {
		for num, line := range textArray {
			fmt.Printf("%4d  %s\n", num, line)
		}
	} else {
		fmt.Print(string(text))
	}

}
