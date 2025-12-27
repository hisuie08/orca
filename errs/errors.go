package errs

import "orca/internal/errdef"

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
