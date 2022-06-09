package logic

import (
	"backend/util/xerr"
	"context"

	"backend/app/user/api/internal/svc"
	"backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyTokenLogic) VerifyToken(req *types.VerifyTokenReq) (*types.VerifyTokenResp, error) {
	username := l.ctx.Value("userName")

	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}

	return &types.VerifyTokenResp{
		Username: username.(string),
	}, nil
}
