package logic

import (
	"context"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskInfoLogic {
	return &TaskInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskInfoLogic) TaskInfo(in *task.TaskInfoRequest) (*task.TaskInfoReply, error) {
	// todo: add your logic here and delete this line

	return &task.TaskInfoReply{}, nil
}
