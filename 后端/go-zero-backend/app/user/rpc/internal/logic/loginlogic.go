package logic

import (
	"backend/app/user/rpc/internal/models"
	"context"
	"errors"
	"strings"

	"backend/app/user/rpc/internal/svc"
	"backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// 判断传入值
	if len(strings.TrimSpace(in.Username)) == 0 || len(strings.TrimSpace(in.Password)) == 0 {
		return nil, errors.New("用户名/昵称不能为空")
	}

	// 判断是否存在/密码是否正确正确
	userInfo := models.Userinfo{}
	notExist := l.svcCtx.DbEngin.Where("username = ?", in.Username).First(&userInfo).RecordNotFound()

	if notExist {
		return nil, errors.New("用户不存在")
	}
	if in.Password != userInfo.Password {
		return nil, errors.New("密码不正确")
	}

	return &user.LoginResp{
		Username: userInfo.Username,
		Nickname: userInfo.Nickname,
		Gender:   userInfo.Gender,
	}, nil
}
