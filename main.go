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
	excelMvDel: excelMvDel "with_actions_column.xlsx"
	`)
	mode := os.Args[1]
	if mode == "ls" {
		fmt.Println(handle.Ls(os.Args[2]))
	} else if mode == "excel" {
		handle.Excel(handle.Ls(os.Args[2]), os.Args[3])
	} else if mode == "excelMvDel" {
		handle.ExcelMvDel(os.Args[2])
	} else if mode == "LsDupes" {
		dupes := handle.LsDupes(os.Args[2])
		if !(len(dupes) > 0){
			fmt.Println("No dupes found. Hurray!")
		}
		for i := range dupes {
			fmt.Println("These files are dupes of each other:---------------------")
			fmt.Println(dupes[i])
		}
	}
}

