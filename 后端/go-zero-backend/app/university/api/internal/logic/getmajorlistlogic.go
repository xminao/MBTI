package logic

import (
	"context"

	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMajorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMajorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMajorListLogic {
	return &GetMajorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMajorListLogic) GetMajorList(req *types.GetMajorListReq) (*types.GetMajorListResp, error) {
	whereBuilder := l.svcCtx.MajorInfoModel.RowBuilder()
	//getlist from select type []*models.CollegeInfo
	majorList, err := l.svcCtx.MajorInfoModel.GetList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	// need trans, from CollegeInfo to College (type equal,but name not)
	//
	var resp []types.Major
	if len(majorList) > 0 {
		for _, item := range majorList {
			//user college id to get
			if item.YearId == req.YearId {
				resp = append(resp, types.Major{
					MajorId:   item.MajorId,
					MajorName: item.MajorName,
				})
			}
		}
	}

	return &types.GetMajorListResp{
		MajorList: resp,
	}, nil
}
