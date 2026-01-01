package executor

import (
	"orca/internal/context"
	"orca/internal/executor/docker"
	"orca/internal/executor/filesystem"
)

type Docker interface {
	ComposeUp(string) (string, error)
	ComposeDown(string) (string, error)
	CreateNetwork(string, ...string) (string, error)
	CreateVolume(string, ...string) (string, error)
}

func NewDocker(p context.WithPolicy) Docker {
	return docker.NewExecutor(p)
}

type FileSystem interface {
	WriteFile(path string, data []byte) error
	CreateDir(path string) error
	RemoveFile(path string) error
	RemoveDir(path string) error
}

func NewFilesystem(p context.WithPolicy) FileSystem {
	return filesystem.NewExecutor(p)
}
