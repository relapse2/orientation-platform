package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/h2non/filetype"
	"github.com/thinkeridea/go-extend/exnet"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/oss"
	"orientation-platform/common/utils"
	logic "orientation-platform/service/http/internal/logic/user"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
	"orientation-platform/service/rpc/user/types/user"
	"path/filepath"
	"unicode"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		//参数检查
		// 检查身份证号的长度是否为 18
		if len(req.IdCard) != 18 {
			httpx.Error(w, apiErr.InvalidParams.WithDetails("invalid ID card length"))
			return
		}

		// 检查身份证号第 17 位是否为数字
		// 使用 unicode.IsDigit() 来验证
		if !unicode.IsDigit(rune(req.IdCard[16])) {
			httpx.Error(w, apiErr.InvalidParams.WithDetails("invalid character at position 17"))
			return
		}
		//查找ip
		ip := exnet.ClientPublicIP(r)
		if ip == "" {
			ip = exnet.ClientIP(r)
		}

		//用于本地测试api的ip
		//ip = "111.206.214.37"

		// 调用百度地图api找地址，记得改成自己的ak
		ak := "我的ak"

		// 服务地址
		host := "https://api.map.baidu.com"

		// 接口地址
		uri := "/location/ip"

		// 设置请求参数
		params := url.Values{
			"ip":   []string{ip},
			"coor": []string{"d09llb"},
			"ak":   []string{ak},
		}

		// 发起请求
		request, err := url.Parse(host + uri + "?" + params.Encode())
		if nil != err {
			fmt.Printf("host error: %v", err)
			return
		}

		resp1, err1 := http.Get(request.String())
		defer resp1.Body.Close()
		if err1 != nil {
			fmt.Printf("request error: %v", err1)
			return
		}
		body, err2 := ioutil.ReadAll(resp1.Body)
		if err2 != nil {
			fmt.Printf("response error: %v", err2)
		}
		//测试
		fmt.Println(string(body))
		//返回参数对应数据结构
		var response struct {
			Address string `json:"address""`
			Status  int    `json:"status"`
			Content struct {
				Address       string `json:"address"`
				AddressDetail struct {
					City     string `json:"city"`
					CityCode string `json:"cityCode"`
					Province string `json:"province"`
				} `json:"address_detail"`
				Point struct {
					X string `json:"x"`
					Y string `json:"y"`
				} `json:"point"`
			} `json:"content"`
		}

		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalf("解析 JSON 失败：%v", err)
		}

		// 检查响应状态码
		if response.Status != 0 {
			log.Fatalf("百度地图API 请求错误，状态码：%d", response.Status)
		}

		//存储头像图片并返回url

		file, fileHeader, err := r.FormFile("avatar")
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		defer func(file multipart.File) {
			_ = file.Close()
		}(file)

		// 判断是否为图片
		tmpFile, err := fileHeader.Open()
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, tmpFile); err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		if !filetype.IsImage(buf.Bytes()) {
			httpx.Error(w, apiErr.FileIsNotImage)
			return
		}
		if err = tmpFile.Close(); err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}

		// 使用uuid重新生成文件名
		fileName := utils.GetUUID() + filepath.Ext(fileHeader.Filename)

		// 存储到oss
		ok, err := oss.UploadImageToOss(svcCtx.AliyunClient, svcCtx.Config.OSS.BucketName, fileName, file)
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		} else if !ok {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails("upload image to oss error"))
			return
		}

		AvatarUrl := oss.GetOssImgUrl(svcCtx.Config.OSS, fileName)

		//要加还是可以加一个参数检查

		// 调用rpc服务
		_, err = svcCtx.UserRpc.CreateUser(r.Context(), &user.CreateUserRequest{
			Name:      req.Username,
			Password:  req.Password,
			Collage:   req.Collage,
			IdCard:    req.IdCard,
			Address:   response.Content.Address,
			Latitude:  response.Content.Point.X,
			Longitude: response.Content.Point.Y,
			AvatarUrl: AvatarUrl,
		})

		if err != nil {
			logx.WithContext(r.Context()).Errorf("CreateUser rpc error: %v", err)
			httpx.Error(w, apiErr.InternalError(r.Context(), err.Error()))
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
