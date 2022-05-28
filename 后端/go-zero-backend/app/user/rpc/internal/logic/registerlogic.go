package logic

import (
	"backend/app/user/rpc/internal/models"
	"backend/app/user/rpc/internal/svc"
	"backend/app/user/rpc/user"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册逻辑
func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line

	// 判断传入值
	if len(strings.TrimSpace(in.Username)) == 0 || len(strings.TrimSpace(in.Nickname)) == 0 {
		return nil, errors.New("用户名/昵称不能为空")
	}
	if len(strings.TrimSpace(in.Password)) < 7 {
		return nil, errors.New("密码应大于六位")
	}
	if strings.TrimSpace(in.Gender) != "男" && strings.TrimSpace(in.Gender) != "女" {
		return nil, errors.New("性别为男或女")
	}

	// 判断是否存在
	userinfo := models.Userinfo{}
	unexist := l.svcCtx.DbEngin.Where("username = ?", in.Username).First(&userinfo).RecordNotFound()

	if unexist {
		userinfo := models.Userinfo{
			Model:    gorm.Model{},
			Username: in.Username,
			Nickname: in.Nickname,
			Password: in.Password,
			Gender:   in.Gender,
		}
		//创建记录
		l.svcCtx.DbEngin.Create(&userinfo)
	} else {
		return nil, errors.New("用户已经存在")
	}

	return &user.RegisterResp{
		Msg: "注册成功",
	}, nil
}
