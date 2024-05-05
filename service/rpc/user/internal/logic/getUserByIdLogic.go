package logic

import (
	"context"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"orientation-platform/common/error/rpcErr"
	"orientation-platform/common/model"
	"orientation-platform/common/utils"

	"orientation-platform/service/rpc/user/internal/svc"
	"orientation-platform/service/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserByIdRequest) (*user.GetUserReply, error) {
	// 查询缓存
	cacheData, err := l.svcCtx.DBList.Redis.HGetAll(l.ctx, utils.GenUserInfoCacheKey(in.Id)).Result()
	if err == nil && len(cacheData) != 0 {
		l.Logger.Info("Get user info from cache")
		return &user.GetUserReply{
			Id:           utils.Str2Int64(cacheData["Id"]),
			Name:         cacheData["Name"],
			Password:     cacheData["Password"],
			Point:        utils.Str2Int64(cacheData["Point"]),
			AvatarUrl:    cacheData["AvatarUrl"],
			CharacterUrl: cacheData["CharacterUrl"],
			Sex:          utils.Str2Int64(cacheData["Sex"]),
			Collage:      cacheData["Collage"],
		}, nil
	} else if err != nil && err != redis.Nil {
		l.Logger.Error(rpcErr.CacheError.Code, err.Error())
	}

	// 准备模型
	result := &model.User{}

	// 查询数据
	err = l.svcCtx.DBList.Mysql.Where("id = ?", in.Id).First(result).Error

	if err == gorm.ErrRecordNotFound {
		return nil, status.Error(rpcErr.UserNotExist.Code, rpcErr.UserNotExist.Message)
	} else if err != nil {
		return nil, status.Error(rpcErr.DataBaseError.Code, err.Error())
	}

	// 如果是排行榜上，此处用Point来判断可能的热门程度，因为point越高，越可能上排行榜，缓存个人信息
	//此处主要是想用一下缓存
	if model.IsPopularUser(result.Point) {
		err := l.svcCtx.DBList.Redis.HSet(l.ctx, utils.GenUserInfoCacheKey(int64(result.ID)), map[string]interface{}{
			"Id":           result.ID,
			"Name":         result.Username,
			"Password":     result.Password,
			"Point":        result.Point,
			"AvatarUrl":    result.AvatarUrl,
			"CharacterUrl": result.CharacterUrl,
			"Sex":          result.Sex,
			"Collage":      result.Collage,
		}).Err()
		if err != nil {
			return nil, status.Error(rpcErr.CacheError.Code, err.Error())
		}

		err = l.svcCtx.DBList.Redis.LPush(l.ctx, utils.GenPopUserListCacheKey(), result.ID).Err()
		if err != nil {
			return nil, status.Error(rpcErr.CacheError.Code, err.Error())
		}
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
