package docker

import "strings"

var _ Executor = (*fakeExecutor)(nil)

type fakeExecutor struct {
	Ops             []string
	AllowSideEffect bool
}

func newFakeExecutor(allow bool) Executor {
	return &fakeExecutor{Ops: []string{}, AllowSideEffect: allow}
}

func (f *fakeExecutor) ComposeUp(name string) (string, error) {
	cmd := "docker compose -f " + name + " up -d"
	return f.runFake(cmd)
}
func (f *fakeExecutor) ComposeDown(name string) (string, error) {
	cmd := "docker compose -f " + name + " down"
	return f.runFake(cmd)
}
func (f *fakeExecutor) CreateNetwork(name string, opts ...string) (string, error) {
	cmd := strings.Join(append([]string{"docker", "network", "create", name}, opts...), " ")
	return f.runFake(cmd)
}
func (f *fakeExecutor) CreateVolume(name string, opts ...string) (string, error) {
	cmd := strings.Join(append([]string{"docker", "volume", "create", name}, opts...), " ")
	return f.runFake(cmd)
}

func (f *fakeExecutor) runFake(cmd string) (string, error) {
	if f.AllowSideEffect {
		f.Ops = append(f.Ops, cmd)
	}
	return cmd, nil
}
