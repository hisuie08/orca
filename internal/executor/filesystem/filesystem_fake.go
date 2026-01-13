package filesystem

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
