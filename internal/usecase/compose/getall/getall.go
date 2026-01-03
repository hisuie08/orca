package getall

import (
	"errors"
	"orca/errs"
	"orca/internal/context"
	"orca/internal/inspector"
	"orca/model/compose"
	"path/filepath"
)

type GetAllComposeContext interface {
	context.WithRoot
}

func GetAllCompose(ctx GetAllComposeContext) (*compose.ComposeMap, error) {
	return getAllCompose(ctx, inspector.NewDocker(), inspector.NewFilesystem())
}

type dockerInspector interface {
	Compose(string) (*compose.ComposeSpec, error)
}
type fsInspector interface {
	Dirs(string) ([]string, error)
}

func getAllCompose(ctx GetAllComposeContext,
	di dockerInspector, fi fsInspector) (
	*compose.ComposeMap, error) {
	result := compose.ComposeMap{}
	dirs, err := fi.Dirs(ctx.Root())
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		data, err := di.Compose(dir)
		if err != nil {
			if errors.Is(err, errs.ErrComposeNotFound) {
				continue
			}
			return nil, err
		}

		result[filepath.Base(dir)] = data
	}
	return &result, nil
}
