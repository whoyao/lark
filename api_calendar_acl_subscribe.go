// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// SubscribeCalendarACL 该接口用于以用户身份订阅指定日历下的日历成员变更事件。
//
// 用户必须对日历有访问权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/subscription
func (r *CalendarService) SubscribeCalendarACL(ctx context.Context, request *SubscribeCalendarACLReq, options ...MethodOptionFunc) (*SubscribeCalendarACLResp, *Response, error) {
	if r.cli.mock.mockCalendarSubscribeCalendarACL != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#SubscribeCalendarACL mock enable")
		return r.cli.mock.mockCalendarSubscribeCalendarACL(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Calendar",
		API:                 "SubscribeCalendarACL",
		Method:              "POST",
		URL:                 "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls/subscription",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(subscribeCalendarACLResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockCalendarSubscribeCalendarACL(f func(ctx context.Context, request *SubscribeCalendarACLReq, options ...MethodOptionFunc) (*SubscribeCalendarACLResp, *Response, error)) {
	r.mockCalendarSubscribeCalendarACL = f
}

func (r *Mock) UnMockCalendarSubscribeCalendarACL() {
	r.mockCalendarSubscribeCalendarACL = nil
}

type SubscribeCalendarACLReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
}

type subscribeCalendarACLResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *SubscribeCalendarACLResp `json:"data,omitempty"`
}

type SubscribeCalendarACLResp struct{}