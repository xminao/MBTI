package logic

import (
	"context"
	"errors"

	"backend/app/question_goctl/api/internal/svc"
	"backend/app/question_goctl/api/internal/types"
	"backend/app/question_goctl/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteQuestionLogic {
	return &DeleteQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteQuestionLogic) DeleteQuestion(req *types.DeleteQuestionReq) (*types.DeleteQuestionResp, error) {

	_, err := l.svcCtx.QuestionInfoModel.FindOne(l.ctx, int64(req.QuestionId))
	if err == models.ErrNotFound {
		return nil, errors.New("题目不存在")
	}

	err = l.svcCtx.QuestionInfoModel.Delete(l.ctx, req.QuestionId)
	if err != nil {
		return nil, err
	}

	return &types.DeleteQuestionResp{
		Msg: "删除成功",
	}, nil
}
