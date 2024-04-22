package logic

import (
	"context"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type FailTaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFailTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FailTaskListLogic {
	return &FailTaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FailTaskListLogic) FailTaskList(in *task.Empty) (*task.FailTaskListReply, error) {
	// todo: add your logic here and delete this line

	return &task.FailTaskListReply{}, nil
}
