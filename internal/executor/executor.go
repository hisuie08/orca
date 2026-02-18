package executor

import (
	"orca/internal/capability"
	"orca/internal/executor/internal/docker"
	"orca/internal/executor/internal/filesystem"
)

type Docker interface {
	ComposeUp(string) ([]byte, error)
	ComposeDown(string) ([]byte, error)
	CreateNetwork(string, ...string) ([]byte, error)
	CreateVolume(string, ...string) ([]byte, error)
}

type execCapability interface {
	capability.WithPolicy
	capability.WithLog
}

func NewDocker(p execCapability) Docker {
	return docker.NewExecutor(p)
}

type FileSystem interface {
	WriteFile(path string, data []byte) error
	CreateDir(path string) error
	RemoveFile(path string) error
	RemoveDir(path string) error
}

func NewFilesystem(p execCapability) FileSystem {
	return filesystem.NewExecutor(p)
}
