package config_test

import (
	"orca/internal/loader/config"
	"os"
	"testing"
)

func Test(t *testing.T) {
	dir := t.TempDir()
	testpath := os.DirFS("/workspace/orca/testdata/config")
	os.CopyFS(dir, testpath)
	fakeLoader := config.NewLoader(dir)
	l, e := fakeLoader.Load()
	if e != nil {
		t.Fatal(e)
	}
	t.Logf("%#v", l)

}
