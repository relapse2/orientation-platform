package ad

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orientation-platform/service/http/internal/logic/ad"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
)

func AdminSetAdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminSetAdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ad.NewAdminSetAdLogic(r.Context(), svcCtx)
		resp, err := l.AdminSetAd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
