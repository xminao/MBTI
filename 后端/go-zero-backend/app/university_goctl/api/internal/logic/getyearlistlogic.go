package logic

import (
	"backend/app/university_goctl/api/internal/svc"
	"backend/app/university_goctl/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetYearListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetYearListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetYearListLogic {
	return &GetYearListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetYearListLogic) GetYearList(req *types.GetYearListReq) (*types.GetYearListResp, error) {
	whereBuilder := l.svcCtx.YearInfoModel.RowBuilder()
	//getlist from select type []*models.CollegeInfo
	yearList, err := l.svcCtx.YearInfoModel.GetList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	// need trans, from CollegeInfo to College (type equal,but name not)
	//
	var resp []types.Year
	if len(yearList) > 0 {
		for _, item := range yearList {
			//user college id to get
			if item.CollegeId == req.CollegeId {
				resp = append(resp, types.Year{
					YearId:   item.YearId,
					YearName: item.YearName,
				})
			}
		}
	}

	return &types.GetYearListResp{
		YearList: resp,
	}, nil
}
