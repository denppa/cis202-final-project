package handle

import (
	"fmt"
	"os"
	"strings"
)

type VerboseFile struct {
	Name string
	// Unix Timestamp.
	Timestamp int64
	// FileSize in bytes.
	FileSize int64

}

func Ls(paths string) []VerboseFile {
	// The paths string is empty, return.
	if len(paths) == 0 {
		return []VerboseFile{}
	}

	// Delclare return struct
	found := []VerboseFile{}

	// Turn comma separated paths string into a list
	list := strings.Split(paths, ",")
	// Trim heading trailing whitespace
	for i := range list {
		list[i] = strings.TrimSpace(list[i])
	}
	// m := fmt.Sprintf("Raw list of files: \n%#v\n", list)
	// fmt.Println(m)
	
	// If list only has one item, and that item is a dir,
	// read the files and use this as the list instead.
	first := list[0]
	firstInfo, err := os.Stat(first)
	if len(list) == 1 && err == nil && firstInfo.IsDir() {
		tmp, err := os.ReadDir(first)
		if err != nil {
			fmt.Println(err)
			return []VerboseFile{}
		}

		new := []string{}
		for i := range tmp {
			new =  append(new, first + tmp[i].Name())
		}
		list = new
	}

	// m2 := fmt.Sprintf("Parsed list of files: \n%#v\n", list)
	// fmt.Println(m2)

	for i := range list {
		info, err := os.Stat(list[i])
		if err != nil {
			// There was a problem reading this file, skip.
			m := fmt.Sprintf("Skipped %s, os.Stat errored: %s.", list[i], err)
			fmt.Println(m)
			continue
		}
		
		if info.IsDir() {
			found = append(found, Ls(list[i])...)
		} else {
			file := VerboseFile{
				Name: list[i],
				Timestamp: info.ModTime().Unix(),
				FileSize: info.Size(),
			}
			found = append(found, file)
		}
	}
	// fmt.Println(valid)
	return found
}