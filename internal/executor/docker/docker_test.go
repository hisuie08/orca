package docker_test

import (
	"orca/internal/executor/docker"
	"orca/internal/policy"
	"testing"
)

func Test(t *testing.T) {
	fake := docker.NewExecutor(policy.DryPolicy{})
	o, e := fake.ComposeUp("compose.test.yml")
	if e != nil {
		t.Fatal(e)
	}
	t.Log(o)
}
