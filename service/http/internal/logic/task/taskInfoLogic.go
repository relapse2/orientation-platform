package task

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskInfoLogic {
	return &TaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskInfoLogic) TaskInfo(req *types.TaskInfoRequest) (resp *types.TaskInfoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
