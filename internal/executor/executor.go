package executor

import (
	"orca/internal/context"
	"orca/internal/executor/internal/docker"
	"orca/internal/executor/internal/filesystem"
)

type Docker interface {
	ComposeUp(string) ([]byte, error)
	ComposeDown(string) ([]byte, error)
	CreateNetwork(string, ...string) ([]byte, error)
	CreateVolume(string, ...string) ([]byte, error)
}

type execContext interface {
	context.WithPolicy
	context.WithLog
}

func NewDocker(p execContext) Docker {
	return docker.NewExecutor(p)
}

type FileSystem interface {
	WriteFile(path string, data []byte) error
	CreateDir(path string) error
	RemoveFile(path string) error
	RemoveDir(path string) error
}

func NewFilesystem(p execContext) FileSystem {
	return filesystem.NewExecutor(p)
}
