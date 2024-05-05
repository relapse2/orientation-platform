package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/utils"
	"orientation-platform/service/http/internal/config"
)

type AdminAuthMiddleware struct {
	Config config.Config
}

func NewAdminAuthMiddleware(c config.Config) *AdminAuthMiddleware {
	return &AdminAuthMiddleware{
		Config: c,
	}
}

func (m *AdminAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var token string
		if token = r.URL.Query().Get("token"); token == "" {
			token = r.PostFormValue("token")
		}
		// 如果 token 仍然为空，尝试从 JSON 请求体中获取，这里复制一个出来，防止json解析时改变r
		if token == "" {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("readall error: %v", err)
				return
			}
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			var jsonData struct {
				Token string `json:"token"`
			}

			// 解析 JSON 请求体
			if err := json.NewDecoder(r.Body).Decode(&jsonData); err != nil {
				fmt.Printf("response body error: %v", err)
				return
			}
			// 获取 token
			token = jsonData.Token
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}
		if token == "" {
			httpx.Error(w, apiErr.NotLogin)
			return
		}
		isTimeOut, err := utils.ValidToken(token, m.Config.AdminAuth.AccessSecret)
		if err != nil || isTimeOut {
			httpx.Error(w, apiErr.InvalidToken)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
