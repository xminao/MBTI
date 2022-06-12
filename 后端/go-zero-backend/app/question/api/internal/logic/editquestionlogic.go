package logic

import (
	"backend/app/question/api/internal/svc"
	"backend/app/question/api/internal/types"
	"backend/app/question/models"
	"context"
	"encoding/json"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditQuestionLogic {
	return &EditQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditQuestionLogic) EditQuestion(req *types.EditQuestionReq) (resp *types.EditQuestionResp, err error) {
	if l.ctx.Value("authGroup") != "admin" {
		return nil, errors.New("非管理员用户")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致
	ques, err := l.svcCtx.QuestionInfoModel.FindOne(l.ctx, int64(req.QuestionId))
	if err == models.ErrNotFound {
		return nil, errors.New("题目不存在")
	}

	option, err := json.Marshal(Option{
		OptionADesc:   req.OptionA,
		OptionATarget: req.TargetA,
		OptionBDesc:   req.OptionB,
		OptionBTarget: req.TargetB,
	})

	if err != nil {
		return nil, err
	}

	ques.QuestionDesc = req.QuestionDesc
	ques.QuestionOption = string(option)

	l.svcCtx.QuestionInfoModel.Update(l.ctx, ques)

	return &types.EditQuestionResp{
		Msg: "修改成功",
	}, nil
}
