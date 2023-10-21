package logic

import (
	"context"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest, authorization string) (resp *types.RefreshAuthorizationSaveReply, err error) {
	uc, err := helper.AnalyzeToken(authorization)
	if err != nil {
		return nil, err
	}
	token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, 3600)
	if err != nil {
		return nil, err
	}
	refreshToken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, 7200)
	if err != nil {
		return nil, err
	}
	resp = new(types.RefreshAuthorizationSaveReply)
	resp.Token = token
	resp.RefreshToken = refreshToken

	return
}
