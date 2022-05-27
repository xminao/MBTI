package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
	"user/api/internal/models"
	"user/api/internal/svc"
	"user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.LoginReq) (*types.LoginResp, error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("用户名和密码不能为空")
	}

	// 用gorm查询数据
	user := models.Userinfo{}
	unexist := l.svcCtx.DbEngin.Where("name = ?", req.Username).First(&user).RecordNotFound()
	fmt.Println(user)

	//判断用户名密码
	if unexist {
		return nil, errors.New("用户名不存在/密码错误")
	}
	if req.Password != user.Password {
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

	return &types.LoginResp{
		Username:     user.Name,
		Age:          int(user.Age),
		Gender:       user.Gender,
		Token:        token,
		ExpireTime:   now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil

}

func (l *GetUserLogic) getJwtToken(key string, iat, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(key))
}
