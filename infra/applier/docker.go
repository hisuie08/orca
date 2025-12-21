package applier

import (
	"orca/internal/ostools"
	"os/exec"
)

type NetworkApplier interface {
	CreateNetwork(name string, opt ...string) error
}

type VolumeApplier interface {
	CreateVolume(name string, opt ...string) error
}

type FsApplier interface {
	CreateDir(path string) error
}

type DockerApplier struct{}

func (d DockerApplier) CreateNetwork(name string, opt ...string) error {
	opts := append([]string{"network", "create"}, opt...)
	cmd := exec.Command("docker", opts...)
	_, err := cmd.CombinedOutput()
	return err
}
func (d DockerApplier) CreateVolume(name string, opt ...string) error {
	opts := append([]string{"volume", "create"}, opt...)
	cmd := exec.Command("docker", opts...)
	_, err := cmd.CombinedOutput()
	return err
}
func (d DockerApplier) CreateDir(path string) error {
	return ostools.CreateDir(path)
}
