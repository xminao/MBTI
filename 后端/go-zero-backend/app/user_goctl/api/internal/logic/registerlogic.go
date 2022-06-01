package logic

import (
	"backend/app/user_goctl/models"
	"context"
	"errors"
	"strings"
	"time"

	"backend/app/user_goctl/api/internal/svc"
	"backend/app/user_goctl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	// unsolute :
	// nickname judge, username judge
	if len(strings.TrimSpace(req.Username)) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(strings.TrimSpace(req.NickName)) == 0 {
		return nil, errors.New("昵称不能为空")
	}
	if len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("密码不能为空")
	}
	if len(strings.TrimSpace(req.Gender)) == 0 {
		return nil, errors.New("性别不能为空")
	}

	//type models.UserInfo
	_, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, req.Username)
	// not found
	// simple , need to change
	if err != nil && err != models.ErrNotFound {
		return nil, errors.New(err.Error())
	}

	_, err = l.svcCtx.UserInfoModel.Insert(l.ctx, &models.UserInfo{
		Username:  req.Username,
		Nickname:  req.NickName,
		Password:  req.Password,
		Gender:    req.Gender,
		CreatedAt: time.Now(),
		AuthGroup: "default",
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Msg: "注册成功",
	}, nil
}
