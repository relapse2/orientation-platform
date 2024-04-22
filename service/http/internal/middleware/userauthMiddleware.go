package middleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/utils"
	"orientation-platform/service/http/internal/config"
)

type UserAuthMiddleware struct {
	Config config.Config
}

func NewUserAuthMiddleware(c config.Config) *UserAuthMiddleware {
	return &UserAuthMiddleware{
		Config: c,
	}
}

func (m *UserAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var token string
		if token = r.URL.Query().Get("token"); token == "" {
			token = r.PostFormValue("token")
		}
		if token == "" {
			httpx.OkJson(w, apiErr.NotLogin)
			return
		}
		isTimeOut, err := utils.ValidToken(token, m.Config.UserAuth.AccessSecret)
		if err != nil || isTimeOut {
			httpx.OkJson(w, apiErr.InvalidToken)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
