package fakeinspector

type FakeFilesystem struct {
	FileMap map[string][]byte
	DirMap  map[string][]string
	Reads   []string
}

func (f *FakeFilesystem) FileExists(path string) bool {
	_, ok := f.FileMap[path]
	return ok
}

func (f *FakeFilesystem) DirExists(path string) bool {
	_, ok := f.DirMap[path]
	return ok
}

func (f *FakeFilesystem) Dirs(path string) ([]string, error) {
	return f.DirMap[path], nil
}

func (f *FakeFilesystem) Files(path string) ([]string, error) {
	return f.DirMap[path], nil
}

func (f *FakeFilesystem) Read(path string) ([]byte, error) {
	return f.FileMap[path], nil
}
