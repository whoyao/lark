// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetSheetProtectedDimension 该接口用于根据保护范围ID查询详细的保护行列信息，最多支持同时查询5个ID。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQTM5YjL0ETO24CNxkjN
func (r *DriveService) GetSheetProtectedDimension(ctx context.Context, request *GetSheetProtectedDimensionReq, options ...MethodOptionFunc) (*GetSheetProtectedDimensionResp, *Response, error) {
	if r.cli.mock.mockDriveGetSheetProtectedDimension != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSheetProtectedDimension mock enable")
		return r.cli.mock.mockDriveGetSheetProtectedDimension(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetSheetProtectedDimension",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/protected_range_batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getSheetProtectedDimensionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveGetSheetProtectedDimension(f func(ctx context.Context, request *GetSheetProtectedDimensionReq, options ...MethodOptionFunc) (*GetSheetProtectedDimensionResp, *Response, error)) {
	r.mockDriveGetSheetProtectedDimension = f
}

func (r *Mock) UnMockDriveGetSheetProtectedDimension() {
	r.mockDriveGetSheetProtectedDimension = nil
}

type GetSheetProtectedDimensionReq struct {
	ProtectIDs       []string `query:"protectIds" json:"-"`      // 保护范围ID，可以通过[获取表格元数据](https://open.feishu.cn/document/ukTMukTMukTM/uETMzUjLxEzM14SMxMTN)接口获取，多个ID用逗号分隔，如xxxID1,xxxID2
	MemberType       *string  `query:"memberType" json:"-"`      // 返回的用户类型，可选userId,openId,unionId,默认使用userId
	SpreadSheetToken string   `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token，获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
}

type getSheetProtectedDimensionResp struct {
	Code int64                           `json:"code,omitempty"`
	Msg  string                          `json:"msg,omitempty"`
	Data *GetSheetProtectedDimensionResp `json:"data,omitempty"`
}

type GetSheetProtectedDimensionResp struct {
	ProtectedRange *GetSheetProtectedDimensionRespProtectedRange `json:"protectedRange,omitempty"` // 保护范围
}

type GetSheetProtectedDimensionRespProtectedRange struct {
	ProtectID string                                                 `json:"protectId,omitempty"` // 保护范围ID
	Dimension *GetSheetProtectedDimensionRespProtectedRangeDimension `json:"dimension,omitempty"` // 保护范围，如果为空，则为保护子表
	SheetID   string                                                 `json:"sheetId,omitempty"`   // sheet的id
	LockInfo  string                                                 `json:"lockInfo,omitempty"`  // 保护说明
	Editors   *GetSheetProtectedDimensionRespProtectedRangeEditors   `json:"editors,omitempty"`   // 用户信息
}

type GetSheetProtectedDimensionRespProtectedRangeDimension struct {
	SheetID        string `json:"sheetId,omitempty"`        // sheet 的 id
	StartIndex     int64  `json:"startIndex,omitempty"`     // 保护行列起始下标，下标从1开始
	EndIndex       int64  `json:"endIndex,omitempty"`       // 保护行列终止下标，下标从1开始
	MajorDimension string `json:"majorDimension,omitempty"` // 保护范围的维度，COLUMNS为保护列，ROWS为保护行
}

type GetSheetProtectedDimensionRespProtectedRangeEditors struct {
	Users []*GetSheetProtectedDimensionRespProtectedRangeEditorsUser `json:"users,omitempty"` // 用户信息列表
}

type GetSheetProtectedDimensionRespProtectedRangeEditorsUser struct {
	MemberType string `json:"memberType,omitempty"` // 用户类型
	MemberID   string `json:"memberId,omitempty"`   // 用户ID
}
