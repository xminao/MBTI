package logic

import (
	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"
	"backend/app/university/rpc/university"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CollegeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollegeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollegeListLogic {
	return &CollegeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollegeListLogic) CollegeList(req *types.CollegeListReq) (*types.CollegeListResp, error) {
	// 请求- 获取resp
	collegeListResp, err := l.svcCtx.UniversityRpc.CollegeList(l.ctx, &university.CollegeListReq{})

	if err != nil {
		return nil, err
	}

	var resp types.CollegeListResp
	_ = copier.Copy(&resp, collegeListResp)

	return &resp, nil
}
