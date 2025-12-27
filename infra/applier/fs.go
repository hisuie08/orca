package applier

import "orca/ostools"

type DirCreator interface {
	CreateDir(path string) error
}

type fsApplier struct{}

func (fsApplier) CreateDir(path string) error {
	return ostools.CreateDir(path)
}
