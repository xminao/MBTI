package logic

import (
	"backend/app/university/api/internal/svc"
	"backend/app/university/api/internal/types"
	"backend/app/university/rpc/universityrpc"
	"backend/util/xerr"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentInfoLogic {
	return &GetStudentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentInfoLogic) GetStudentInfo(req *types.GetStudentInfoReq) (*types.GetStudentInfoResp, error) {
	stu, err := l.svcCtx.UniversityRpc.GetStudentInfo(l.ctx, &universityrpc.GetStudentInfoReq{
		StudentId: req.StudentId,
	})

	if err != nil {
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}

	var temp = types.Student{}
	temp.StudentName = stu.StudentInfo.StudentName
	temp.StudentId = stu.StudentInfo.StudentId
	temp.College = stu.StudentInfo.College
	temp.Year = stu.StudentInfo.Year
	temp.Major = stu.StudentInfo.Major
	temp.Class = stu.StudentInfo.Class
	temp.CreatedAt = stu.StudentInfo.CreatedAt

	return &types.GetStudentInfoResp{StudentInfo: temp}, nil
}
