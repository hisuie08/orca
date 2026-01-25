package filesystem

import (
	"fmt"
	"io/fs"
	"orca/errs"
	"orca/internal/context"
	"orca/internal/logger"
	"os"
	"path/filepath"
)

type executor interface {
	WriteFile(path string, data []byte) error
	CreateDir(path string) error
	RemoveFile(path string) error
	RemoveDir(path string) error
}
type execContext interface {
	context.WithPolicy
	context.WithLog
}

var _ executor = (*fsExecutor)(nil)

type fsExecutor struct {
	ctx execContext
	log logger.Logger
}

func NewExecutor(ctx execContext) *fsExecutor {
	l := logger.New(ctx.LogTarget(), ctx.LogLevel())
	return &fsExecutor{ctx: ctx, log: l}
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
	if f.ctx.Policy().AllowSideEffect() {
		if err := os.WriteFile(path, content, perm); err != nil {
			return &errs.FileError{Path: path, Err: err}
		}
	}
	return nil
}

func (f *fsExecutor) createDir(path string, perm fs.FileMode) error {
	if f.ctx.Policy().AllowSideEffect() {
		if err := os.MkdirAll(path, perm); err != nil {
			return &errs.FileError{Path: path, Err: err}
		}
	}
	return nil
}

func (f *fsExecutor) removeFile(path string) error {
	if f.ctx.Policy().AllowSideEffect() {
		if err := os.Remove(path); err != nil {
			return &errs.FileError{Path: path, Err: err}
		}
	}
	return nil
}
func (f *fsExecutor) removeDir(path string) error {
	if f.ctx.Policy().AllowSideEffect() {
		if err := os.RemoveAll(path); err != nil {
			return &errs.FileError{Path: path, Err: err}
		}
	}
	return nil
}

func (f *fsExecutor) report(cmd string) {
	mode := "[DRY-RUN]"
	if f.ctx.Policy().AllowSideEffect() {
		mode = "[RUN]"
	}
	msg := fmt.Sprintf("%s %s\n", mode, cmd)
	f.log.Log(logger.LogDebug, []byte(msg))
}
