package ad

import (
	"context"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminSetAdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminSetAdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminSetAdLogic {
	return &AdminSetAdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminSetAdLogic) AdminSetAd(req *types.AdminSetAdRequest) (resp *types.AdminSetAdReply, err error) {
	// todo: add your logic here and delete this line

	return
}
