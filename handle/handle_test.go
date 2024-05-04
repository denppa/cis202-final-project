package handle

import (
	"fmt"
	"testing"
)

func TestLsMode(t *testing.T) {
	output := Ls("../test/nonexistent, ../test/file1, ../test/file2, ../test/emptyFolder/, ../test/another/")
	snapshot := []VerboseFile{
			{Name: "../test/file1", Timestamp: 1524182400, FileSize: 4096}, 
			{Name: "../test/file2", Timestamp: 1524182400, FileSize: 4096},
			{Name: "../test/another/file3", Timestamp: 1714784952, FileSize: 4096},
			{Name: "../test/another/file4", Timestamp: 1714784957, FileSize: 4096},
	}

	for i := range output {
		if output[i] != snapshot[i] ||
		 output[i].Timestamp != snapshot[i].Timestamp ||
		 output[i].FileSize != snapshot[i].FileSize {
			fmt.Println(output[i])
			t.Fatalf("Does not match snapshot. \noutput: %#v \nsnapshot: %#v",
			output[i], snapshot[i],
			)
		}
	}
	fmt.Println(output)
}

// === RUN   TestLsMode
// Skipped ../test/nonexistent, os.Stat errored: stat ../test/nonexistent: no such file or directory.
// [{../test/file1 1524182400 4096} {../test/file2 1524182400 4096} {../test/another/file3 1714784952 4096} {../test/another/file4 1714784957 4096}]
// --- PASS: TestLsMode (0.00s)
// PASS
// ok      main/handle     0.088s

func TestExcel(t *testing.T) {
	Excel(Ls("../test/"), "../test-output/test.xlsx")
}

// Please see saved xlsx file for verification.

func TestPrepareExcelMvDel(t *testing.T) {
	Excel(Ls("../test2/"), "../test-output/test2.xlsx")
}

// Made the second excel file. Then I modified it, adding the actions columns and 2 actions to be performed.

func TestExcelMvDel(t *testing.T) {
	ExcelMvDel("../test-output/test2.xlsx")
}