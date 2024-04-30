package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"orientation-platform/common/oss"
)

type Config struct {
	// 启动配置
	rest.RestConf

	// RPC 配置
	TaskRpc zrpc.RpcClientConf
	UserRpc zrpc.RpcClientConf
	AdRpc   zrpc.RpcClientConf

	// JWT 配置
	UserAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	AdminAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	OSS oss.AliyunCfg
}
