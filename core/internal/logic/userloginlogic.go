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

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// 从数据库里查询用户
	ub := new(models.UserBasic)
	has, err := l.svcCtx.Engine.Where("name=? AND password=?", req.Name, helper.Md5(req.Password)).Get(ub)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户名或密码错误")
	}
	token, err := helper.GenerateToken(int64(ub.Id), ub.Identity, ub.Name, 20)
	if err != nil {
		return nil, err
	}
	//生成一个用于刷新token的token
	refreshToken, err := helper.GenerateToken(int64(ub.Id), ub.Identity, ub.Name, 7200)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}
