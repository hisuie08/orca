package filesystem

import (
	"testing"
)

var _ executor = (*fakeExecutor)(nil)

type fakeExecutor struct {
	Files           map[string][]byte
	Dirs            map[string]bool
	Issued          []string
	Done            []string
	AllowSideEffect bool
}

const (
	opWrite  = "WriteFile"
	opMkdir  = "CreateDir"
	opRmFile = "RemoveFile"
	opRmDir  = "RemoveDir"
)

func newFakeExecutor(allow bool) *fakeExecutor {
	return &fakeExecutor{
		Files:           map[string][]byte{},
		Dirs:            map[string]bool{},
		Issued:          []string{},
		Done:            []string{},
		AllowSideEffect: allow,
	}
}

func (f *fakeExecutor) WriteFile(path string, data []byte) error {
	op := opWrite + ":" + path
	f.Issued = append(f.Issued, op)
	if !f.AllowSideEffect {
		return nil
	}
	f.Done = append(f.Done, op)
	dir := dirOf(path)
	f.Dirs[dir] = true
	f.Files[path] = data
	return nil
}

func (f *fakeExecutor) CreateDir(path string) error {
	op := opMkdir + ":" + path
	f.Issued = append(f.Issued, op)
	if !f.AllowSideEffect {
		return nil
	}
	f.Done = append(f.Done, op)
	f.Dirs[path] = true
	return nil
}

func (f *fakeExecutor) RemoveFile(path string) error {
	op := opRmFile + ":" + path
	f.Issued = append(f.Issued, op)
	if !f.AllowSideEffect {
		return nil
	}
	f.Done = append(f.Done, op)
	delete(f.Files, path)
	return nil
}

func (f *fakeExecutor) RemoveDir(path string) error {
	op := opRmDir + ":" + path
	f.Issued = append(f.Issued, op)
	if !f.AllowSideEffect {
		return nil
	}
	f.Done = append(f.Done, op)
	delete(f.Dirs, path)
	return nil
}

// dirOf simulates filepath.Dir(path)
func dirOf(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[:i]
		}
	}
	return "."
}

func TestWriteFileCreatesParentDir(t *testing.T) {
	exec := newFakeExecutor(true)

	err := exec.WriteFile("a/b/c.txt", []byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	if !exec.Dirs["a/b"] {
		t.Fatalf("parent dir not created")
	}

	if string(exec.Files["a/b/c.txt"]) != "hello" {
		t.Fatalf("file content mismatch")
	}

	wantOps := []string{
		"WriteFile:a/b/c.txt",
	}

	if len(exec.Done) != len(wantOps) {
		t.Fatalf("ops mismatch: %+v", exec.Done)
	}
}

func TestDryRunDoesNotModifyState(t *testing.T) {
	exec := newFakeExecutor(false)

	_ = exec.WriteFile("x/y.txt", []byte("dry"))

	if len(exec.Files) != 0 {
		t.Fatalf("file should not be written in dry-run")
	}

	if len(exec.Issued) == 0 {
		t.Fatalf("op should be recorded")
	}
}
