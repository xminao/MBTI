package logic

import (
	"backend/app/user/models"
	"backend/util/xerr"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"strings"
	"time"

	"backend/app/user/api/internal/svc"
	"backend/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	if len(strings.TrimSpace(req.Username)) == 0 {
		return nil, errors.New("用户名不能为空")
	}
	if len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errors.New("密码不能为空")
	}

	//type models.UserInfo
	user, err := l.svcCtx.UserInfoModel.FindOne(l.ctx, req.Username)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.USER_ERROR)
	}

	if err == models.ErrNotFound {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "用户不存在哟")
	}
	if user.Password != req.Password {
		l.svcCtx.UserInfoModel.DeleteCache(l.ctx, req.Username)
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "密码错啦")
	}

	// Jwt
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire //get expire
	// call getJwtToken func to create jwt token
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, user.Username, user.AuthGroup)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR)
	}
	fmt.Println(user.AuthGroup)
	//return msg and token_info
	return &types.LoginResp{
		Msg:       "登录成功",
		Username:  user.Username,
		Nickname:  user.Nickname,
		AuthGroup: user.AuthGroup,
		Jwt: types.JwtToken{
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}

// get jwt token
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, userName, authGroup string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userName"] = userName
	claims["authGroup"] = authGroup
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
