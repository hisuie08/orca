package filesystem

import (
	"io/fs"
	"orca/errs"
	"testing"
)

var _ Inspector = (*fakeInspector)(nil)

type fakeInspector struct {
	FileMap map[string][]byte
	DirMap  map[string]bool
}

func newFakeInspector() Inspector {
	return &fakeInspector{
		FileMap: map[string][]byte{},
		DirMap:  map[string]bool{},
	}
}

func (f *fakeInspector) FileExists(path string) bool {
	_, ok := f.FileMap[path]
	return ok
}

func (f *fakeInspector) DirExists(path string) bool {
	return f.DirMap[path]
}

func (f *fakeInspector) Dirs(path string) ([]string, error) {
	out := []string{}
	for d := range f.DirMap {
		if parentOf(d) == path {
			out = append(out, d)
		}
	}
	return out, nil
}

func (f *fakeInspector) Files(path string) ([]string, error) {
	out := []string{}
	for file := range f.FileMap {
		if parentOf(file) == path {
			out = append(out, file)
		}
	}
	return out, nil
}

func (f *fakeInspector) Read(path string) ([]byte, error) {
	if r, ok := f.FileMap[path]; ok {
		return r, &errs.FileError{Path: path, Err: fs.ErrNotExist}
	}
	return f.FileMap[path], nil
}

func parentOf(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[:i]
		}
	}
	return "."
}

func TestInspectorListsFiles(t *testing.T) {
	ins := newFakeInspector().(*fakeInspector)

	ins.DirMap["a"] = true
	ins.FileMap["a/x.txt"] = []byte("x")
	ins.FileMap["a/y.txt"] = []byte("y")

	files, _ := ins.Files("a")

	if len(files) != 2 {
		t.Fatalf("expected 2 files")
	}
}
