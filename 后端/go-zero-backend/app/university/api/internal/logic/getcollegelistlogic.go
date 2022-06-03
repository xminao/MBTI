package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"backend/app/university_goctl/api/internal/svc"
	"backend/app/university_goctl/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollegeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCollegeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollegeListLogic {
	return &GetCollegeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCollegeListLogic) GetCollegeList(req *types.GetCollegeListReq) (*types.GetCollegeListResp, error) {

	whereBuilder := l.svcCtx.CollegeInfoModel.RowBuilder()
	//getlist from select type []*models.CollegeInfo
	collegeList, err := l.svcCtx.CollegeInfoModel.GetList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	// need trans, from CollegeInfo to College (type equal,but name not)
	//
	var resp []types.College
	if len(collegeList) > 0 {
		for _, item := range collegeList {
			var temp types.College
			_ = copier.Copy(&temp, item)

			resp = append(resp, temp)
		}
	}

	return &types.GetCollegeListResp{
		CollegeList: resp,
	}, nil
}
