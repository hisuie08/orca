package filesystem

import (
	"fmt"
	"io/fs"
	"orca/errs"
	"orca/internal/capability"
	"orca/internal/logger"
	"orca/model/policy/log"
	"os"
	"path/filepath"
)

type executor interface {
	WriteFile(path string, data []byte) error
	CreateDir(path string) error
	RemoveFile(path string) error
	RemoveDir(path string) error
}
type execCapability interface {
	capability.WithPolicy
	capability.WithLog
}

var _ executor = (*fsExecutor)(nil)

type fsExecutor struct {
	caps   execCapability
	logger logger.Logger
}

func NewExecutor(caps execCapability) *fsExecutor {
	l := logger.New(caps)
	return &fsExecutor{caps: caps, logger: l}
}

func (f *fsExecutor) WriteFile(path string, data []byte) error {
	// create dir to path if not exist
	defer f.report(fmt.Sprintf("%s: %s", "create file", path))
	dir := filepath.Dir(path)
	if dir != "." {
		if err := f.createDir(dir, 0o755); err != nil {
			return err
		}
	}
	return f.writeFile(path, data, 0o644)
}

func (f *fsExecutor) CreateDir(path string) error {
	defer f.report(fmt.Sprintf("%s: %s", "create dir", path))
	return f.createDir(path, 0o755)
}

func (f *fsExecutor) RemoveFile(path string) error {
	defer f.report(fmt.Sprintf("%s: %s", "remove file", path))
	return f.removeFile(path)
}
func (f *fsExecutor) RemoveDir(path string) error {
	defer f.report(fmt.Sprintf("%s: %s", "remove dir", path))
	return f.removeDir(path)
}

func (f *fsExecutor) writeFile(path string, content []byte, perm fs.FileMode) error {
	if f.caps.Policy().AllowSideEffect() {
		if err := os.WriteFile(path, content, perm); err != nil {
			return &errs.FileError{Path: path, Err: err}
		}
	}
	return nil
}

func (f *fsExecutor) createDir(path string, perm fs.FileMode) error {
	if f.caps.Policy().AllowSideEffect() {
		if err := os.MkdirAll(path, perm); err != nil {
			return &errs.FileError{Path: path, Err: err}
		}
	}
	return nil
}

func (f *fsExecutor) removeFile(path string) error {
	if f.caps.Policy().AllowSideEffect() {
		if err := os.Remove(path); err != nil {
			return &errs.FileError{Path: path, Err: err}
		}
	}
	return nil
}
func (f *fsExecutor) removeDir(path string) error {
	if f.caps.Policy().AllowSideEffect() {
		if err := os.RemoveAll(path); err != nil {
			return &errs.FileError{Path: path, Err: err}
		}
	}
	return nil
}

func (f *fsExecutor) report(cmd string) {
	mode := "[DRY-RUN]"
	if f.caps.Policy().AllowSideEffect() {
		mode = "[RUN]"
	}
	msg := fmt.Sprintf("%s %s\n", mode, cmd)
	f.logger.Log(log.LogDetail, []byte(msg))
}
