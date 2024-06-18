package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/common/model"

	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserIdByCollageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserIdByCollageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserIdByCollageLogic {
	return &GetUserIdByCollageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserIdByCollageLogic) GetUserIdByCollage(in *user.GetUserIdByCollageRequest) (*user.GetUserIdByCollageReply, error) {
	// 获取对应大学的user列表
	var users []model.User

	if err := l.svcCtx.DBList.Mysql.
		Where("collage = ?", in.Collage).
		Find(&users).Error; err != nil {
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	IdList := make([]int64, 0, len(users))
	for _, user := range users {
		IdList = append(IdList, int64(user.ID))
	}

	return &user.GetUserIdByCollageReply{
		IdList: IdList,
	}, nil
}
