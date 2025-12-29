package errs

import (
	"fmt"
	"orca/internal/errdef"
)

var (
	ErrAlreadyInitialized = errdef.ErrAlreadyInitialized
	ErrNotInitialized     = errdef.ErrNotInitialized

	ErrPlanNotBuilt = errdef.ErrPlanNotBuilt
	ErrPlanDirty    = errdef.ErrPlanDirty

	ErrDryRunViolation = errdef.ErrDryRunViolation

	ErrComposeNotFound = errdef.ErrComposeNotFound
	ErrClusterNotFound = errdef.ErrClusterNotFound

	ErrExternalDependency = errdef.ErrExternalDependency
)

type ExternalError struct {
	Cmd string
	Err error
}

func (e *ExternalError) Error() string {
	return fmt.Sprintf("%s failed: %v", e.Cmd, e.Err)
}

func (e *ExternalError) Unwrap() error {
	return errdef.ErrExternalDependency
}

type FileError struct {
	Path string
	Err  error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("file operation failed: %s: %v", e.Path, e.Err)
}

func (e *FileError) Unwrap() error {
	return errdef.ErrFileOperation
}
