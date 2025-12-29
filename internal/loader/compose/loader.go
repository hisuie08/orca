package compose

import (
	"errors"
	"orca/errs"
	ins "orca/infra/inspector/compose"
	"orca/infra/inspector/fs"
	"orca/internal/context"
	"orca/model/compose"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var _ compose.ComposeLoader = (*composeLoader)(nil)

type composeLoader struct {
	context.WithRoot
	compose ins.ComposeInspector
	fs      fs.FsInspector
}

func NewLoader(root string) *composeLoader {
	return &composeLoader{
		WithRoot: context.NewWithRoot(root),
		fs:       fs.NewInspector,
		compose:  ins.NewInspector(root),
	}
}

func (c *composeLoader) Load() (compose.ComposeMap, error) {
	result := compose.ComposeMap{}
	dirs, err := c.fs.Dirs(c.Root())
	if err != nil {
		return result, &errs.FileError{Path: c.Root(), Err: err}
	}
	for _, dir := range dirs {
		data, err := c.compose.Config(dir)
		if err != nil && errors.Is(err, errs.ErrComposeNotFound) {
			continue
		}

		spec := &compose.ComposeSpec{}
		if err := yaml.Unmarshal(data, spec); err != nil {
			return nil, err
		}

		result[filepath.Base(dir)] = spec
	}
	return result, nil
}
