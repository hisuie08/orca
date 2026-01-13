package docker

import "strings"

var _ executor = (*fakeExecutor)(nil)

type fakeExecutor struct {
	Issued          []string
	Done            []string
	AllowSideEffect bool
}

func newFakeExecutor(allow bool) executor {
	return &fakeExecutor{Issued: []string{}, Done: []string{}, AllowSideEffect: allow}
}

func (f *fakeExecutor) ComposeUp(name string) ([]byte, error) {
	cmd := "docker compose -f " + name + " up -d"
	return f.runFake(cmd)
}
func (f *fakeExecutor) ComposeDown(name string) ([]byte, error) {
	cmd := "docker compose -f " + name + " down"
	return f.runFake(cmd)
}
func (f *fakeExecutor) CreateNetwork(name string, opts ...string) ([]byte, error) {
	cmd := strings.Join(append([]string{"docker", "network", "create", name}, opts...), " ")
	return f.runFake(cmd)
}
func (f *fakeExecutor) CreateVolume(name string, opts ...string) ([]byte, error) {
	cmd := strings.Join(append([]string{"docker", "volume", "create", name}, opts...), " ")
	return f.runFake(cmd)
}

func (f *fakeExecutor) runFake(cmd string) ([]byte, error) {
	f.Issued = append(f.Issued, cmd)
	if f.AllowSideEffect {
		f.Done = append(f.Done, cmd)
	}
	return []byte(cmd), nil
}
