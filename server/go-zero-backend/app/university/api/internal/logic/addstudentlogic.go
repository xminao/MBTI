package logic

import (
	"backend/app/university/models"
	"backend/util/xerr"
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
	username := l.ctx.Value("userName")
	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}
	if l.ctx.Value("authGroup") != "admin" {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "没有权限进行操作")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致
	if len(strings.TrimSpace(req.StudentId)) == 0 {
		return nil, errors.New("学号不能为空")
	}
	if len(strings.TrimSpace(req.StudentName)) == 0 {
		return nil, errors.New("姓名不能为空")
	}

	student := new(models.StudentInfo)
	student.StudentId = req.StudentId
	student.StudentName = req.StudentName
	student.CollegeId = req.CollegeId
	student.YearId = req.YearId
	student.MajorId = req.MajorId
	student.ClassId = req.ClassId
	student.CreatedAt = time.Now()

	_, err = l.svcCtx.StudentInfoModel.Insert(l.ctx, student)
	if err != nil {
		return nil, err
	}

	return &types.AddStudentResp{
		Msg: "插入成功",
	}, nil
}
