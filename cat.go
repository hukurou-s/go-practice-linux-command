package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	//fmt.Println(flag.Arg(0))
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Could not open file", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 100)
	for {
		n, err := file.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Could not read file", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}

}
