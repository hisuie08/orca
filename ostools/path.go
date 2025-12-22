package ostools

import (
	orca "orca/helper"
	"os"
	"path/filepath"
)

// パスの存在確認
func pathExists(path string, expectDir bool) bool {
	if f, err := os.Stat(path); err == nil {
		return f.IsDir() == expectDir
	}
	return false
}

// ディレクトリの存在確認
func DirExists(path string) bool {
	return pathExists(path, true)
}

// ファイルの存在確認
func FileExisists(path string) bool {
	return pathExists(path, false)
}

// ファイルとディレクトリを全走査
func fileDirs(path string, isDir bool) ([]string, error) {
	result := []string{}
	readDir, err := os.ReadDir(path)
	if err != nil {
		return nil, orca.OrcaError("directory read failed", err)
	}
	for _, d := range readDir {
		if d.IsDir() != isDir {
			continue
		}
		if abs, err := filepath.Abs(path); err != nil {
			return nil, orca.OrcaError("path resolve failed", err)
		} else {
			dir := filepath.Join(abs, d.Name())
			result = append(result, dir)
		}
	}
	return result, nil
}

// サブディレクトリを走査
func Directories(path string) ([]string, error) {
	return fileDirs(path, true)
}

// ディレクトリ内のファイルを走査
func Files(path string) ([]string, error) {
	return fileDirs(path, false)
}

// target の親ディレクトリを再帰的に作成
func ensureParentDir(target string) error {
	dir := filepath.Dir(target)
	return os.MkdirAll(dir, 0o755)
}

// - target がなければ再帰的に作成
// - content を書き込んで閉じる（既存なら上書き）
func CreateFile(target string, content []byte) error {
	if err := ensureParentDir(target); err != nil {
		return err
	}

	return os.WriteFile(target, content, 0o644)
}

func ReadFile(target string) ([]byte, error) {
	data, err := os.ReadFile(target)
	if err != nil {
		return nil, orca.OrcaError("file read error", err)
	}
	return data, nil
}

func CreateDir(target string) error {
	return os.MkdirAll(target, 0o755)
}

// - target がなければ再帰的に作成
// - ファイル末尾に「content+改行」追記
func AppendToFile(target string, content string) error {
	if err := ensureParentDir(target); err != nil {
		return err
	}

	f, err := os.OpenFile(
		target,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0o644,
	)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content + "\n")
	return err
}
