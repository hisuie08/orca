package inspector

import (
	"orca/ostools"
	"os/exec"
)

type NetworkInspector interface {
	NetworkExists(name string) bool
}

type VolumeInspector interface {
	VolumeExists(name string) bool
}

type FsInspector interface {
	DirExists(dir string) bool
}

// DockerInspector 実装
type DockerInspector struct {
}

// VolumeExists docker volume inspect <name>
func (d DockerInspector) VolumeExists(name string) bool {
	cmd := exec.Command("docker", "volume", "inspect", name)
	if _, err := cmd.CombinedOutput(); err != nil {
		return false
	}
	return true
}

// NetworkExists docker network inspect <name>
func (d DockerInspector) NetworkExists(name string) bool {
	cmd := exec.Command("docker", "network", "inspect", name)
	if _, err := cmd.CombinedOutput(); err != nil {
		return false
	}
	return true
}

// DirExists ボリュームのマウント先確認用
func (d DockerInspector) DirExists(path string) bool {
	return ostools.DirExists(path)
}
