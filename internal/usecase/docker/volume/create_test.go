package volume

import (
	"orca/model/plan"
	"os/exec"
	"testing"
)

var _ dockerExecutor = (*fakeExecutor)(nil)

type fakeExecutor struct {
}

func (f *fakeExecutor) ComposeUp(name string) ([]byte, error) {
	cmd := exec.Command("docker compose -f " + name + " up -d")
	return f.runFake(cmd)
}
func (f *fakeExecutor) ComposeDown(name string) ([]byte, error) {
	cmd := exec.Command("docker compose -f " + name + " down")
	return f.runFake(cmd)
}
func (f *fakeExecutor) CreateNetwork(name string, opts ...string) ([]byte, error) {
	c := append(append([]string{"network", "create"}, opts...), name)
	cmd := exec.Command("docker", c...)
	return f.runFake(cmd)
}
func (f *fakeExecutor) CreateVolume(name string, opts ...string) ([]byte, error) {
	c := append(append([]string{"volume", "create"}, opts...), name)
	cmd := exec.Command("docker", c...)
	return f.runFake(cmd)
}

func (f *fakeExecutor) runFake(cmd *exec.Cmd) ([]byte, error) {
	return []byte(cmd.String()), nil
}
func Test_createVolume(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		caps createVolCapability
		vp   plan.VolumePlan
		de   dockerExecutor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createVolume(tt.caps, tt.vp, tt.de)
		})
	}
}
