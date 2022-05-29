package logic

import (
	"context"

	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollegeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCollegeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollegeLogic {
	return &AddCollegeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCollegeLogic) AddCollege(req *types.AddCollegeReq) (resp *types.AddCollegeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
