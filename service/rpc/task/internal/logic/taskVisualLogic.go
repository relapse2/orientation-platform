package logic

import (
	"context"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskVisualLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskVisualLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskVisualLogic {
	return &TaskVisualLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskVisualLogic) TaskVisual(in *task.Empty) (*task.TaskVisualReply, error) {
	// todo: add your logic here and delete this line

	return &task.TaskVisualReply{}, nil
}
