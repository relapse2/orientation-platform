package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	oss2 "orientation-platform/common/oss"
	"orientation-platform/service/http/internal/config"
	"orientation-platform/service/http/internal/middleware"
	"orientation-platform/service/rpc/ad/adclient"
	"orientation-platform/service/rpc/task/taskclient"
	"orientation-platform/service/rpc/user/userclient"
)

type ServiceContext struct {
	Config       config.Config
	AdminAuth    rest.Middleware
	UserAuth     rest.Middleware
	AliyunClient *oss.Client
	TaskRpc      taskclient.Task
	UserRpc      userclient.User
	AdRpc        adclient.Ad
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		AdminAuth:    middleware.NewAdminAuthMiddleware(c).Handle,
		UserAuth:     middleware.NewUserAuthMiddleware(c).Handle,
		AliyunClient: oss2.AliyunInit(c.OSS),
		TaskRpc:      taskclient.NewTask(zrpc.MustNewClient(c.TaskRpc)),
		UserRpc:      userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AdRpc:        adclient.NewAd(zrpc.MustNewClient(c.AdRpc)),
	}
}
