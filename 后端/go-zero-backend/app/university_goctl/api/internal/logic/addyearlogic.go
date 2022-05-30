package logic

import (
	"context"

	"backend/app/university_goctl/api/internal/svc"
	"backend/app/university_goctl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddYearLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddYearLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddYearLogic {
	return &AddYearLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddYearLogic) AddYear(req *types.AddYearReq) (resp *types.AddYearResp, err error) {
	// todo: add your logic here and delete this line

	return
}
