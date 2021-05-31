// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHireJobManager 根据职位 ID 获取职位上的招聘人员信息，如招聘负责人、用人经理
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job-manager/get
func (r *HireService) GetHireJobManager(ctx context.Context, request *GetHireJobManagerReq, options ...MethodOptionFunc) (*GetHireJobManagerResp, *Response, error) {
	if r.cli.mock.mockHireGetHireJobManager != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireJobManager mock enable")
		return r.cli.mock.mockHireGetHireJobManager(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireJobManager",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/hire/v1/jobs/:job_id/managers/:manager_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireJobManagerResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHireGetHireJobManager(f func(ctx context.Context, request *GetHireJobManagerReq, options ...MethodOptionFunc) (*GetHireJobManagerResp, *Response, error)) {
	r.mockHireGetHireJobManager = f
}

func (r *Mock) UnMockHireGetHireJobManager() {
	r.mockHireGetHireJobManager = nil
}

type GetHireJobManagerReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	JobID      string  `path:"job_id" json:"-"`        // 职位ID, 示例值："1618209327096"
	ManagerID  string  `path:"manager_id" json:"-"`    // 人员ID，目前传职位ID, 示例值："1618209327096"
}

type getHireJobManagerResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetHireJobManagerResp `json:"data,omitempty"`
}

type GetHireJobManagerResp struct {
	Info *GetHireJobManagerRespInfo `json:"info,omitempty"` // 职位负责人
}

type GetHireJobManagerRespInfo struct {
	ID                  string   `json:"id,omitempty"`                     // 职位ID
	RecruiterID         string   `json:"recruiter_id,omitempty"`           // 招聘负责人ID
	HiringManagerIDList []string `json:"hiring_manager_id_list,omitempty"` // 用人经理ID列表
	AssistantIDList     []string `json:"assistant_id_list,omitempty"`      // 协助人ID列表
}