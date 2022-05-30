package logic

import (
	"backend/app/university_goctl/api/internal/svc"
	"backend/app/university_goctl/api/internal/types"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
)

type AddCollegeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCollegeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollegeLogic {
	return &AddCollegeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCollegeLogic) AddCollege(req *types.AddCollegeReq) (*types.AddCollegeResp, error) {

	if len(strings.TrimSpace(req.CollegeName)) == 0 {
		return nil, errors.New("学院名不能为空")
	}

	maxBuilder := l.svcCtx.CollegeInfoModel.RowBuilder()
	data, err := l.svcCtx.CollegeInfoModel.FindOneByQuery(l.ctx, maxBuilder)
	fmt.Println(data)
	if err != nil {
		return nil, nil
	}
	return nil, nil

	//countBuilder := l.svcCtx.CollegeInfoModel.CountBuilder("*")
	//count, err := l.svcCtx.CollegeInfoModel.FindCount(l.ctx, countBuilder)
	//fmt.Println(count)
	//college := new(models.CollegeInfo)
	//// 获取最新的ID，并+1
	//if count == 0 {
	//	college.CollegeId = 1
	//} else {
	//	maxBuilder := l.svcCtx.CollegeInfoModel.CountBuilder("*")
	//	LatestRecord, err := l.svcCtx.CollegeInfoModel.FindOneByQuery(l.ctx, maxBuilder)
	//	if err != nil {
	//		return nil, err
	//	}
	//	college.CollegeId = LatestRecord.CollegeId + 1
	//}
	//
	//college.CollegeName = req.CollegeName
	//college.CreatedAt = time.Now()
	//
	//_, err = l.svcCtx.CollegeInfoModel.Insert(l.ctx, college)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &types.AddCollegeResp{
	//	Msg: "插入成功",
	//}, nil
}
