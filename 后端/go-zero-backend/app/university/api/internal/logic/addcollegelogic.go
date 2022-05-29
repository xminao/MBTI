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
	if len(strings.TrimSpace(req.CollegeName)) == 0 {
		return nil, errors.New("学院名字不能为空")
	}

	data := models.CollegeInfo{
		CollegeId:   req.CollegeId,
		CollegeName: req.CollegeName,
	}

	err := l.svcCtx.DbEngin.Create(&data).Error

	if err != nil {
		return nil, err
	}

	return &types.AddCollegeResp{
		Msg: "插入成功",
	}, nil
}
