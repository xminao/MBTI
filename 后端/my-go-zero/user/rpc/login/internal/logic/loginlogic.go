package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"user/rpc/login/internal/models"

	"user/rpc/login/internal/svc"
	"user/rpc/login/login"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginResp, error) {
	// 用gorm查询数据
	user := models.Userinfo{}
	unexist := l.svcCtx.DbEngin.Where("name = ?", in.Username).First(&user).RecordNotFound()
	fmt.Println(user)

	//判断用户名密码
	if unexist {
		return nil, errors.New("用户名不存在/密码错误")
	}
	if in.Password != user.Password {
		return nil, errors.New("用户名不存在/密码错误")
	}

	//Jwt
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Jwt.AccessExpire
	accessSecret := l.svcCtx.Config.Jwt.AccessSecret
	token, err := l.getJwtToken(accessSecret,
		now,
		accessExpire)
	if err != nil {
		return nil, err
	}

	return &login.LoginResp{
		Username:     user.Name,
		Age:          int64(user.Age),
		Gender:       user.Gender,
		Token:        token,
		ExpireTime:   now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) getJwtToken(key string, iat, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(key))
}
