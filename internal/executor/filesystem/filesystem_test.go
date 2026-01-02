package filesystem

import (
	"testing"
)

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

	if len(exec.Issued) ==0 {
		t.Fatalf("op should be recorded")
	}
}
