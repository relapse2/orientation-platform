package task

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskVisualLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskVisualLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskVisualLogic {
	return &TaskVisualLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskVisualLogic) TaskVisual(req *types.TaskVisualRequest) (resp *types.TaskVisualReply, err error) {
	// todo: add your logic here and delete this line

	return
}
