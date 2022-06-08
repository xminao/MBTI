package logic

import (
	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"
	"context"
	"encoding/json"
	"time"

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

	type College struct {
		CollegeId   int64  `json:"college_id"`
		CollegeName string `json:"college_name"`
	}

	type Year struct {
		YearId   int64  `json:"year_id"`
		YearName string `json:"year_name"`
	}

	type Major struct {
		MajorId   int64  `json:"major_id"`
		MajorName string `json:"major_name"`
	}

	type Class struct {
		ClassId   int64  `json:"class_id"`
		ClassName string `json:"class_name"`
	}

	var resp []types.Student
	if len(studentList) > 0 {
		for _, item := range studentList {
			college, _ := json.Marshal(College{
				CollegeId:   item.CollegeId,
				CollegeName: l.svcCtx.CollegeInfoModel.GetCollegeName(l.ctx, item.CollegeId),
			})

			year, _ := json.Marshal(Year{
				YearId:   item.YearId,
				YearName: l.svcCtx.YearInfoModel.GetYearName(l.ctx, item.YearId),
			})

			major, _ := json.Marshal(Major{
				MajorId:   item.MajorId,
				MajorName: l.svcCtx.MajorInfoModel.GetMajorName(l.ctx, item.MajorId),
			})

			class, _ := json.Marshal(Class{
				ClassId:   item.ClassId,
				ClassName: l.svcCtx.ClassInfoModel.GetClassName(l.ctx, item.ClassId),
			})

			var temp = types.Student{
				StudentId:       item.StudentId,
				StudentName:     item.StudentName,
				College:         string(college),
				Year:            string(year),
				Major:           string(major),
				Class:           string(class),
				IsBinding:       false,
				BindingUsername: item.BindingUsername.String,
				CreatedAt:       time.Now().Unix(),
			}
			resp = append(resp, temp)
		}
	}

	return &types.GetStudentListResp{
		StudentList: resp,
	}, nil
}
