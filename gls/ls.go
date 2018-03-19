package main

import (
	"flag"
	"fmt"
	"github.com/hukurou-s/go-command/gls/arguments"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

var (
	R *bool
	a *bool
	l *bool
	S *bool
	r *bool
)

func init() {
	R = flag.Bool("R", false, "Recursively list")
	a = flag.Bool("a", false, "Display names begin with a dot")
	l = flag.Bool("l", false, "Long format")
	S = flag.Bool("S", false, "Sort by file size")
	r = flag.Bool("r", false, "Reverse list")
}

func main() {
	flag.Parse()
	options := arguments.GetFlags(R, a, l, S, r)

	path := getPath()

	if !isExist(path) {
		fmt.Println("No such file or directory")
		return
	}
	if !isDir(path) {
		fmt.Println(path)
		return
	}

	if options.LongFormatOpt() {
		outputLongFormat(path, options)
	} else {
		outputFileList(path, options)
	}

}

func getPath() string {
	if flag.Arg(0) != "" {
		return flag.Arg(0)
	}

	return "./"
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

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() >= files[j].Size()
	})
}

func reverseArray(files []os.FileInfo) {
	for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
		files[i], files[j] = files[j], files[i]
	}
}

func outputFileList(path string, opts *arguments.CommandOpts) {

	files := readDirectory(path, opts)

	for _, file := range files {
		if !opts.NameBeginWithADotOpt() && isBeginADot(file.Name()) {
			continue
		}
		fmt.Printf("%s ", file.Name())
	}
	fmt.Print("\n\n")

	if !opts.RecursivelyOpt() {
		return
	}

	for _, file := range files {
		if !opts.NameBeginWithADotOpt() && isBeginADot(file.Name()) {
			continue
		}
		newpath := filepath.Join(path, file.Name())
		if isDir(newpath) {
			fmt.Println(newpath)
			outputFileList(newpath, opts)
		}
	}
}

func outputLongFormat(path string, opts *arguments.CommandOpts) {

	files := readDirectory(path, opts)

	for _, file := range files {
		if !opts.NameBeginWithADotOpt() && isBeginADot(file.Name()) {
			continue
		}

		// TODO : get file owner, group
		fmt.Printf("%s %8d %s %s\n", file.Mode(), file.Size(), file.ModTime(), file.Name())
	}

}

func readDirectory(path string, opts *arguments.CommandOpts) []os.FileInfo {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.SortOpt() {
		sortByFileSize(files)
	}

	if opts.ReverseArrayOpt() {
		reverseArray(files)
	}

	return files
}
