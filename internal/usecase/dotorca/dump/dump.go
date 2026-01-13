package dump

import (
	"fmt"
	"io/fs"
	"orca/errs"
	"orca/internal/context"
	"orca/internal/executor"
	"orca/internal/inspector"
	"orca/model/compose"
	"orca/model/plan"
	"path/filepath"
	"sort"

	"gopkg.in/yaml.v3"
)

var _ Dumper = (*dumper)(nil)

type Dumper interface {
	DumpComposes(compose.ComposeMap) ([]string, error)
	DumpPlan(plan.OrcaPlan) (string, error)
}

type dumpContext interface {
	context.WithRoot
	context.WithPolicy
	context.WithReport
}

type dumper struct {
	ctx   dumpContext
	force bool
	fi    inspector.FileSystem
	fe    executor.FileSystem
}

func DotOrcaDumper(ctx dumpContext, force bool) *dumper {
	return &dumper{ctx: ctx,
		force: force,
		fi:    inspector.NewFilesystem(),
		fe:    executor.NewFilesystem(ctx),
	}
}
func (d *dumper) DumpComposes(cm compose.ComposeMap) ([]string, error) {
	written := []string{}
	names := make([]string, 0, len(cm))
	for name := range cm {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		spec := cm[name]
		filename := fmt.Sprintf("compose.%s.yml", name)
		path := filepath.Join(d.ctx.OrcaDir(), filename)
		data, err := yaml.Marshal(spec)
		if err != nil {
			return written, err
		}
		if err := d.dumpDotOrca(path, data); err != nil {
			return written, err
		}
		written = append(written, path)
	}
	return written, nil
}

func (d *dumper) DumpPlan(pl plan.OrcaPlan) (string, error) {
	path := filepath.Join(d.ctx.OrcaDir(), "plan.yml")
	data, err := yaml.Marshal(pl)
	if err != nil {
		return "", err
	}
	if err := d.dumpDotOrca(path, data); err != nil {
		return "", &errs.FileError{Path: path, Err: err}
	}
	return path, nil
}

func (d *dumper) dumpDotOrca(path string, data []byte) error {
	if d.fi.FileExists(path) && !d.force {
		return &errs.FileError{Path: path, Err: fs.ErrExist}
	}
	return d.fe.WriteFile(path, data)
}
