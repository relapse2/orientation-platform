package task

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RankLogic {
	return &RankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RankLogic) Rank(req *types.RankRequest) (resp *types.RankReply, err error) {
	// todo: add your logic here and delete this line

	return
}
