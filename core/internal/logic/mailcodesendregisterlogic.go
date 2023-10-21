package logic

import (
	"context"
	"errors"
	"time"

	"cloud-disk/core/helper"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeRequest) (resp *types.MailCodeReply, err error) {
	//该邮箱未被注册
	cnt, err := l.svcCtx.Engine.Where("email=?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("该邮箱已被注册")
	}
	//储存验证码
	code := helper.GetRand()
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*300)
	err = helper.SendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}
