package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/common/model"

	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminByNameLogic {
	return &GetAdminByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminByNameLogic) GetAdminByName(in *user.GetAdminByNameRequest) (*user.GetAdminReply, error) {
	// 准备数据
	result := &model.Admin{}

	// 查询数据
	err := l.svcCtx.DBList.Mysql.Where("username = ?", in.Name).First(result).Error

	if err == gorm.ErrRecordNotFound {
		return nil, status.Error(rpcErr.UserNotExist.Code, rpcErr.UserNotExist.Message)
	} else if err != nil {
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	return &user.GetAdminReply{
		Id:       int64(result.ID),
		Name:     result.Username,
		Password: result.Password,
	}, nil
}
