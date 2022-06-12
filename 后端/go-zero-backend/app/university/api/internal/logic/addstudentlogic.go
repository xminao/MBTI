package logic

import (
	"backend/app/university/models"
	"context"
	"errors"
	"strings"
	"time"

	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStudentLogic {
	return &AddStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStudentLogic) AddStudent(req *types.AddStudentReq) (resp *types.AddStudentResp, err error) {
	if l.ctx.Value("authGroup") != "admin" {
		return nil, errors.New("非管理员用户")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致
	if len(strings.TrimSpace(req.StudentId)) == 0 {
		return nil, errors.New("学号不能为空")
	}
	if len(strings.TrimSpace(req.StudentName)) == 0 {
		return nil, errors.New("姓名不能为空")
	}

	student := new(models.StudentInfo)
	student.StudentId = req.StudentId     //学号
	student.StudentName = req.StudentName //姓名
	student.CollegeId = req.CollegeId     //学院
	student.YearId = req.YearId           //年级
	student.MajorId = req.MajorId         //专业
	student.ClassId = req.ClassId         //班级
	student.CreatedAt = time.Now()        //创建事件

	_, err = l.svcCtx.StudentInfoModel.Insert(l.ctx, student)
	if err != nil {
		return nil, err
	}

	return &types.AddStudentResp{
		Msg: "插入成功",
	}, nil
}
