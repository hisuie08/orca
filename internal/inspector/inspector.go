package inspector

import (
	"orca/internal/inspector/docker"
	"orca/internal/inspector/filesystem"
	"orca/model/compose"
)

type Docker interface {
	NetworkExists(name string) bool
	VolumeExists(name string) bool
	BindExists(dir string) bool
	Compose(dir string) (*compose.ComposeSpec, error)
}

func NewDocker() Docker {
	return docker.NewInspector()
}

type FileSystem interface {
	FileExists(string) bool
	DirExists(string) bool
	Dirs(string) ([]string, error)
	Files(string) ([]string, error)
	Read(string) ([]byte, error)
}

func NewFilesystem() FileSystem {
	return filesystem.NewInspector()
}
