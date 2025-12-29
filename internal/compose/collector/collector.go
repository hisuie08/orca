package collector

import (
	"errors"
	"orca/errs"
	"orca/internal/context"
	inspector "orca/internal/inspector/compose"
	model "orca/model/compose"
	"os"
	"path/filepath"
)

var _ ComposeCollector = (*composeCollector)(nil)

type ComposeCollector interface {
	GetAll() (*model.ComposeMap, error)
}

func NewCollector(root string, ins inspector.ComposeInspector) ComposeCollector {
	return &composeCollector{WithRoot: context.NewWithRoot(root), ins: ins}
}

type composeCollector struct {
	context.WithRoot
	ins inspector.ComposeInspector
}

func (c *composeCollector) GetAll() (*model.ComposeMap, error) {
	result := model.ComposeMap{}
	dirs, err := c.dirs()
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		data, err := c.ins.Config(dir)
		if err != nil && errors.Is(err, errs.ErrComposeNotFound) {
			continue
		}

		result[filepath.Base(dir)] = data
	}
	return &result, nil

}

func (c *composeCollector) dirs() ([]string, error) {
	result := []string{}
	readDir, err := os.ReadDir(c.Root())
	if err != nil {
		return nil, err
	}
	for _, d := range readDir {
		if d.IsDir() {
			result = append(result, filepath.Join(c.Root(), d.Name()))
		}
	}
	return result, nil
}
