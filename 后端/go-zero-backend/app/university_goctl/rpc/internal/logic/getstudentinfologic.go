package logic

import (
	"context"

	"backend/app/university_goctl/rpc/internal/svc"
	"backend/app/university_goctl/rpc/university"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentInfoLogic {
	return &GetStudentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentInfoLogic) GetStudentInfo(in *university.GetStudentInfoReq) (*university.GetStudentInfoResp, error) {
	// todo: add your logic here and delete this line

	return &university.GetStudentInfoResp{}, nil
}
