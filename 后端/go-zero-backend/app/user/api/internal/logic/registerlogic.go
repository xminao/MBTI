package logic

import (
	"backend/app/user/models"
	"backend/util/xerr"
	"context"
	"errors"
	"strings"
	"time"

	"backend/app/user/api/internal/svc"
	"backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	// unsolute :
	// nickname judge, username judge
	if len(strings.TrimSpace(req.Username)) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(strings.TrimSpace(req.NickName)) == 0 {
		return nil, errors.New("昵称不能为空")
	}
	if len(strings.TrimSpace(req.Password)) < 6 {
		return nil, errors.New("密码位数不小于6位")
	}
	if len(strings.TrimSpace(req.Gender)) == 0 {
		return nil, errors.New("性别不能为空")
	}
	if req.Gender != "男" && req.Gender != "女" {
		return nil, errors.New("性别只能为男/女")
	}

	//type models.UserInfo
	user, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, req.Username)
	// not found
	// simple , need to change
	if err != nil && err != models.ErrNotFound {
		return nil, errors.New(err.Error())
	}

	if user != nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "用户名已经被注册咯，换一个吧")
	}

	stu, err := l.svcCtx.UserInfoModel.FindOneByBindingStudentId(l.ctx, req.BindingStudentId)
	if stu != nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "学号已经绑定别的用户啦")
	}

	_, err = l.svcCtx.UserInfoModel.Insert(l.ctx, &models.UserInfo{
		Username:         req.Username,
		Nickname:         req.NickName,
		Password:         req.Password,
		Gender:           req.Gender,
		CreatedAt:        time.Now(),
		UpdateAt:         time.Now(),
		AuthGroup:        "default",
		BindingStudentId: req.BindingStudentId,
	})
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	return &types.RegisterResp{
		Nickname: req.NickName,
		Msg:      "注册成功",
	}, nil
}
