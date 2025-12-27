package fake

import (
	"orca/errs"
	"orca/infra/inspector"
)

var _ inspector.ComposeInspector = (*FakeComposeInspector)(nil)

type FakeComposeInspector struct {
	composes map[string][]byte
	root     string
}

func (f FakeComposeInspector) Root() string {
	return f.root
}

func (f FakeComposeInspector) Config(composeDir string) ([]byte, error) {
	if res, ok := f.composes[composeDir]; ok {
		return res, nil
	}
	return []byte{}, errs.ErrComposeNotFound
}
