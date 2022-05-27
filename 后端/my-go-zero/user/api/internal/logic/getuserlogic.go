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
	err := l.svcCtx.DbEngin.Where("name = ?", req.Username).First(&user).Error
	fmt.Println(user)
	if err != nil {
		panic(err)
	}

	//Jwt
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Jwt.AccessExpire
	accessSecret := l.svcCtx.Config.Jwt.AccessSecret
	token, err := l.getJwtToken(accessSecret,
		now,
		accessExpire,
		user.Id)

	return &types.LoginResp{
		Username:     user.Name,
		Age:          int(user.Age),
		Gender:       user.Gender,
		Token:        token,
		ExpireTime:   now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil

}

func (l *GetUserLogic) getJwtToken(key string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(key))
}
