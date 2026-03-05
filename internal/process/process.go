package process

import (
	"context"
)

type Process interface {
	Run(ctx context.Context, o ...any)
}
