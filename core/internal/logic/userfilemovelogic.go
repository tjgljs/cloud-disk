package logic

import (
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	//查parentid
	ParentData := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("user_identity= ? AND identity= ?", userIdentity, req.ParentIdentity).Get(ParentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件不存在")
	}
	//移动文件 更新parentid
	_, err = l.svcCtx.Engine.Where("user_identity= ? AND identity= ?", userIdentity, req.Identity).Update(models.UserRepository{
		ParentId: int64(ParentData.Id),
	})
	if err != nil {
		return nil, err
	}
	return
}
