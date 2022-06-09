package logic

import (
	"backend/app/user/api/internal/svc"
	"backend/app/user/api/internal/types"
	"backend/app/user/rpc/userrpc"
	"backend/util/xerr"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetuserlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetuserlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetuserlistLogic {
	return &GetuserlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetuserlistLogic) Getuserlist(req *types.GetUserListReq) (*types.GetUserListResp, error) {
	username := l.ctx.Value("userName")
	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}
	if l.ctx.Value("authGroup") != "admin" {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "没有权限进行操作")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致
	r, err := l.svcCtx.UserRpc.GetUserList(l.ctx, &userrpc.GetUserListReq{})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	var resp []types.UserInfo
	if len(r.UserList) > 0 {
		for _, item := range r.UserList {
			var temp types.UserInfo
			_ = copier.Copy(&temp, item)
			resp = append(resp, temp)
		}
	}

	return &types.GetUserListResp{
		UserList: resp,
	}, nil
}
