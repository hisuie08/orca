package orca

import (
	"fmt"
)

func OrcaError(op string, err error) error {
	return fmt.Errorf("%s: %w", op, err)
}

