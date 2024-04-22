package task

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskInitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskInitLogic {
	return &TaskInitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskInitLogic) TaskInit(req *types.TaskInitRequest) (resp *types.TaskInitReply, err error) {
	// todo: add your logic here and delete this line

	return
}
