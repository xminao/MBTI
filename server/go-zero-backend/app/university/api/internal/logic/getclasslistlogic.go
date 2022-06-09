package logic

import (
	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"
	"backend/util/xerr"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassListLogic {
	return &GetClassListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassListLogic) GetClassList(req *types.GetClassListReq) (*types.GetClassListResp, error) {
	username := l.ctx.Value("userName")
	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}
	if l.ctx.Value("authGroup") != "admin" {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "没有权限进行操作")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致
	whereBuilder := l.svcCtx.ClassInfoModel.RowBuilder()
	//getlist from select type []*models.CollegeInfo
	classList, err := l.svcCtx.ClassInfoModel.GetList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	// need trans, from CollegeInfo to College (type equal,but name not)
	//

	var resp []types.Class
	if len(classList) > 0 {
		for _, item := range classList {
			//user college id to get
			if item.MajorId == req.MajorId {
				resp = append(resp, types.Class{
					ClassId:   item.ClassId,
					ClassName: item.ClassName,
					CollegeId: item.CollegeId,
					YearId:    item.YearId,
					MajorId:   item.MajorId,
				})
			}
		}
	}

	return &types.GetClassListResp{
		ClassList: resp,
	}, nil
}
