package main

import (
	"fmt"
	"./filestructure"
	"os"
)

func main() {
	fs, err := filestructure.NewFileStruct("root")
	check(err)
	fs.Print()

	fp, err := filestructure.NewDirectory("./dir1/dir2/dir2/\\")
	check(err)
	fmt.Println(fp)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
