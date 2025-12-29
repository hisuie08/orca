package fs

import (
	"orca/errs"
	"orca/internal/policy"
	"os"
	"path/filepath"
)

type FileWriter interface {
	Write(string, []byte) (string, error)
}

var _ FileWriter = (*fileWriter)(nil)

type fileWriter struct {
	policy policy.ExecPolicy
}

func NewFileWriter(p policy.ExecPolicy) *fileWriter {
	return &fileWriter{policy: p}
}

func (f *fileWriter) Write(target string, content []byte) (string, error) {
	path, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}
	return f.writeFile(path, content)
}

func (f *fileWriter) writeFile(p string, content []byte) (string, error) {
	if f.policy.AllowSideEffect() {
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			return p, &errs.FileError{Path: p, Err: err}
		}
		if err := os.WriteFile(p, content, 0o644); err != nil {
			return p, &errs.FileError{Path: p, Err: err}
		}
		return p, nil
	}
	return p, nil
}
