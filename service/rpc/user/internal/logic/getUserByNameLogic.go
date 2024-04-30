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

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *user.GetUserByNameRequest) (*user.GetUserReply, error) {
	// 准备数据
	result := &model.User{}

	// 查询数据
	err := l.svcCtx.DBList.Mysql.Where("username = ?", in.Name).First(result).Error

	if err == gorm.ErrRecordNotFound {
		return nil, status.Error(rpcErr.UserNotExist.Code, rpcErr.UserNotExist.Message)
	} else if err != nil {
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	return &user.GetUserReply{
		Id:           int64(result.ID),
		Name:         result.Username,
		Password:     result.Password,
		Point:        result.Point,
		AvatarUrl:    result.AvatarUrl,
		CharacterUrl: result.CharacterUrl,
		Sex:          result.Sex,
		Collage:      result.Collage,
	}, nil
}
