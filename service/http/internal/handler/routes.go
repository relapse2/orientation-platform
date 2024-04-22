// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	ad "orientation-platform/service/http/internal/handler/ad"
	task "orientation-platform/service/http/internal/handler/task"
	user "orientation-platform/service/http/internal/handler/user"
	"orientation-platform/service/http/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/platform/ad/adminset",
					Handler: ad.AdminSetAdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/platform/ad/checkcost",
					Handler: ad.CheckCostHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/platform/ad/adlist",
					Handler: ad.AdListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/platform/ad/adinfo",
					Handler: ad.AdInfoHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/platform/ad/sendad",
					Handler: ad.SendAdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/platform/ad/clickad",
					Handler: ad.ClickAdHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/platform/task/list",
					Handler: task.TaskListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/platform/task/init",
					Handler: task.TaskInitHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/douyin/task/do",
					Handler: task.DoTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/platform/task",
					Handler: task.TaskInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/douyin/task/rank",
					Handler: task.RankHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/platform/task/admininfo",
					Handler: task.AdminTaskInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/platform/task/faillist",
					Handler: task.FailTaskListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/douyin/task/check",
					Handler: task.AdminCheckTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/douyin/task/visual",
					Handler: task.TaskVisualHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/platform/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/platform/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/platform/user/adminlogin",
				Handler: user.AdminLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/platform/user",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/platform/user/upgradechar",
					Handler: user.UpgradeCharHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/platform/user/load",
					Handler: user.AdminLoadStuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/platform/user/visual",
					Handler: user.UserVisualHandler(serverCtx),
				},
			}...,
		),
	)
}
