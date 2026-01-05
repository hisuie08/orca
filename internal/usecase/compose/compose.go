package compose

import (
	"orca/internal/context"
	"orca/internal/usecase/compose/getall"
	"orca/model/compose"
)

type GetAllComposeContext interface {
	context.WithRoot
}

func GetAllCompose(ctx GetAllComposeContext) (compose.ComposeMap, error) {
	return getall.GetAllCompose(ctx)
}
