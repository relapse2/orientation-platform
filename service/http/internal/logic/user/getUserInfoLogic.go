package user

import (
	"context"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/service/rpc/user/types/user"

	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoReply, err error) {
	logx.WithContext(l.ctx).Infof("获取用户信息: %v", req)

	//获取用户信息(名字与id)
	getUserByIdReply, err := l.svcCtx.UserRpc.GetUserById(l.ctx, &user.GetUserByIdRequest{
		Id: req.UserId,
	})
	if rpcErr.Is(err, rpcErr.UserNotExist) {
		return nil, apiErr.UserNotFound
	} else if err != nil {
		logx.WithContext(l.ctx).Errorf("获取用户信息失败: %v", err)
		return nil, apiErr.InternalError(l.ctx, err.Error())
	}

	return &types.GetUserInfoReply{
		BasicReply: types.BasicReply(apiErr.Success),
		User: types.User{
			Id:           getUserByIdReply.Id,
			Name:         getUserByIdReply.Name,
			Point:        getUserByIdReply.Point,
			AvatarUrl:    getUserByIdReply.AvatarUrl,
			CharacterUrl: getUserByIdReply.CharacterUrl,
			Sex:          getUserByIdReply.Sex,
			Collage:      getUserByIdReply.Collage,
		},
	}, nil
}
