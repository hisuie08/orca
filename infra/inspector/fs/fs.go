package fs

import (
	"os"
	"path/filepath"
)

var _ FsInspector = (*fsInspector)(nil)

type FsInspector interface {
	FileExists(string) bool
	DirExists(string) bool
	Dirs(string) ([]string, error)
	Files(string) ([]string, error)
}

type fsInspector struct{}

var (
	NewInspector FsInspector = &fsInspector{}
)

// ディレクトリの存在確認
func (f *fsInspector) DirExists(path string) bool {
	return pathExists(path, true)
}

// ファイルの存在確認
func (f *fsInspector) FileExists(path string) bool {
	return pathExists(path, false)
}

// サブディレクトリを走査
func (f *fsInspector) Dirs(path string) ([]string, error) {
	return fileDirs(path, true)
}

// ディレクトリ内のファイルを走査
func (f *fsInspector) Files(path string) ([]string, error) {
	return fileDirs(path, false)
}

// パスの存在確認
func pathExists(path string, expectDir bool) bool {
	if f, err := os.Stat(path); err == nil {
		return f.IsDir() == expectDir
	}
	return false
}

// ファイルとディレクトリを全走査
func fileDirs(path string, isDir bool) ([]string, error) {
	result := []string{}
	readDir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, d := range readDir {
		if d.IsDir() != isDir {
			continue
		}
		if abs, err := filepath.Abs(path); err != nil {
			return nil, err
		} else {
			dir := filepath.Join(abs, d.Name())
			result = append(result, dir)
		}
	}
	return result, nil
}
