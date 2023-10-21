package logic

import (
	"context"
	"errors"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	ur := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity= ?", req.UserRepositoryIdentity).Get(ur)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("没有找到 userRepositoryIdentity")
	}
	data := &models.ShareBasic{
		RepositoryIdentity: ur.RepositoryIdentity,
		ExpiredTime:        req.ExpiredTime,
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}
	resp = new(types.ShareBasicCreateReply)
	resp.Identity = data.Identity

	return
}
