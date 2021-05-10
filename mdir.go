package mdir

import (
	"errors"
	"path/filepath"
)

type Cmd struct {
	CopyFile bool
	Force    bool
	Src      string
	Dest     string
	Segments []int

	_destRoot []string
}

func (cmd *Cmd) MvFiles() error {
	if cmd.Src == "" || cmd.Dest == "" {
		return errors.New("no src or dest")
	}

	list, err := listFiles(cmd.Src)
	if err != nil {
		return err
	}

	action := cmd.action()

	for _, file := range list {
		newPath := cmd.newPath(file)
		if err := action(file.path, newPath); err != nil {
			return err
		}
	}

	return nil
}

func (cmd *Cmd) action() func(src string, dest string) error {
	if !cmd.Force {
		return _dryRunFiles
	} else if cmd.CopyFile {
		return _cpFiles
	} else {
		return _mvFiles
	}
}

func (cmd *Cmd) destRoot() []string {
	if cmd._destRoot != nil {
		return cmd._destRoot
	}
	// add two more cap for root and file
	root := make([]string, 0, len(cmd.Segments)+2)
	root = append(root, cmd.Dest)
	cmd._destRoot = root
	return root
}

func (cmd *Cmd) newPath(file *fileInfo) string {
	md5path := md5Str(file.baseNameNoExt)
	dirs, _ := splitStr(md5path, cmd.Segments...)
	newDir := mkdirs(cmd.Force, append(cmd.destRoot(), dirs...)...)
	newPath := filepath.Join(newDir, file.baseName)
	return newPath
}
