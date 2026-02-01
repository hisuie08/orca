package filesystem

import (
	"orca/errs"
	"os"
	"path/filepath"
)

var _ Inspector = (*inspector)(nil)

type Inspector interface {
	FileExists(string) bool
	DirExists(string) bool
	Dirs(string) ([]string, error)
	Files(string) ([]string, error)
	Read(string) ([]byte, error)
}

type inspector struct{}

func NewInspector() Inspector {
	return &inspector{}
}

// ディレクトリの存在確認
func (f *inspector) DirExists(path string) bool {
	return pathExists(path, true)
}

// ファイルの存在確認
func (f *inspector) FileExists(path string) bool {
	return pathExists(path, false)
}

// Dirs returns full paths of direct sub-directories under path.
func (f *inspector) Dirs(path string) ([]string, error) {
	return fileDirs(path, true)
}

// Files returns full paths of direct files under path.
func (f *inspector) Files(path string) ([]string, error) {
	return fileDirs(path, false)
}

// Read returns content in file path
func (f *inspector) Read(path string) ([]byte, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, &errs.FileError{Path: path, Err: err}
	}
	return b, nil
}

// pathExists pathの存在確認のみ; Statのエラーは「pathは存在しない」とみなす
func pathExists(path string, expectDir bool) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir() == expectDir
}

// ファイルとディレクトリを全走査
func fileDirs(path string, isDir bool) ([]string, error) {
	result := []string{}
	readDir, err := os.ReadDir(path)
	if err != nil {
		return nil, &errs.FileError{Path: path, Err: err}
	}
	for _, d := range readDir {
		if d.IsDir() != isDir {
			continue
		}
		dir := filepath.Join(path, d.Name())
		result = append(result, dir)
	}
	return result, nil
}
