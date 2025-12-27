package compose_test

import (
	"errors"
	"orca/infra/applier"
	"orca/internal/compose"
	"os"
	"path/filepath"

	"testing"
)

var _ compose.ComposeInspector = (*fakeInspector)(nil)

type fakeInspector struct {
	results map[string][]byte
	errors  map[string]error
}

func (f *fakeInspector) Config(dir string) ([]byte, error) {
	base := filepath.Base(dir)
	if err, ok := f.errors[base]; ok {
		return nil, err
	}
	return f.results[base], nil
}

func TestGetAllCompose_PartialFailure(t *testing.T) {
	root := t.TempDir()
	os.Mkdir(filepath.Join(root, "ok"), 0755)
	os.Mkdir(filepath.Join(root, "ng"), 0755)

	inspector := &fakeInspector{
		results: map[string][]byte{
			"ok": []byte("volumes: {}\nnetworks: {}"),
		},
		errors: map[string]error{
			"ng": errors.New("compose error"),
		},
	}

	got, err := compose.GetAllCompose(root, inspector)
	if err != nil {
		t.Fatal(err)
	}

	if len(*got) != 1 {
		t.Fatalf("expected 1 compose, got %d", len(*got))
	}

	if _, ok := (*got)["ok"]; !ok {
		t.Fatal("expected 'ok' to be present")
	}
}

var _ applier.ComposeWriter = (*fakeWriter)(nil)

type fakeWriter struct {
	dumped map[string][]byte
	errOn  string
}

func (f *fakeWriter) DumpCompose(name string, b []byte) (string, error) {
	if name == f.errOn {
		return "", errors.New("dump error")
	}
	if f.dumped == nil {
		f.dumped = map[string][]byte{}
	}
	f.dumped[name] = b
	return "compose." + name + ".yml", nil
}
func TestDumpAllComposes_OK(t *testing.T) {
	m := compose.ComposeMap{
		"a": &compose.ComposeSpec{Volumes: compose.VolumesSection{}},
		"b": &compose.ComposeSpec{Networks: compose.NetworksSection{}},
	}

	w := &fakeWriter{errOn: ""}

	got, err := m.DumpAllComposes(w)
	t.Log(got)
	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 2 {
		t.Fatalf("expected 2 results, got %d", len(got))
	}

	if _, ok := w.dumped["a"]; !ok {
		t.Fatal("compose a not dumped")
	}
}
