package logic

import (
	"context"
	"errors"
	"strings"

	"backend/app/university/models"
	"backend/app/user/rpc/internal/svc"
	"backend/app/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	//加一个jwt管理员验证得
	if len(strings.TrimSpace(in.Username)) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	data, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, in.Username)
	if err != nil && err != models.ErrNotFound {
		return nil, err
	}
	if err == models.ErrNotFound {
		return nil, errors.New("用户不存在")
	}

	userinfo := &user.GetUserInfoResp{
		Username:         data.Username,
		Nickname:         data.Nickname,
		Gender:           data.Gender,
		BindingStudentId: data.BindingStudentId,
	}
	return userinfo, nil
}
