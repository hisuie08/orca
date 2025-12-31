package config

import (
	"orca/model/config"
	"testing"
)

var _ loader = (*fakeLoader)(nil)

type fakeLoader struct {
}

func (f *fakeLoader) Load() (*config.ResolvedConfig, error) {
	return &config.ResolvedConfig{}, nil
}
func Test(t *testing.T) {
	fakeLoader := fakeLoader{}
	_, e := fakeLoader.Load()
	if e != nil {
		t.Fatal(e)
	}

}
