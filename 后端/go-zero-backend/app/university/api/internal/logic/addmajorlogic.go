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

func (l *AddMajorLogic) AddMajor(req *types.AddMajorReq) (*types.AddMajorResp, error) {
	if len(strings.TrimSpace(req.MajorName)) == 0 {
		return nil, errors.New("专业名字不能为空")
	}

	data := models.MajorInfo{
		MajorId:     req.MajorId,
		MajorName:   req.MajorName,
		CollegeInfo: models.CollegeInfo{},
		CollegeId:   req.CollegeId,
		YearInfo:    models.YearInfo{},
		Year:        req.Year,
	}

	err := l.svcCtx.DbEngin.Create(&data).Error

	if err != nil {
		return nil, err
	}

	return &types.AddMajorResp{
		Msg: "插入成功",
	}, nil
}
