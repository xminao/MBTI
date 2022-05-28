package logic

import (
	"backend/app/user/api/internal/svc"
	"backend/app/user/api/internal/types"
	"backend/app/user/rpc/user"
	"context"
	"github.com/jinzhu/copier"

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
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Username: req.Username,
		Nickname: req.NickName,
		Password: req.Password,
		Gender:   req.Gender,
	})

	if err != nil {
		return nil, err
	}

	var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)

	return &resp, nil
}
