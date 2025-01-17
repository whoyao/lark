// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHelpdeskCategory 该接口用于获取知识库分类。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/get
func (r *HelpdeskService) GetHelpdeskCategory(ctx context.Context, request *GetHelpdeskCategoryReq, options ...MethodOptionFunc) (*GetHelpdeskCategoryResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskCategory != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskCategory mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskCategory(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskCategory",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/categories/:id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskCategoryResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskGetHelpdeskCategory(f func(ctx context.Context, request *GetHelpdeskCategoryReq, options ...MethodOptionFunc) (*GetHelpdeskCategoryResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskCategory = f
}

func (r *Mock) UnMockHelpdeskGetHelpdeskCategory() {
	r.mockHelpdeskGetHelpdeskCategory = nil
}

type GetHelpdeskCategoryReq struct {
	ID string `path:"id" json:"-"` // 知识库分类ID, 示例值："6948728206392295444"
}

type getHelpdeskCategoryResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskCategoryResp `json:"data,omitempty"`
}

type GetHelpdeskCategoryResp struct {
	CategoryID string `json:"category_id,omitempty"` // 知识库分类ID
	ID         string `json:"id,omitempty"`          // 知识库分类ID，（旧版，请使用category_id）
	Name       string `json:"name,omitempty"`        // 名称
	HelpdeskID string `json:"helpdesk_id,omitempty"` // 服务台ID
	Language   string `json:"language,omitempty"`    // 语言
}
