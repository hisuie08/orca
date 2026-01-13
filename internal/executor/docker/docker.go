package docker

import (
	"fmt"
	"orca/errs"
	"orca/internal/context"
	"os/exec"
)

type executor interface {
	ComposeUp(string) ([]byte, error)
	ComposeDown(string) ([]byte, error)
	CreateNetwork(string, ...string) ([]byte, error)
	CreateVolume(string, ...string) ([]byte, error)
}

type execContext interface {
	context.WithPolicy
	context.WithReport
}

var _ executor = (*dockerExecutor)(nil)

func NewExecutor(ctx execContext) executor {
	return &dockerExecutor{ctx: ctx}
}

type dockerExecutor struct {
	ctx execContext
}

func (d *dockerExecutor) ComposeUp(composeFile string) ([]byte, error) {
	cmd := exec.Command("docker", "compose", "-f", composeFile, "up", "-d")
	return d.run(cmd)
}

// docker compose down -f <file>
func (d *dockerExecutor) ComposeDown(composeFile string) ([]byte, error) {
	cmd := exec.Command("docker", "compose", "-f", composeFile, "down")
	return d.run(cmd)
}

// docker network create <name> [opt...]
func (d *dockerExecutor) CreateNetwork(name string, opt ...string) ([]byte, error) {
	c := append([]string{"network", "create", name}, opt...)
	cmd := exec.Command("docker", c...)
	return d.run(cmd)
}

// docker volume create <name> [opt...]
func (d *dockerExecutor) CreateVolume(name string, opt ...string) ([]byte, error) {
	c := append([]string{"volume", "create", name}, opt...)
	cmd := exec.Command("docker", c...)
	return d.run(cmd)
}

func (d *dockerExecutor) run(cmd *exec.Cmd) ([]byte, error) {
	mode := "[DRY-RUN]"
	msg := fmt.Sprintf("%s %s\n", mode, cmd.String())
	defer d.ctx.Report().Write([]byte(msg))
	if d.ctx.Policy().AllowSideEffect() {
		mode = "[RUN]"
		out, err := cmd.CombinedOutput()
		if err != nil {
			return out, &errs.ExternalError{Cmd: cmd.String(), Err: err}
		} else {
			return out, nil
		}
	}
	return []byte{}, nil
}
