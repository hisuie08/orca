package executor

import (
	"orca/internal/executor/docker"
	"orca/internal/executor/filesystem"
	"orca/model/policy"
)

type Docker interface {
	ComposeUp(string) (string, error)
	ComposeDown(string) (string, error)
	CreateNetwork(string, ...string) (string, error)
	CreateVolume(string, ...string) (string, error)
}

func NewDocker(p policy.ExecPolicy) Docker {
	return docker.NewExecutor(p)
}

type FileSystem interface {
	WriteFile(path string, data []byte) error
	CreateDir(path string) error
	RemoveFile(path string) error
	RemoveDir(path string) error
}

func NewFilesystem(p policy.ExecPolicy) FileSystem {
	return filesystem.NewExecutor(p)
}
