package logic

import (
	"backend/util/xerr"
	"context"

	"backend/app/data/api/internal/svc"
	"backend/app/data/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLatestDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLatestDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLatestDataLogic {
	return &GetLatestDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLatestDataLogic) GetLatestData(req *types.GetLatestDataReq) (resp *types.GetLatestDataResp, err error) {
	username := l.ctx.Value("userName")

	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}

	maxBuilder := l.svcCtx.TestDataModel.MaxRowBuilderByUser("created_at", username.(string))
	LatestRecord, err := l.svcCtx.TestDataModel.FindOneByQuery(l.ctx, maxBuilder)

	if err != nil {
		return nil, xerr.NewErrCode(xerr.RESULT_NOT_FOUND)
		//return nil, xerr.NewErrCodeMsg(xerr.DATA_ERROR, "没有记录哟，快去测试吧")
	}

	return &types.GetLatestDataResp{Type: LatestRecord.Type}, nil
}
