// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHireAttachmentPreview 根据附件 ID 获取附件预览信息
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/attachment/preview
func (r *HireService) GetHireAttachmentPreview(ctx context.Context, request *GetHireAttachmentPreviewReq, options ...MethodOptionFunc) (*GetHireAttachmentPreviewResp, *Response, error) {
	if r.cli.mock.mockHireGetHireAttachmentPreview != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireAttachmentPreview mock enable")
		return r.cli.mock.mockHireGetHireAttachmentPreview(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireAttachmentPreview",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/hire/v1/attachments/:attachment_id/preview",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireAttachmentPreviewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHireGetHireAttachmentPreview(f func(ctx context.Context, request *GetHireAttachmentPreviewReq, options ...MethodOptionFunc) (*GetHireAttachmentPreviewResp, *Response, error)) {
	r.mockHireGetHireAttachmentPreview = f
}

func (r *Mock) UnMockHireGetHireAttachmentPreview() {
	r.mockHireGetHireAttachmentPreview = nil
}

type GetHireAttachmentPreviewReq struct {
	AttachmentID string `path:"attachment_id" json:"-"` // 附件id, 示例值："11111"
}

type getHireAttachmentPreviewResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *GetHireAttachmentPreviewResp `json:"data,omitempty"`
}

type GetHireAttachmentPreviewResp struct {
	URL string `json:"url,omitempty"` // 预览链接
}
