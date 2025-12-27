package errdef

import (
	"errors"
)

var (
	ErrAlreadyInitialized = errors.New("orca already initialized")
	ErrNotInitialized     = errors.New("orca not initialized")

	ErrPlanNotBuilt = errors.New("plan not built")
	ErrPlanDirty    = errors.New("plan is out of date")

	ErrDryRunViolation = errors.New("operation violates dry-run")

	
	ErrComposeNotFound = errors.New("compose file not found")
	ErrClusterNotFound = errors.New("cluster not found")

	ErrExternalDependency = errors.New("external dependency failed")
)
