package logic

import (
	"context"
	"errors"

	"backend/app/data/api/internal/svc"
	"backend/app/data/api/internal/types"
	"backend/app/data/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataLogic {
	return &GetDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataLogic) GetData(req *types.GetTestDataReq) (resp *types.GetTestDataResp, err error) {
	question, err := l.svcCtx.TestDataModel.FindOne(l.ctx, int64(req.DataId))
	if err == models.ErrNotFound {
		return nil, errors.New("数据不存在")
	}

	data := types.Data{
		Username:  question.Username,
		StudentId: question.StudentId,
		Type:      question.Type,
		Time:      question.CreatedAt.String(),
		Selection: question.Selection,
	}

	return &types.GetTestDataResp{
		DataInfo: data,
	}, nil
}
