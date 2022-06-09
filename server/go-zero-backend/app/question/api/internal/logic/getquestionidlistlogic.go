package logic

import (
	"backend/app/question/api/internal/svc"
	"backend/app/question/api/internal/types"
	"backend/util/xerr"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQuestionIdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQuestionIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQuestionIdListLogic {
	return &GetQuestionIdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQuestionIdListLogic) GetQuestionIdList(req *types.GetQuestionIdListReq) (resp *types.GetQuestionIdListResp, err error) {
	username := l.ctx.Value("userName")
	if username == nil {
		return nil, xerr.NewErrCodeMsg(xerr.USER_ERROR, "无效的登录令牌")
	}
	logx.Infof("操作用户名：%v", l.ctx.Value("userName")) // 与传入一致
	build := l.svcCtx.QuestionInfoModel.RowBuilder()
	idList, err := l.svcCtx.QuestionInfoModel.GetIDList(l.ctx, build)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}
	return &types.GetQuestionIdListResp{
		QuestionIdList: idList,
	}, nil
}
