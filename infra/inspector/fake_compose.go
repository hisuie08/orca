package inspector

// FakeInpsector Compose取得テスト用
type FakeComposeInspector struct {
	Mock map[string]string
}

func (f FakeComposeInspector) Directories() ([]string, error) {
	keys := make([]string, 0, len(f.Mock))
	for k := range f.Mock {
		keys = append(keys, k)
	}
	keys = append(keys, "nocompose")
	return keys, nil
}

func (f FakeComposeInspector) Config(fakeComposeDir string) ([]byte, error) {
	v, ok := f.Mock[fakeComposeDir]
	if !ok {
		return nil, ErrNoCompose
	}
	return []byte(v), nil
}
