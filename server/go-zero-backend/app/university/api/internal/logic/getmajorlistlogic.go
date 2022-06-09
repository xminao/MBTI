package logic

import (
	"backend/util/xerr"
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
	username := l.ctx.Value("userName")
	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}
	if l.ctx.Value("authGroup") != "admin" {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "没有权限进行操作")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致
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
					CollegeId: item.MajorId,
					YearId:    item.YearId,
				})
			}
		}
	}

	return &types.GetMajorListResp{
		MajorList: resp,
	}, nil
}
