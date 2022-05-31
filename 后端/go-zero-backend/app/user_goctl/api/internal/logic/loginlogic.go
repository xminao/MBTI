package logic

import (
	"backend/app/user_goctl/models"
	"context"
	"errors"
	"strings"

	"backend/app/user_goctl/api/internal/svc"
	"backend/app/user_goctl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	if len(strings.TrimSpace(req.Username)) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("密码不能为空")
	}

	//type models.UserInfo
	user, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, req.Username)
	if err == models.ErrNotFound {
		return nil, errors.New("用户不存在")
	}
	if user.Password != req.Password {
		return nil, errors.New("密码错误")
	}

	return &types.LoginResp{
		Msg: "登录成功",
	}, nil
}
