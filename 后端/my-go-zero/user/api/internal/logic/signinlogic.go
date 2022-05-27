package logic

import (
	"context"
	"errors"
	"strings"
	"user/api/internal/models"
	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

//添加svcCtx *svc.ServiceContext,因为需要用到里面的DbEngin
func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//填充业务逻辑
func (l *SignInLogic) SignIn(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("用户名和密码不能为空")
	}

	//插入数据
	user := models.Userinfo{
		Name:     req.Username,
		Password: req.Password,
		Age:      int64(req.Age),
		Gender:   req.Gender,
	}
	l.svcCtx.DbEngin.Create(&user)

	return &types.RegisterResp{
		Msg: "用户注册成功",
	}, nil
}
