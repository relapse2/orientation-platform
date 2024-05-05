package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orientation-platform/service/http/internal/logic/user"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
)

func AdminLoadStuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminLoadStuRequest

		if err := httpx.Parse(r, &req); err != nil {
			println(err)
			return
		}

		l := user.NewAdminLoadStuLogic(r.Context(), svcCtx)
		resp, err := l.AdminLoadStu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
