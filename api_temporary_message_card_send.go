package lark

import "context"

// SendTemporaryMassageCard 用于机器人在群会话中发送指定用户可见的消息卡片。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要机器人在会话群里
// - 仅触发临时卡片的用户自己可见
// - 不支持转发
// - 只能在群聊使用
// - 仅在用户处于在线状态的飞书客户端上可见
// - 临时消息卡片的呈现能力、交互能力与消息卡片一致
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uETOyYjLxkjM24SM5IjN
func (r *MessageService) SendTemporaryMassageCard(ctx context.Context, request *SendTemporaryMassageCardReq, options ...MethodOptionFunc) (*SendTemporaryMassageCardResp, *Response, error) {

	if r.cli.mock.mockSendTemporaryMassageCard != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MessageCard#SendTemporaryMassageCard mock enable")
		return r.cli.mock.mockSendTemporaryMassageCard(ctx, request, options...)
	}

	request.MsgType = "interactive"

	req := &RawRequestReq{
		Scope:                 "MessageCard",
		API:                   "SendTemporaryMassageCard",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/ephemeral/v1/send",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(sendTemporaryMassageCardResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockSendTemporaryMassageCard(f func(ctx context.Context, request *SendTemporaryMassageCardReq, options ...MethodOptionFunc) (*SendTemporaryMassageCardResp, *Response, error)) {
	r.mockSendTemporaryMassageCard = f
}

func (r *Mock) UnMockSendTemporaryMassageCard() {
	r.mockSendTemporaryMassageCard = nil
}

type SendTemporaryMassageCardReq struct {
	ChatID  *string             `json:"chat_id"`           // 发送临时消息的群的chat_id
	OpenID  *string             `json:"open_id,omitempty"` // 指定发送临时消息卡片的用户的open_id，其他人将无法看到临时消息卡片
	UserID  *string             `json:"user_id,omitempty"` // 指定发送临时消息卡片的用户的user_id
	Email   *string             `json:"email,omitempty"`   // 指定发送临时消息卡片的用户的email, open_id、user_id、email 不能同时为空， 优先级 open_id > user_id > email
	MsgType string              `json:"msg_type"`          // 消息的类型，此处固定填 "interactive"
	Card    *MessageContentCard `json:"card"`              // 消息卡片的描述内容，具体参考 https://open.feishu.cn/document/ukTMukTMukTM/ugTNwUjL4UDM14CO1ATN
}

type sendTemporaryMassageCardResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 返回码描述
	Data *SendTemporaryMassageCardResp `json:"data,omitempty"`
}

type SendTemporaryMassageCardResp struct {
	MessageID *string `json:"message_id,omitempty"` // 消息id open_message_id
}
