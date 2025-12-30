package filesystem

import (
	"io/fs"
	"orca/errs"
	"orca/internal/context"
	"orca/internal/policy"
	"os"
	"path/filepath"
)

type Executor interface {
	WriteFile(path string, data []byte) error
	CreateDir(path string) error
	RemoveFile(path string) error
	RemoveDir(path string) error
}

var _ Executor = (*executor)(nil)

type executor struct {
	context.WithPolicy
}

func NewExecutor(p policy.ExecPolicy) *executor {
	return &executor{WithPolicy: context.NewWithPolicy(p)}
}

func (f *executor) WriteFile(path string, data []byte) error {
	// create dir to path if not exist
	dir := filepath.Dir(path)
	if dir != "." {
		if err := f.createDir(dir, 0o755); err != nil {
			return err
		}
	}
	return f.writeFile(path, data, 0o644)
}

func (f *executor) CreateDir(path string) error {
	return f.createDir(path, 0o755)
}

func (f *executor) RemoveFile(path string) error {
	return f.removeFile(path)
}
func (f *executor) RemoveDir(path string) error {
	return f.removeDir(path)
}

func (f *executor) writeFile(p string, content []byte, perm fs.FileMode) error {
	if f.Policy().AllowSideEffect() {
		if err := os.WriteFile(p, content, perm); err != nil {
			return &errs.FileError{Path: p, Err: err}
		}
	}
	return nil
}

func (f *executor) createDir(path string, perm fs.FileMode) error {
	if f.Policy().AllowSideEffect() {
		if err := os.MkdirAll(path, perm); err != nil {
			return &errs.FileError{Path: path, Err: err}
		}
	}
	return nil
}

func (f *executor) removeFile(p string) error {
	if f.Policy().AllowSideEffect() {
		if err := os.Remove(p); err != nil {
			return &errs.FileError{Path: p, Err: err}
		}
	}
	return nil
}
func (f *executor) removeDir(p string) error {
	if f.Policy().AllowSideEffect() {
		if err := os.RemoveAll(p); err != nil {
			return &errs.FileError{Path: p, Err: err}
		}
	}
	return nil
}
