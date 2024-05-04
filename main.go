package main

import (
	"fmt"
	"os"

	"main/handle"
)

func main() {
	fmt.Println(`go run main.go mode arg
	mode: ls "path/to/list"
	excel: excel "path/to/list" "save_file_name"
	`)
	mode := os.Args[1]
	if mode == "ls" {
		fmt.Println(handle.Ls(os.Args[2]))
	} else if mode == "excel" {
		handle.Excel(handle.Ls(os.Args[2]), os.Args[3])
	}

}

