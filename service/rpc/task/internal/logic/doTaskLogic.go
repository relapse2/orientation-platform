package logic

import (
	"context"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDoTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoTaskLogic {
	return &DoTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DoTaskLogic) DoTask(in *task.DoTaskRequest) (*task.DoTaskReply, error) {
	// todo: add your logic here and delete this line

	return &task.DoTaskReply{}, nil
}
