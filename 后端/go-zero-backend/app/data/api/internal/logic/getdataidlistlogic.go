package logic

import (
	"backend/util/xerr"
	"context"

	"backend/app/data/api/internal/svc"
	"backend/app/data/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataIdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataIdListLogic {
	return &GetDataIdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataIdListLogic) GetDataIdList(req *types.GetDataIdListReq) (resp *types.GetDataIdListResp, err error) {
	build := l.svcCtx.TestDataModel.RowBuilder()
	idList, err := l.svcCtx.TestDataModel.GetIDList(l.ctx, build)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	return &types.GetDataIdListResp{
		DataIdList: idList,
	}, nil
}
