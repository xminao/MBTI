package logic

import (
	"backend/app/university/rpc/internal/models"
	"backend/app/university/rpc/internal/svc"
	"backend/app/university/rpc/university"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollegeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollegeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollegeListLogic {
	return &CollegeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollegeListLogic) CollegeList(in *university.CollegeListReq) (*university.CollegeListResp, error) {

	//转换为可返回的
	// models.CollegeInfo 数据库的表结构体
	// university.CollegeInfo proto自定义的resp结构体
	var collegeData []*university.College

	err := l.svcCtx.DbEngin.Model(&models.CollegeInfo{}).Scan(&collegeData).Error

	if err != nil {
		return nil, err
	}

	return &university.CollegeListResp{
		CollegeList: collegeData,
	}, nil
}
