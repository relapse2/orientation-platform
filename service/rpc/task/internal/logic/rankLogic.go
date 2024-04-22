package logic

import (
	"context"

	"orientation-platform/service/rpc/task/internal/svc"
	"orientation-platform/service/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type RankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RankLogic {
	return &RankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RankLogic) Rank(in *task.RankRequest) (*task.RankReply, error) {
	// todo: add your logic here and delete this line

	return &task.RankReply{}, nil
}
