package mdir

import (
	"os"
	"testing"
)

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
	wd, _ := os.Getwd()
	subDir := []string{wd, "a", "c"}
	path := mkdirs(false, subDir...)
	if path != wd+"/"+"a/c" {
		t.Error(path)
	}
}

func TestGenerateNamePathsMap(t *testing.T) {
	wd, _ := os.Getwd()
	namePaths, err := generateNamePathsMap(wd)
	if err != nil {
		t.Error(err)
	}
	for k, v := range namePaths {
		t.Logf("%s -> %s\n", k, v)
	}
}

func TestCmd(t *testing.T) {
	wd, _ := os.Getwd()
	cmd := Cmd{Src: wd, Dest: "/tmp/mdirtest", Segments: []int{2, 2, 2}}
	cmd.mvFiles()
}
