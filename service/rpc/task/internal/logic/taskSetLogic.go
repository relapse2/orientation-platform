package logic

import (
	"context"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskSetLogic {
	return &TaskSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskSetLogic) TaskSet(in *task.TaskSetRequest) (*task.Empty, error) {
	// todo: add your logic here and delete this line

	return &task.Empty{}, nil
}
