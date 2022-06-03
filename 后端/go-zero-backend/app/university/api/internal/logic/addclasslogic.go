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

func (l *AddClassLogic) AddClass(req *types.AddClassReq) (resp *types.AddClassResp, err error) {
	if len(strings.TrimSpace(req.ClassName)) == 0 {
		return nil, errors.New("年级不能为空")
	}

	countBuilder := l.svcCtx.ClassInfoModel.CountBuilder("*")
	count, err := l.svcCtx.ClassInfoModel.FindCount(l.ctx, countBuilder)

	class := new(models.ClassInfo)

	if count == 0 {
		class.ClassId = 1
	} else {
		maxBuilder := l.svcCtx.ClassInfoModel.MaxRowBuilder("created_at")
		LatestRecord, err := l.svcCtx.ClassInfoModel.FindOneByQuery(l.ctx, maxBuilder)
		if err != nil {
			return nil, err
		}
		class.ClassId = LatestRecord.ClassId + 1
	}

	class.ClassName = req.ClassName
	class.CollegeId = req.CollegeId
	class.YearId = req.YearId
	class.MajorId = req.MajorId
	class.CreatedAt = time.Now()

	_, err = l.svcCtx.ClassInfoModel.Insert(l.ctx, class)
	if err != nil {
		return nil, err
	}

	return &types.AddClassResp{
		Msg: "插入成功",
	}, nil
}
