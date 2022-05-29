package logic

import (
	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"
	"backend/app/university/models"
	"context"
	"errors"

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

func (l *AddYearLogic) AddYear(req *types.AddYearReq) (*types.AddYearResp, error) {
	if req.Year < 0 {
		return nil, errors.New("年级不应小于0")
	}

	data := models.YearInfo{
		CollegeId: req.CollegeId,
		Year:      req.Year,
	}

	err := l.svcCtx.DbEngin.Create(&data).Error

	if err != nil {
		return nil, err
	}

	return &types.AddYearResp{
		Msg: "插入成功",
	}, nil
}
