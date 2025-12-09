package ostools

import (
	"fmt"
	"os"
	"strings"
)

// パスからファル読み取り
func LoadFile(path string) ([]byte, error) {
	if data, err := os.ReadFile(path); err != nil {
		return nil, fmt.Errorf("file read error: %w", err)
	} else {
		return data, nil
	}
}

func pathExists(path string, expectDir bool) bool {
	if f, err := os.Stat(path); err == nil {
		return f.IsDir() == expectDir
	}
	return false
}

func DirExists(path string) bool {
	return pathExists(path, true)
}
func FileExisists(path string) bool {
	return pathExists(path, false)
}

// 現在のディレクトリ名を取得
func DirName(path string) string {
	split := strings.Split(path, "/")
	return split[len(split)-1]
}

// サブディレクトリを走査
func Directories(path string) []string {
	result := []string{}
	os.Chdir(path)
	cd, _ := os.Getwd()
	dir, _ := os.ReadDir(cd)
	for _, d := range dir {
		if d.IsDir() {
			result = append(result, cd+"/"+d.Name())
		}
	}
	return result
}

func ToFile(content []byte, path string) error {
	return os.WriteFile(path, content, 0775)
}
