package creator_test

import (
	"orca/internal/config/creator"
	"orca/internal/policy"
	"testing"
)

func TestConfigCreator(t *testing.T) {
	dir := t.TempDir()
	fakeCreator := creator.NewCreator(dir, policy.RealPolicy{})

	_, err := fakeCreator.Create("name")
	if err != nil {
		t.Error(err)
	}
}
