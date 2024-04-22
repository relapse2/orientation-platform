package logic

import (
	"context"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskInitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskInitLogic {
	return &TaskInitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskInitLogic) TaskInit(in *task.TaskInitRequest) (*task.TaskListReply, error) {
	// todo: add your logic here and delete this line

	return &task.TaskListReply{}, nil
}
