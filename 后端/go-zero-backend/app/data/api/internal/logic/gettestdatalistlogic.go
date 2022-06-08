package logic

import (
	"backend/app/data/api/internal/svc"
	"backend/app/data/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTestDataListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTestDataListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTestDataListLogic {
	return &GetTestDataListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTestDataListLogic) GetTestDataList(req *types.GetTestDataListReq) (*types.GetTestDataListResp, error) {
	whereBuilder := l.svcCtx.TestDataModel.RowBuilder()

	userList, err := l.svcCtx.TestDataModel.GetDataList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	// need trans, from CollegeInfo to College (type equal,but name not)
	//
	var resp []types.Data
	if len(userList) > 0 {
		for _, item := range userList {
			var temp types.Data
			//fmt.Println(item)
			//_ = copier.Copy(&temp, item.)
			// 我真他妈服啦，time类型 copier不能复制\
			temp.Username = item.Username
			//学号
			temp.Time = item.CreatedAt.String()
			temp.Type = item.Type
			resp = append(resp, temp)
		}
	}

	return &types.GetTestDataListResp{List: resp}, nil
}
