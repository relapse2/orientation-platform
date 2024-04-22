package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orientation-platform/service/http/internal/logic/task"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
)

func AdminCheckTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminCheckTaskRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := task.NewAdminCheckTaskLogic(r.Context(), svcCtx)
		resp, err := l.AdminCheckTask(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
