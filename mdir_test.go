package mdir

import (
	"path/filepath"
	"testing"
)

const srcDir = "/bin"

func TestMd5Str(t *testing.T) {
	goodMorningMd5 := md5Str("good morning")
	if goodMorningMd5 != "2b849500e4585dab4196ec9a415edf8f" {
		t.Error(goodMorningMd5)
	}
}

func TestSplitStr(t *testing.T) {
	str := "1234567890"

	result, err := splitStr(str, 1, 2, 3)
	if err != nil || result[0] != "1" || result[1] != "23" || result[2] != "456" {
		t.Error(result)
	}

	result1, err := splitStr(str, 10)
	if err != nil || result1[0] != str {
		t.Error(result1)
	}

	result2, err := splitStr(str, 11)
	if err == nil {
		t.Error(result2)
	}
}

func TestMkdirs(t *testing.T) {
	subDir := []string{srcDir, "a", "c"}
	path := mkdirs(false, subDir...)
	if path != filepath.Join(srcDir, "a", "c") {
		t.Error(path)
	}
}

func TestListFiles(t *testing.T) {
	list, err := listFiles(srcDir)
	if err != nil {
		t.Error(err)
	}
	for _, file := range list {
		t.Logf("%s -> %s\n", file.baseNameNoExt, file.oldPath)
	}
}

func TestCmd(t *testing.T) {
	cmd := Cmd{
		Src:      srcDir,
		Dest:     "./tmp/",
		Segments: []int{1},
		Progress: true,
		// CopyFile: true,
		// Force:    true,
	}
	if err := cmd.MvFiles(); err != nil {
		t.Error(err)
	}
}
