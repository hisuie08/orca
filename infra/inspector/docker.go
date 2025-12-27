package inspector

import (
	"orca/ostools"
	"os/exec"
)

var _ NetworkInspector = (*dockerInspector)(nil)
var _ Volumenspector = (*dockerInspector)(nil)
var _ BindInspector = (*dockerInspector)(nil)
var (
	Docker = &dockerInspector{}
)

type NetworkInspector interface {
	NetworkExists(name string) bool
}
type Volumenspector interface {
	VolumeExists(name string) bool
}
type BindInspector interface {
	BindExists(dir string) bool
}

// DockerInspector 実装
type dockerInspector struct {
}

// VolumeExists docker volume inspect <name>
func (d dockerInspector) VolumeExists(name string) bool {
	cmd := exec.Command("docker", "volume", "inspect", name)
	if _, err := cmd.CombinedOutput(); err != nil {
		return false
	}
	return true
}

// NetworkExists docker network inspect <name>
func (d dockerInspector) NetworkExists(name string) bool {
	cmd := exec.Command("docker", "network", "inspect", name)
	if _, err := cmd.CombinedOutput(); err != nil {
		return false
	}
	return true
}

// BindExists ボリュームのマウント先確認用
func (f dockerInspector) BindExists(path string) bool {
	return ostools.DirExists(path)
}
