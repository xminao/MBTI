package logic

import (
	"backend/app/user/api/internal/svc"
	"backend/app/user/api/internal/types"
	"backend/app/user/rpc/userrpc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetuserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetuserinfoLogic {
	return &GetuserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserinfoLogic) Getuserinfo(req *types.GetUserInfoReq) (*types.GetUserInfoResp, error) {

	user, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &userrpc.GetUserInfoReq{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	userinfo := types.UserInfo{
		Username:         user.Username,
		Nickname:         user.Nickname,
		Gender:           user.Gender,
		BindingStudentId: user.BindingStudentId,
	}

	return &types.GetUserInfoResp{
		UserInfoResp: userinfo,
	}, nil
}
