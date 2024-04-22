package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orientation-platform/service/http/internal/logic/user"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
)

func UpgradeCharHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeCharRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpgradeCharLogic(r.Context(), svcCtx)
		resp, err := l.UpgradeChar(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
