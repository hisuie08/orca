package filesystem

type fakeExecutor struct {
	Files map[string][]byte
	Dirs  map[string]bool
	Ops   []string

	AllowSideEffect bool
}

func newFakeExecutor(allow bool) Executor {
	return &fakeExecutor{
		Files:           map[string][]byte{},
		Dirs:            map[string]bool{},
		Ops:             []string{},
		AllowSideEffect: allow,
	}
}

func (f *fakeExecutor) WriteFile(path string, data []byte) error {
	f.Ops = append(f.Ops, "WriteFile:"+path)

	if !f.AllowSideEffect {
		return nil
	}

	dir := dirOf(path)
	f.Dirs[dir] = true
	f.Files[path] = data
	return nil
}

func (f *fakeExecutor) CreateDir(path string) error {
	f.Ops = append(f.Ops, "CreateDir:"+path)

	if !f.AllowSideEffect {
		return nil
	}

	f.Dirs[path] = true
	return nil
}

func (f *fakeExecutor) RemoveFile(path string) error {
	f.Ops = append(f.Ops, "RemoveFile:"+path)

	if !f.AllowSideEffect {
		return nil
	}

	delete(f.Files, path)
	return nil
}

func (f *fakeExecutor) RemoveDir(path string) error {
	f.Ops = append(f.Ops, "RemoveDir:"+path)

	if !f.AllowSideEffect {
		return nil
	}

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
