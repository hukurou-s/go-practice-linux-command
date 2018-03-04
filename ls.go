package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	recursivelyOpt       = flag.Bool("R", false, "Recursively list")
	nameBeginWithADotOpt = flag.Bool("a", false, "Display names begin with a dot")
	longFormatOpt        = flag.Bool("l", false, "Long format")
	sortOpt              = flag.Bool("S", false, "Sort by file size")
	reverseArrayOpt      = flag.Bool("r", false, "Reverse list")
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

	if *longFormatOpt {
		outputLongFormat(path)
	} else {
		outputFileList(path)
	}

}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isDir(path string) bool {
	Info, _ := os.Stat(path)
	return Info.IsDir()
}

func isBeginADot(name string) bool {
	return []rune(name)[0] == []rune(".")[0]
}

func sortByFileSize(files []os.FileInfo) {
	k := 1
	n := len(files)

	for k < n {
		if files[k-1].Size() >= files[k].Size() {
			k++
		} else {
			swap(files, k-1, k)
			k--
			if k == 0 {
				k++
			}
		}

	}
}

func reverseArray(files []os.FileInfo) {
	n := len(files)
	k := 0
	for k < n/2 {
		swap(files, k, n-k-1)
		k++
	}
}

func swap(files []os.FileInfo, i int, j int) {
	tmp := files[i]
	files[i] = files[j]
	files[j] = tmp
}

func outputFileList(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	if *sortOpt {
		sortByFileSize(files)
	}
	if *reverseArrayOpt {
		reverseArray(files)
	}

	for _, file := range files {
		if !*nameBeginWithADotOpt && isBeginADot(file.Name()) {
			continue
		}
		fmt.Printf("%s ", file.Name())
	}
	fmt.Print("\n\n")

	if !*recursivelyOpt {
		return
	}

	for _, file := range files {
		if !*nameBeginWithADotOpt && isBeginADot(file.Name()) {
			continue
		}
		//newpath := path + "/" + file.Name()
		newpath := filepath.Join(path, file.Name())
		if isDir(newpath) {
			fmt.Println(newpath)
			outputFileList(newpath)
		}
	}
}

func outputLongFormat(path string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	if *sortOpt {
		sortByFileSize(files)
	}

	if *reverseArrayOpt {
		reverseArray(files)
	}

	for _, file := range files {
		if !*nameBeginWithADotOpt && isBeginADot(file.Name()) {
			continue
		}

		// TODO : get file owner, group
		fmt.Printf("%s %8d %s %s\n", file.Mode(), file.Size(), file.ModTime(), file.Name())
	}

}
