package handler

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		//判断文件是否存在
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)
		has, err := svcCtx.Engine.Where("hash=?", hash).Get(rp)
		if err != nil {
			return
		}
		if has {
			httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}
		//往cos上传文件
		cosPath, err := helper.CosUpLoad(r)
		if err != nil {
			return
		}
		//往logic中传递数据
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
