package docker

import (
	"fmt"
	"strings"
	"testing"
)

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

func TestDockerExecutor(t *testing.T) {
	testCases := []struct {
		desc       string
		allow      bool
		wantIssued int
		wantDone   int
	}{
		{desc: "real", allow: true, wantIssued: 4, wantDone: 4},
		{desc: "dry", allow: false, wantIssued: 4, wantDone: 0},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%s_allow=%v", tC.desc, tC.allow), func(t *testing.T) {
			fake := newFakeExecutor(tC.allow).(*fakeExecutor)
			if _, e := fake.ComposeUp(""); e != nil {
				t.Fatal(e)
			}
			if _, e := fake.ComposeDown(""); e != nil {
				t.Fatal(e)
			}
			if _, e := fake.CreateVolume(""); e != nil {
				t.Fatal(e)
			}
			if _, e := fake.CreateNetwork(""); e != nil {
				t.Fatal(e)
			}
			if len(fake.Issued) != tC.wantIssued {
				t.Errorf("expected %d in Issued but got %d", len(fake.Issued), tC.wantIssued)
			}
			if len(fake.Done) != tC.wantDone {
				t.Errorf("expected %d in Done but got %d", len(fake.Done), tC.wantDone)
			}
		})
	}
}
