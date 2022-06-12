package logic

import (
	"backend/app/data/api/internal/svc"
	"backend/app/data/api/internal/types"
	"backend/app/data/models"
	"backend/util/xerr"
	"context"
	"encoding/json"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTestDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTestDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTestDataLogic {
	return &AddTestDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type SelectInfo struct {
	Id         int64  `json:"id"`
	Desc       string `json:"desc"`
	Option     string `json:"option"`
	OptionDesc string `json:"option_desc"`
}

func (l *AddTestDataLogic) AddTestData(req *types.AddTestDataReq) (*types.AddTestDataResp, error) {
	username := l.ctx.Value("userName")

	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}

	var c []SelectInfo
	err := json.Unmarshal([]byte(req.Selection), &c)
	save, err := json.Marshal(c)

	countBuilder := l.svcCtx.TestDataModel.CountBuilder("*")
	count, err := l.svcCtx.TestDataModel.FindCount(l.ctx, countBuilder)

	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	data := new(models.TestData)

	if count == 0 {
		data.Id = 1
	} else {
		maxBuilder := l.svcCtx.TestDataModel.MaxRowBuilder("created_at")
		LatestRecord, err := l.svcCtx.TestDataModel.FindOneByQuery(l.ctx, maxBuilder)
		if err != nil {
			return nil, err
		}
		data.Id = LatestRecord.Id + 1
	}

	//var sel SelectInfo
	//err = json.Unmarshal([]byte(req.Selection), &sel)
	//if err != nil {
	//	return nil, err
	//}
	//
	//fmt.Println(sel)

	data.Username = req.Username
	data.Type = req.Type
	data.Selection = string(save)
	data.StudentId = req.StudentId
	data.CreatedAt = time.Now()

	_, err = l.svcCtx.TestDataModel.Insert(l.ctx, data)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	return &types.AddTestDataResp{Msg: "添加成功"}, nil
}
