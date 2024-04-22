package logic

import (
	"context"

	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RankByPointsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRankByPointsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RankByPointsLogic {
	return &RankByPointsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RankByPointsLogic) RankByPoints(in *user.Empty) (*user.RankByPointsReply, error) {
	// todo: add your logic here and delete this line

	return &user.RankByPointsReply{}, nil
}
