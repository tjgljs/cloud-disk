package logic

import (
	"context"
	"errors"
	"fmt"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	//判断验证码是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未发送验证码")
	}
	if code != req.Code {
		return nil, errors.New("验证码不正确")
	}
	//判断用户名是否已经存在
	cnt, err := l.svcCtx.Engine.Where("name=?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("用户名已存在")
	}
	//通过判断 加入数据库
	ub := models.UserBasic{
		Identity: helper.UUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	n, err := l.svcCtx.Engine.Insert(ub)
	if err != nil {
		return nil, err
	}
	fmt.Printf("n: %v\n", n)
	return
}
