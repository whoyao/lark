package lark

import "context"

// SendMassageCard 给指定用户或者会话发送消息卡片，其中会话包括私聊会话和群会话。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 私聊会话时机器人需要拥有对用户的可见性
// - 群会话需要机器人在群里
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN
func (r *MessageService) SendMassageCard(ctx context.Context, request *SendMassageCardReq, options ...MethodOptionFunc) (*SendMassageCardResp, *Response, error) {

	if r.cli.mock.mockSendMassageCard != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MessageCard#mockSendMassageCard mock enable")
		return r.cli.mock.mockSendMassageCard(ctx, request, options...)
	}

	request.MsgType = "interactive"

	req := &RawRequestReq{
		Scope:                 "MessageCard",
		API:                   "SendMassageCard",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/message/v4/send/",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(sendMessageCardResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockSendMassageCard(f func(ctx context.Context, request *SendMassageCardReq, options ...MethodOptionFunc) (*SendMassageCardResp, *Response, error)) {
	r.mockSendMassageCard = f
}

func (r *Mock) UnMockSendMassageCard() {
	r.mockSendMassageCard = nil
}

type SendMassageCardReq struct {
	ChatID      *string             `json:"chat_id,omitempty"` // 接收卡片消息的群的chat_id
	OpenID      *string             `json:"open_id,omitempty"` // 接收卡片消息的人的open_id
	UserID      *string             `json:"user_id,omitempty"` // 接收卡片消息的人的user_id
	Email       *string             `json:"email,omitempty"`   // 接收卡片消息的人的email, chat_id、open_id、user_id、email 不能同时为空， 优先级 chat_id > open_id > user_id > email
	MsgType     string              `json:"msg_type"`          // 消息的类型，此处固定填 "interactive"
	Card        *MessageContentCard `json:"card"`              // 消息卡片的描述内容，具体参考 https://open.feishu.cn/document/ukTMukTMukTM/ugTNwUjL4UDM14CO1ATN
	RootID      string              `json:"root_id,omitempty"` // 需要回复的消息的open_message_id
	UpdateMulti bool                `json:"update_multi"`      // 控制卡片是否是共享卡片(所有用户共享同一张消息卡片），默认为 false，流程参考 https://open.feishu.cn/document/ukTMukTMukTM/uYjNwUjL2YDM14iN2ATN
}

type sendMessageCardResp struct {
	Code int64                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 返回码描述
	Data *SendMassageCardResp `json:"data,omitempty"`
}

type SendMassageCardResp struct {
	MessageID *string `json:"message_id,omitempty"` // 消息id open_message_id
}
