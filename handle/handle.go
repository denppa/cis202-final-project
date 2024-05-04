package handle

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type VerboseFile struct {
	Name string
	// Unix Timestamp.
	Timestamp int64
	// FileSize in bytes.
	FileSize int64
	Md5hash string

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
		// We know first is a dir now, append / if not exist,
		// ReadDir is omits trailing / even for dirs.
		end := len(first)-1
		if string(first[end]) != "/" {
			first += "/"
		}
		for i := range tmp {
			new =  append(new, first + tmp[i].Name())
		}
		list = new
	}

	// m2 := fmt.Sprintf("Parsed list of files: \n%#v\n", list)
	// fmt.Println(m2)

	// Now populate the files/dirs in list with verbose info.
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

func Excel(vf []VerboseFile, fileName string) {
    f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    // Create a new sheet.
    index, err := f.NewSheet("Files")
    if err != nil {
        fmt.Println(err)
        return
    }

	f.SetCellValue("Files", "A1", "File Name")
	f.SetCellValue("Files", "B1", "Timestamp")
	f.SetCellValue("Files", "C1", "File Size (Bytes)")

	for i := range vf {
		f.SetCellValue("Files", "A"+strconv.Itoa(i+2), vf[i].Name)
		f.SetCellValue("Files", "B"+strconv.Itoa(i+2), vf[i].Timestamp)
		f.SetCellValue("Files", "C"+strconv.Itoa(i+2), vf[i].FileSize)
	}

    // Set active sheet of the workbook.
    f.SetActiveSheet(index)
    // Save spreadsheet by the given path.
    if err := f.SaveAs(fileName); err != nil {
        fmt.Println(err)
    }
}

func ExcelMvDel(excelFile string) {
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	for i :=2;;i++{
		colFilePath :=  "A"+strconv.Itoa(i)
		colAction :=  "D"+strconv.Itoa(i)

		filePath, err := f.GetCellValue("Files", colFilePath)
		if len(filePath) == 0 || err != nil {
			fmt.Println("Looped through all the files, error: ", err)
			break
		}

		action, err := f.GetCellValue("Files", colAction)
		if len(action)==0 || err != nil {
			fmt.Println("Skipping ", colAction ,", no action to be performed, or error: ", err)
			continue
		}

		actions := strings.Split(action, " ")
		if actions[0] == "del" {
			err:=os.Remove(filePath)
			if err != nil {
				fmt.Println(err)
			}
		} else if actions[0] == "mv" {
			err:=os.Rename(filePath, actions[1])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			m:=fmt.Sprintf("Action %s not recognized, skipping.", actions[i])
			fmt.Println(m)
			continue
		}

	}
	cell, err := f.GetCellValue("Files", "D")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

// Accepts a path argument that points to a file.
func hash(path string) (string, error) {
	// This is an implementation that llm gave me, should be standard.
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString, nil
}

// Files are only considered to be dupes of each other if they have the same size, timestamp, and then hash. If their timpstamp is different, then it is not considered a dupe.
func LsDupes(paths string) []VerboseFile {
	dupes := []VerboseFile{}
	vf := Ls(paths)

	// Loop over every file in the list, check if there are any duplicates.
	for i := range vf {
		// The ones that dupe vf[i] will be stored in these
		these := []VerboseFile{}

		for j := range vf {
			if i==j {continue} // This is actually the same file, skip.

			if vf[i].FileSize == vf[j].FileSize &&
			vf[i].Timestamp == vf[j].Timestamp {
				m := fmt.Sprintf("%s has same size and timestamp as %s, hasing to verify", vf[i].Name, vf[j].Name)
				fmt.Println(m)

				hero, err := hash(vf[i].Name); if err != nil {
					m := fmt.Sprintf("Error hashing file %s, skipping to next.", vf[i].Name)
					fmt.Println(m)
					continue
				}
				vf[i].Md5hash = hero

				villain, err := hash(vf[j].Name) ; if err != nil {
					fmt.Println(err)
					continue
				}
				if hero  == villain {
					vf[j].Md5hash = villain
					these = append(these, vf[j])
				}
			}
		}
		// If the dupes for vf[i] has been found, they would be in these.
		if len(these) > 0 {
			these = append(these, vf[i])
			for i := range these {
				alreadyInList := false
				for j := range dupes {
					if these[i] == dupes[j] {
						alreadyInList = true
						break
					}
				}
				if !alreadyInList {
					dupes = append(dupes, these[i])
				}
			}
		}
	}

	return dupes
}