package task

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskListLogic) TaskList(req *types.TaskListRequest) (resp *types.TaskListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
