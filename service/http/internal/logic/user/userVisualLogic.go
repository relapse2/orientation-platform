package user

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserVisualLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserVisualLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserVisualLogic {
	return &UserVisualLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserVisualLogic) UserVisual(req *types.UserVisualRequest) (resp *types.UserVisualReply, err error) {
	// todo: add your logic here and delete this line

	return
}
