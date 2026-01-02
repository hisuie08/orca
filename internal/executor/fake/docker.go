package fakeexecutor

import (
	"orca/internal/executor"
)

var _ executor.Docker = (*FakeDocker)(nil)

type FakeDocker struct {
	Issued          []string
	Done            []string
	AllowSideEffect bool
}

// Record as abbreviated command
const (
	cmdUp        = "compose up"
	cmdDown      = "compose down"
	cmdCreateVol = "volume create"
	cmdCreateNet = "network create"
)

func (f *FakeDocker) ComposeUp(name string) (string, error) {
	cmd := cmdUp + ":" + name
	return f.runFake(cmd)
}
func (f *FakeDocker) ComposeDown(name string) (string, error) {
	cmd := cmdDown + ":" + name
	return f.runFake(cmd)
}
func (f *FakeDocker) CreateNetwork(name string, opts ...string) (string, error) {
	cmd := cmdCreateNet + ":" + name
	return f.runFake(cmd)
}
func (f *FakeDocker) CreateVolume(name string, opts ...string) (string, error) {
	cmd := cmdCreateVol + ":" + name
	return f.runFake(cmd)
}

func (f *FakeDocker) runFake(cmd string) (string, error) {
	f.Issued = append(f.Issued, cmd)
	if f.AllowSideEffect {
		f.Done = append(f.Done, cmd)
	}
	return cmd, nil
}
