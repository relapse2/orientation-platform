package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orientation-platform/service/http/internal/logic/user"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
)

func UserVisualHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserVisualRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserVisualLogic(r.Context(), svcCtx)
		resp, err := l.UserVisual(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
