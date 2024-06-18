package task

import (
	"bytes"
	"github.com/h2non/filetype"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"mime/multipart"
	"net/http"
	"orientation-platform/common/error/apiErr"
	"orientation-platform/common/oss"
	"orientation-platform/common/utils"
	logic "orientation-platform/service/http/internal/logic/task"
	"orientation-platform/service/rpc/task/types/task"
	"orientation-platform/service/rpc/user/types/user"
	"path/filepath"

	"github.com/zeromicro/go-zero/rest/httpx"
	"orientation-platform/service/http/internal/svc"
	"orientation-platform/service/http/internal/types"
)

func AdminSetTaskHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminSetTaskRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//存储问题、参考答案图片并返回url

		file1, fileHeader1, err := r.FormFile("question_image")
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		defer func(file multipart.File) {
			_ = file.Close()
		}(file1)
		//defer使得过程中文件不管出现什么错误，最终都会被关闭，一般要处理文件时都要用到它

		file2, fileHeader2, err := r.FormFile("refer_answer_image")
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		defer func(file multipart.File) {
			_ = file.Close()
		}(file2)

		// 判断file1是否为图片
		tmpFile1, err := fileHeader1.Open()
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, tmpFile1); err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		if !filetype.IsImage(buf.Bytes()) {
			httpx.Error(w, apiErr.FileIsNotImage)
			return
		}
		if err = tmpFile1.Close(); err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		//file2
		tmpFile2, err := fileHeader2.Open()
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}

		buf = bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, tmpFile2); err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}
		if !filetype.IsImage(buf.Bytes()) {
			httpx.Error(w, apiErr.FileIsNotImage)
			return
		}
		if err = tmpFile1.Close(); err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		}

		// 使用uuid重新生成文件名
		fileName1 := utils.GetUUID() + filepath.Ext(fileHeader1.Filename)
		fileName2 := utils.GetUUID() + filepath.Ext(fileHeader2.Filename)

		// 存储到oss
		ok, err := oss.UploadImageToOss(svcCtx.AliyunClient, svcCtx.Config.OSS.BucketName, fileName1, file1)
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		} else if !ok {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails("upload image to oss error"))
			return
		}
		QuestionUrl := oss.GetOssImgUrl(svcCtx.Config.OSS, fileName1)

		ok, err = oss.UploadImageToOss(svcCtx.AliyunClient, svcCtx.Config.OSS.BucketName, fileName2, file2)
		if err != nil {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails(err.Error()))
			return
		} else if !ok {
			httpx.Error(w, apiErr.FileUploadFailed.WithDetails("upload image to oss error"))
			return
		}
		ReferAnswerUrl := oss.GetOssImgUrl(svcCtx.Config.OSS, fileName2)

		// 调用rpc服务
		_, err = svcCtx.TaskRpc.TaskSet(r.Context(), &task.TaskSetRequest{
			Collage:        req.Collage,
			Title:          req.Title,
			Text:           req.Text,
			Type:           req.Type,
			Bonus:          req.Bonus,
			QuestionUrl:    QuestionUrl,
			ReferAnswerUrl: ReferAnswerUrl,
		})

		if err != nil {
			logx.WithContext(r.Context()).Errorf("TaskSet rpc error: %v", err)
			httpx.Error(w, apiErr.InternalError(r.Context(), err.Error()))
			return
		}

		//先查询出对应collage的用户list，再进行给每一个对应学生初始化任务的rpc的TaskInit
		IdList, err := svcCtx.UserRpc.GetUserIdByCollage(r.Context(), &user.GetUserIdByCollageRequest{
			Collage: req.Collage,
		})

		if err != nil {
			logx.WithContext(r.Context()).Errorf("GetUserIdByCollage rpc error: %v", err)
			httpx.Error(w, apiErr.InternalError(r.Context(), err.Error()))
			return
		}

		for _, id := range IdList.IdList {
			_, err = svcCtx.TaskRpc.TaskInit(r.Context(), &task.TaskInitRequest{
				UserId: id,
				Collage: req.Collage,
			})

			if err != nil {
				logx.WithContext(r.Context()).Errorf("TaskInit rpc error: %v", err)
				httpx.Error(w, apiErr.InternalError(r.Context(), err.Error()))
				return
			}
		}

		l := logic.NewAdminSetTaskLogic(r.Context(), svcCtx)
		resp, err := l.AdminSetTask(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
