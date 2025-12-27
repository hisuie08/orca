package applier

import (
	"os/exec"
)

type NetworkCreator interface {
	CreateNetwork(name string, opt ...string) error
}

type VolumeCreator interface {
	CreateVolume(name string, opt ...string) error
}

type DockerApplier struct{}

func (d DockerApplier) CreateNetwork(name string, opt ...string) (string, error) {
	cmd := append([]string{"docker", "network", "create"}, opt...)
	return RunCommand(cmd...)
}
func (d DockerApplier) CreateVolume(name string, opt ...string) (string, error) {
	cmd := append([]string{"docker", "volume", "create"}, opt...)
	return RunCommand(cmd...)
}

func RunCommand(cmdline ...string) (string, error) {
	name, opts := cmdline[0], cmdline[1:]
	cmd := exec.Command(name, opts...)
	res, err := cmd.CombinedOutput()
	return string(res), err
}
