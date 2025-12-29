package inspector

import (
	"orca/ostools"
	"os/exec"
)

var _ DockerInspector = (*dockerInspector)(nil)

type DockerInspector interface {
	NetworkExists(name string) bool
	VolumeExists(name string) bool
	BindExists(dir string) bool
}

// DockerInspector 実装
type dockerInspector struct {
}

func NewInspector() DockerInspector {
	return &dockerInspector{}
}

// VolumeExists docker volume inspect <name>
func (d dockerInspector) VolumeExists(name string) bool {
	cmd := exec.Command("docker", "volume", "inspect", name)
	return d.runInspect(cmd)
}

// NetworkExists docker network inspect <name>
func (d dockerInspector) NetworkExists(name string) bool {
	cmd := exec.Command("docker", "network", "inspect", name)
	return d.runInspect(cmd)
}

// BindExists ボリュームのマウント先確認用
func (f dockerInspector) BindExists(path string) bool {
	return ostools.DirExists(path)
}

func (d *dockerInspector) runInspect(cmd *exec.Cmd) bool {
	if _, err := cmd.CombinedOutput(); err != nil {
		return false
	}
	return true
}
