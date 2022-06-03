package logic

import (
	"backend/app/university_goctl/models"
	"context"
	"errors"
	"strings"
	"time"

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
	if len(strings.TrimSpace(req.MajorName)) == 0 {
		return nil, errors.New("年级不能为空")
	}

	countBuilder := l.svcCtx.MajorInfoModel.CountBuilder("*")
	count, err := l.svcCtx.MajorInfoModel.FindCount(l.ctx, countBuilder)

	major := new(models.MajorInfo)

	if count == 0 {
		major.MajorId = 1
	} else {
		maxBuilder := l.svcCtx.MajorInfoModel.MaxRowBuilder("created_at")
		LatestRecord, err := l.svcCtx.MajorInfoModel.FindOneByQuery(l.ctx, maxBuilder)
		if err != nil {
			return nil, err
		}
		major.MajorId = LatestRecord.MajorId + 1
	}

	major.MajorName = req.MajorName
	major.CollegeId = req.CollegeId
	major.YearId = req.YearId
	major.CreatedAt = time.Now()

	_, err = l.svcCtx.MajorInfoModel.Insert(l.ctx, major)
	if err != nil {
		return nil, err
	}

	return &types.AddMajorResp{
		Msg: "插入成功",
	}, nil
}
