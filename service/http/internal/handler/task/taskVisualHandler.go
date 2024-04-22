package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orientation-platform/service/http/internal/logic/task"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
)

func TaskVisualHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskVisualRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := task.NewTaskVisualLogic(r.Context(), svcCtx)
		resp, err := l.TaskVisual(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
