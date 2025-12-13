package ostools

import (
	"os"
)

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
