package logic

import (
	"context"

	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentListLogic) GetStudentList(req *types.GetStudentListReq) (*types.GetStudentListResp, error) {
	whereBuilder := l.svcCtx.StudentInfoModel.RowBuilder()
	//getlist from select type []*models.CollegeInfo
	studentList, err := l.svcCtx.StudentInfoModel.GetList(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}
	// need trans, from CollegeInfo to College (type equal,but name not)
	//
	var resp []types.Student
	if len(studentList) > 0 {
		for _, item := range studentList {
			//user college id to get
			if item.ClassId == req.ClassId {
				resp = append(resp, types.Student{
					StudentId:       item.StudentId,
					StudentName:     item.StudentName,
					CollegeId:       item.CollegeId,
					YearId:          item.YearId,
					MajorId:         item.MajorId,
					ClassId:         item.ClassId,
					IsBinding:       false,
					BindingUsername: "",
				})
			}
		}
	}

	return &types.GetStudentListResp{
		StudentList: resp,
	}, nil
}
