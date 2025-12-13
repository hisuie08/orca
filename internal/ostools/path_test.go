package ostools

import (
	"testing"
)

var testpath = "/workspace/orca/testdata/pathtest"

func TestPathExists(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path      string
		expectDir bool
		want      bool
	}{
		// TODO: Add test cases.
		{"existsDir", testpath + "/dir", true, true},
		{"existsFile", testpath + "/file", false, true},
		{"ExistsButNotDir", testpath + "/file", true, false},
		{"ExistsButNotFile", testpath + "/dir", false, false},
		{"NotExistsFile", testpath + "/nodir", true, false},
		{"NotExistsFile", testpath + "/nofile", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			println(tt.path)
			got := pathExists(tt.path, tt.expectDir)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("PathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirExists(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path string
		want bool
	}{
		// TODO: Add test cases.

		// TODO: Add test cases.
		{"existsDir", testpath + "/dir", true},
		{"existsFile", testpath + "/file", false},
		{"ExistsButNotDir", testpath + "/file", false},
		{"ExistsButNotFile", testpath + "/dir", true},
		{"NotExistsFile", testpath + "/nodir", false},
		{"NotExistsFile", testpath + "/nofile", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DirExists(tt.path)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("DirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileExisists(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path string
		want bool
	}{

		// TODO: Add test cases.
		{"existsDir", testpath + "/dir", false},
		{"existsFile", testpath + "/file", true},
		{"ExistsButNotDir", testpath + "/file", true},
		{"ExistsButNotFile", testpath + "/dir", false},
		{"NotExistsFile", testpath + "/nodir", false},
		{"NotExistsFile", testpath + "/nofile", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FileExisists(tt.path)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("FileExisists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectories(t *testing.T) {
	testpath = "/workspace/orca/testdata"
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", testpath, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Directories(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Directories() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Directories() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != nil {
				t.Logf("Directories() = %v\n", got)
			}
		})
	}
}

func TestFiles(t *testing.T) {
	testpath = "/workspace/orca/testdata"
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", testpath, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := Files(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Files() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Files() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if got != nil {
				t.Logf("Files() = %v\n", got)
			}
		})
	}
}

func TestAppendFile(t *testing.T) {
	testpath = "/workspace/orca/testdata/append/append.txt"
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		target  string
		content string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", testpath, "log1", false},
		{"test2", testpath, `multiple
		line
		end
		`, false},
		{"test3", testpath, "log3\n", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := AppendFile(tt.target, tt.content)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("AppendFile() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("AppendFile() succeeded unexpectedly")
			}
		})
	}
}

func TestCreateFile(t *testing.T) {
	testpath = "/workspace/orca/testdata/create/create.txt"
	test1, test2, test3 := "log1", `multiple
	line
	`, ""
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		target  string
		content []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test1", testpath, []byte(test1), false},
		{"test2", testpath, []byte(test2), false},
		{"test3", testpath, []byte(test3), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := CreateFile(tt.target, tt.content)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateFile() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateFile() succeeded unexpectedly")
			}
		})
	}
}
