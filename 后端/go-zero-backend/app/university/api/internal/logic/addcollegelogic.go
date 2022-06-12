package logic

import (
	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"
	"backend/app/university/models"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"
)

type AddCollegeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCollegeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollegeLogic {
	return &AddCollegeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCollegeLogic) AddCollege(req *types.AddCollegeReq) (*types.AddCollegeResp, error) {
	if l.ctx.Value("authGroup") != "admin" {
		return nil, errors.New("非管理员用户")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致

	if len(strings.TrimSpace(req.CollegeName)) == 0 {
		return nil, errors.New("学院名不能为空")
	}

	countBuilder := l.svcCtx.CollegeInfoModel.CountBuilder("*")
	count, err := l.svcCtx.CollegeInfoModel.FindCount(l.ctx, countBuilder)

	college := new(models.CollegeInfo)

	if count == 0 {
		college.CollegeId = 1
	} else {
		maxBuilder := l.svcCtx.CollegeInfoModel.MaxRowBuilder("created_at")
		LatestRecord, err := l.svcCtx.CollegeInfoModel.FindOneByQuery(l.ctx, maxBuilder)
		if err != nil {
			return nil, err
		}
		college.CollegeId = LatestRecord.CollegeId + 1
	}

	college.CollegeName = req.CollegeName
	college.CreatedAt = time.Now()

	_, err = l.svcCtx.CollegeInfoModel.Insert(l.ctx, college)
	if err != nil {
		return nil, err
	}

	return &types.AddCollegeResp{
		Msg: "插入成功",
	}, nil
}
