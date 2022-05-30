package logic

import (
	"context"

	"backend/app/university_goctl/api/internal/svc"
	"backend/app/university_goctl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMajorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMajorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMajorLogic {
	return &AddMajorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMajorLogic) AddMajor(req *types.AddMajorReq) (resp *types.AddMajorResp, err error) {
	// todo: add your logic here and delete this line

	return
}
