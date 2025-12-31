package config

import (
	"orca/model/policy"
	"testing"
)

func TestConfigCreator(t *testing.T) {
	dir := t.TempDir()
	fakeCreator := NewCreator(dir, policy.RealPolicy{})

	_, err := fakeCreator.Create("name")
	if err != nil {
		t.Error(err)
	}
}
