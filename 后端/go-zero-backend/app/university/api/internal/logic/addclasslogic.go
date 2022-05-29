package logic

import (
	"backend/app/university/models"
	"context"
	"errors"
	"strings"

	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassLogic {
	return &AddClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassLogic) AddClass(req *types.AddClassReq) (*types.AddClassResp, error) {
	if len(strings.TrimSpace(req.ClassName)) == 0 {
		return nil, errors.New("班级名字不能为空")
	}

	data := models.ClassInfo{
		ClassId:     req.ClassId,
		ClassName:   req.ClassName,
		CollegeInfo: models.CollegeInfo{},
		CollegeId:   req.CollegeId,
		YearInfo:    models.YearInfo{},
		Year:        req.Year,
		MajorInfo:   models.MajorInfo{},
		MajorId:     req.MajorId,
	}

	err := l.svcCtx.DbEngin.Create(&data).Error

	if err != nil {
		return nil, err
	}

	return &types.AddClassResp{
		Msg: "插入成功",
	}, nil
}
