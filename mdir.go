package mdir

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Cmd struct {
	CopyFile bool
	Force    bool
	Src      string
	Dest     string
	Segments []int
}

func md5Str(fileName string) string {
	hash := md5.Sum([]byte(fileName))
	return fmt.Sprintf("%x", hash)
}

func splitStr(str string, lens ...int) ([]string, error) {
	start := 0
	limit := len(str)
	slice := make([]string, 0, len(lens))

	for _, l := range lens {
		end := start + l
		if end > limit {
			return nil, errors.New("1")
		}
		subStr := str[start : start+l]
		slice = append(slice, subStr)
		start = end
	}

	return slice, nil
}

func mkdirs(force bool, dirs ...string) string {
	path := filepath.Join(dirs...)
	if force {
		os.MkdirAll(path, os.ModePerm)
	}
	return path
}

type baseNameNoExtSet map[string]void
type void struct{}

var baseNameNoExtSetMember void

type fileList []*fileInfo
type fileInfo struct {
	path          string
	baseName      string
	baseNameNoExt string
}

func listFiles(dir string) (fileList, error) {
	npm := make(baseNameNoExtSet)
	list := make(fileList, 0, 0)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		baseName := filepath.Base(info.Name())
		baseNameNoExt := strings.TrimSuffix(baseName, filepath.Ext(baseName))
		if _, exists := npm[baseNameNoExt]; exists {
			return errors.New("duplicate " + baseNameNoExt)
		}
		npm[baseNameNoExt] = baseNameNoExtSetMember
		list = append(list, &fileInfo{path, baseName, baseNameNoExt})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return list, nil
}

func _mvFiles(src string, dest string) {

}

func _cpFiles(src string, dest string) {

}

func _dryRunFiles(src string, dest string) {
	fmt.Printf("%s -> %s\n", src, dest)
}

func (cmd *Cmd) mvFiles() error {
	if cmd.Src == "" || cmd.Dest == "" {
		return errors.New("no src or dest")
	}

	var action func(src string, dest string)
	if !cmd.Force {
		action = _dryRunFiles
	} else if cmd.CopyFile {
		action = _cpFiles
	} else {
		action = _mvFiles
	}

	list, err := listFiles(cmd.Src)
	if err != nil {
		return err
	}

	destRoot := cmd.destRoot()

	for _, file := range list {
		md5path := md5Str(file.baseNameNoExt)
		dirs, _ := splitStr(md5path, cmd.Segments...)
		newPath := mkdirs(cmd.Force, dirs...)
		newPaths := append(destRoot, newPath, file.baseName)
		newPath = filepath.Join(newPaths...)
		action(file.path, newPath)
	}

	return nil
}

func (cmd *Cmd) destRoot() []string {
	// add two more cap for root and file
	root := make([]string, 0, len(cmd.Segments)+2)
	root = append(root, cmd.Dest)
	return root
}
