package logic

import (
	"context"

	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserVisualLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserVisualLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserVisualLogic {
	return &UserVisualLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserVisualLogic) UserVisual(in *user.Empty) (*user.UserVisualReply, error) {
	// todo: add your logic here and delete this line

	return &user.UserVisualReply{}, nil
}
