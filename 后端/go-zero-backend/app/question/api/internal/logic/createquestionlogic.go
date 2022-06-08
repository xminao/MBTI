package logic

import (
	"backend/app/question/models"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"backend/app/question/api/internal/svc"
	"backend/app/question/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateQuestionLogic {
	return &CreateQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Option struct {
	OptionADesc   string `json:"option_a_desc"`
	OptionATarget string `json:"option_a_target"`
	OptionBDesc   string `json:"option_b_desc"`
	OptionBTarget string `json:"option_b_target"`
}

func (l *CreateQuestionLogic) CreateQuestion(req *types.CreateQuestionReq) (*types.CreateQuestionResp, error) {
	if len(strings.TrimSpace(req.QuestionDesc)) == 0 {
		return nil, errors.New("题目描述不能为空")
	}

	countBuilder := l.svcCtx.QuestionInfoModel.CountBuilder("*")
	count, err := l.svcCtx.QuestionInfoModel.FindCount(l.ctx, countBuilder)

	question := new(models.QuestionInfo)

	if count == 0 {
		question.QuestionId = 1
	} else {
		maxBuilder := l.svcCtx.QuestionInfoModel.MaxRowBuilder("created_at")
		LatestRecord, err := l.svcCtx.QuestionInfoModel.FindOneByQuery(l.ctx, maxBuilder)
		if err != nil {
			return nil, err
		}
		question.QuestionId = LatestRecord.QuestionId + 1
	}

	// generate json for insert

	option, err := json.Marshal(Option{
		OptionADesc:   req.OptionADesc,
		OptionATarget: req.OptionATarget,
		OptionBDesc:   req.OptionBDesc,
		OptionBTarget: req.OptionBTarget,
	})

	question.QuestionDesc = req.QuestionDesc
	question.QuestionOption = string(option)
	question.CreatedAt = time.Now()

	// insert question data
	_, err = l.svcCtx.QuestionInfoModel.Insert(l.ctx, question)
	if err != nil {
		return nil, err
	}

	return &types.CreateQuestionResp{
		Msg: "插入成功",
	}, nil
}
