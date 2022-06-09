package logic

import (
	"backend/util/xerr"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"backend/app/question/api/internal/svc"
	"backend/app/question/api/internal/types"

	"backend/app/question/models"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionLogic {
	return &GetQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionLogic) GetQuestion(req *types.GetQuestionReq) (*types.GetQuestionResp, error) {
	username := l.ctx.Value("userName")
	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致
	question, err := l.svcCtx.QuestionInfoModel.FindOne(l.ctx, int64(req.QuestionId))
	if err == models.ErrNotFound {
		return nil, errors.New("题目不存在")
	}
	// 存储的是json类型
	fmt.Println(question.QuestionOption)
	option := toJson([]byte(question.QuestionOption))
	fmt.Println(option)

	data := types.Question{
		QuestionId:   int64(req.QuestionId),
		QuestionDesc: question.QuestionDesc,
		OptionA:      option["option_a_desc"],
		OptionB:      option["option_b_desc"],
		TargetA:      option["option_a_target"],
		TargetB:      option["option_b_target"],
	}

	return &types.GetQuestionResp{
		QuestionInfo: data,
	}, nil
}

func toJson(data []byte) map[string]string {
	var ben map[string]string
	json.Unmarshal(data, &ben)
	return ben
}
