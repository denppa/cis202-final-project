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

// Please see that this worked from the git history. You will see one file being renamed/moved and another deleted from test2.

func TestLsDupes(t *testing.T) {
	dupes := LsDupes("../test")
	snapshot := []VerboseFile{
		{Name: "../test/file2", Timestamp: 1524182400, FileSize: 4096, Md5hash: "620f0b67a91f7f74151bc5be745b7110"}, 
		{Name: "../test/file1", Timestamp: 1524182400, FileSize: 4096, Md5hash: "620f0b67a91f7f74151bc5be745b7110"},
	}

	for i := range dupes {
		if dupes[i] != snapshot[i] {
			t.Fatalf("Does not match snapshot. \ndupes[i]: %#v \nsnapshot[i]: %#v",
			dupes[i], snapshot[i],
			)
		}
	}
	m := fmt.Sprintf("Duplicated stuff: \n%#v", dupes)
	fmt.Println(m)
}

// === RUN   TestLsDupes
// ../test/file1 has same size and timestamp as ../test/file2, hasing to verify
// ../test/file2 has same size and timestamp as ../test/file1, hasing to verify
// Duplicated stuff:
// []handle.VerboseFile{handle.VerboseFile{Name:"../test/file2", Timestamp:1524182400, FileSize:4096, Md5hash:"620f0b67a91f7f74151bc5be745b7110"}, handle.VerboseFile{Name:"../test/file1", Timestamp:1524182400, FileSize:4096, Md5hash:"620f0b67a91f7f74151bc5be745b7110"}}
// --- PASS: TestLsDupes (0.00s)
// PASS
// ok      main/handle     0.127s