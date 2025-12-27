package ostools

import (
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
func FileExists(path string) bool {
	return pathExists(path, false)
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

// サブディレクトリを走査
func Dirs(path string) ([]string, error) {
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
func CreateFile(target string, content []byte, dry bool) (string, error) {
	if err := ensureParentDir(target); err != nil {
		return "", err
	}
	if !dry {
		if err := os.WriteFile(target, content, 0o644); err != nil {
			return "", nil
		}
	}
	return target, nil
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
