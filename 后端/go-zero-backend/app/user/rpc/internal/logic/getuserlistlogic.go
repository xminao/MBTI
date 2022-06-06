package logic

import (
	"backend/app/user/rpc/internal/svc"
	"backend/app/user/rpc/user"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *user.GetUserListReq) (*user.GetUserListResp, error) {
	//加一个jwt管理员验证得
	whereBuilder := l.svcCtx.UserInfoModel.RowBuilder()

	userList, err := l.svcCtx.UserInfoModel.GetUserList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	// need trans, from CollegeInfo to College (type equal,but name not)
	//
	var resp []*user.User
	if len(userList) > 0 {
		for _, item := range userList {
			var temp user.User
			_ = copier.Copy(&temp, item)
			resp = append(resp, &temp)
		}
	}

	return &user.GetUserListResp{
		UserList: resp,
	}, nil
}
