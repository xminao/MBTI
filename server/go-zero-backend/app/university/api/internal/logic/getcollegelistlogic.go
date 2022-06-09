package logic

import (
	"backend/util/xerr"
	"context"
	"github.com/jinzhu/copier"

	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"

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
	username := l.ctx.Value("userName")
	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}
	if l.ctx.Value("authGroup") != "admin" {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "没有权限进行操作")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致
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
