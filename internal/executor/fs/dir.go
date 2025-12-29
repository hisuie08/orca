package fs

import (
	"orca/internal/policy"
	"os"
	"path/filepath"
)

type DirCreator interface {
	CreateDir(path string) (string, error)
}

type dirCreator struct {
	policy policy.ExecPolicy
}


func (d *dirCreator) CreateDir(path string) (string, error) {
	t, err := filepath.Abs(path)
	if err != nil {
		return t, err
	}
	if !d.policy.AllowSideEffect() {
		return t, nil
	}
	if err := os.MkdirAll(t, 0o755); err != nil {
		return t, err
	}
	return t, nil
}

func (d *dirCreator) RemoveDir(path string) (string, error) {
	t, err := filepath.Abs(path)
	if err != nil {
		return t, err
	}
	if !d.policy.AllowSideEffect() {
		return t, nil
	}
	if err := os.RemoveAll(t); err != nil {
		return t, err
	}
	return t, nil
}
