package logic

import (
	"context"
	"encoding/json"
	"errors"
	"question/api/internal/models"

	"question/api/internal/svc"
	"question/api/internal/types"

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

/*
 * 查询数据
 * json字符串的byte数组反序列化
 * 反序列化：
 */
func toJson(data []byte) map[string]interface{} {
	var ben map[string]interface{}
	err := json.Unmarshal(data, &ben)
	if err != nil {
		return nil
	}
	return ben
}

func (l *GetQuestionLogic) GetQuestion(req *types.QuestionReq) (*types.QuestionResp, error) {
	// todo: add your logic here and delete this line
	if req.QuestionId > 93 || req.QuestionId < 1 {
		return nil, errors.New("题目不存在")
	}

	question := models.QuestionBank{}
	l.svcCtx.DbEngin.Where("question_id").First(&question)

	return &types.QuestionResp{
		QuestionDesc: toJson(question.QuestionText.RawMessage)["question_desc"],
		OptionA:      toJson(question.QuestionText.RawMessage)["option_a"],
		OptionB:      toJson(question.QuestionText.RawMessage)["option_b"],
	}, nil
}
