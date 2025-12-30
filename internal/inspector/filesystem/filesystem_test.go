package filesystem

import "testing"

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
