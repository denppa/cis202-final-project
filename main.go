package main

import (
	"fmt"
	"os"

	"main/handle"
)

func main() {
	fmt.Println(`go run main.go mode arg
	mode: ls
	`)
	mode := os.Args[1]
	if mode == "ls" {
		fmt.Println(handle.Ls(os.Args[2]))
	}
}

