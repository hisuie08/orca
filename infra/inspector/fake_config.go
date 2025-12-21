package inspector

import "os"

// FakeConfigReader テスト用
type FakeConfigReader struct {
	Mock map[string]string
}

func (f FakeConfigReader) Read(fakeRoot string) ([]byte, error) {
	if b, ok := f.Mock[fakeRoot]; ok {
		return []byte(b), nil
	}
	return nil, os.ErrNotExist
}
