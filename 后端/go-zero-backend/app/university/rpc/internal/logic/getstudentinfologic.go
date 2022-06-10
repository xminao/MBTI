package logic

import (
	"backend/util/xerr"
	"context"
	"encoding/json"
	"time"

	"backend/app/university/rpc/internal/svc"
	"backend/app/university/rpc/university"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentInfoLogic {
	return &GetStudentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentInfoLogic) GetStudentInfo(in *university.GetStudentInfoReq) (*university.GetStudentInfoResp, error) {
	stu, err := l.svcCtx.StudentInfoModel.FindOne(l.ctx, in.StudentId)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
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

	college, _ := json.Marshal(College{
		CollegeId:   stu.CollegeId,
		CollegeName: l.svcCtx.CollegeInfoModel.GetCollegeName(l.ctx, stu.CollegeId),
	})

	year, _ := json.Marshal(Year{
		YearId:   stu.YearId,
		YearName: l.svcCtx.YearInfoModel.GetYearName(l.ctx, stu.YearId),
	})

	major, _ := json.Marshal(Major{
		MajorId:   stu.MajorId,
		MajorName: l.svcCtx.MajorInfoModel.GetMajorName(l.ctx, stu.MajorId),
	})

	class, _ := json.Marshal(Class{
		ClassId:   stu.ClassId,
		ClassName: l.svcCtx.ClassInfoModel.GetClassName(l.ctx, stu.ClassId),
	})

	var resp = university.Student{
		StudentId:   stu.StudentId,
		StudentName: stu.StudentName,
		College:     string(college),
		Year:        string(year),
		Major:       string(major),
		Class:       string(class),
		CreatedAt:   time.Now().Unix(),
	}

	return &university.GetStudentInfoResp{
		StudentInfo: &resp,
	}, nil
}
