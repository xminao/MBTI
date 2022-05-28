package logic

import (
	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"
	"backend/app/university/models"
	"context"
	"fmt"
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

	// 扫描到college结构体数组
	var collegeData []types.College

	err := l.svcCtx.DbEngin.Model(&models.CollegeInfo{}).Scan(&collegeData).Error

	fmt.Println(collegeData)
	if err != nil {
		return nil, err
	}

	return &types.CollegeListResp{
		List: collegeData,
	}, nil
}
