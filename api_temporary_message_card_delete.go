package lark

import "context"

// DeleteTemporaryMassageCard 在群会话中删除指定用户的临时消息卡片。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要机器人在会话群里
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uITOyYjLykjM24iM5IjN
func (r *MessageService) DeleteTemporaryMassageCard(ctx context.Context, request *DeleteTemporaryMassageCardReq, options ...MethodOptionFunc) (*Response, error) {

	if r.cli.mock.mockDeleteTemporaryMassageCard != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MessageCard#DeleteTemporaryMassageCard mock enable")
		return r.cli.mock.mockDeleteTemporaryMassageCard(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MessageCard",
		API:                   "SendTemporaryMassageCard",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/ephemeral/v1/delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteTemporaryMassageCardResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return response, err
}

func (r *Mock) MockDeleteTemporaryMassageCard(f func(ctx context.Context, request *DeleteTemporaryMassageCardReq, options ...MethodOptionFunc) (*Response, error)) {
	r.mockDeleteTemporaryMassageCard = f
}

func (r *Mock) UnMockDeleteTemporaryMassageCard() {
	r.mockDeleteTemporaryMassageCard = nil
}

type DeleteTemporaryMassageCardReq struct {
	MessageID *string `json:"message_id,omitempty"` // 消息id open_message_id
}

type deleteTemporaryMassageCardResp struct {
	Code int64  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string `json:"msg,omitempty"`  // 错误描述
}
