// Code generated by goctl. DO NOT EDIT.
package types

type Question struct {
	QuestionId   int64  `json:"question_id"`
	QuestionDesc string `json:"question_desc"`
	OptionA      string `json:"option_a_desc"`
	OptionB      string `json:"option_b_desc"`
}

type GetQuestionReq struct {
	QuestionId int `form:"id"`
}

type GetQuestionResp struct {
	QuestionInfo Question `json:"question_info"`
}

type CreateQuestionReq struct {
	QuestionDesc  string `json:"question_desc"`
	OptionADesc   string `json:"option_a_desc"`
	OptionATarget string `json:"option_a_target"`
	OptionBDesc   string `json:"option_b_desc"`
	OptionBTarget string `json:"option_b_target"`
}

type CreateQuestionResp struct {
	Msg string `json:"msg"`
}

type DeleteQuestionReq struct {
	QuestionId int64 `form:"id"`
}

type DeleteQuestionResp struct {
	Msg string `json:"msg"`
}
