// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchGetMeetingRoomBuilding 该接口用于获取指定建筑物的详细信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukzNyUjL5cjM14SO3ITN
func (r *MeetingRoomService) BatchGetMeetingRoomBuilding(ctx context.Context, request *BatchGetMeetingRoomBuildingReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomBuildingResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomBatchGetMeetingRoomBuilding != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#BatchGetMeetingRoomBuilding mock enable")
		return r.cli.mock.mockMeetingRoomBatchGetMeetingRoomBuilding(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "BatchGetMeetingRoomBuilding",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/building/batch_get?building_ids=omb_8ec170b937536a5d87c23b418b83f9bb&building_ids=omb_38570e4f0fd9ecf15030d3cc8b388f3a&fields=*",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetMeetingRoomBuildingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockMeetingRoomBatchGetMeetingRoomBuilding(f func(ctx context.Context, request *BatchGetMeetingRoomBuildingReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomBuildingResp, *Response, error)) {
	r.mockMeetingRoomBatchGetMeetingRoomBuilding = f
}

func (r *Mock) UnMockMeetingRoomBatchGetMeetingRoomBuilding() {
	r.mockMeetingRoomBatchGetMeetingRoomBuilding = nil
}

type BatchGetMeetingRoomBuildingReq struct {
	BuildingIDs string  `query:"building_ids" json:"-"` // 用于查询指定建筑物的 ID
	Fields      *string `query:"fields" json:"-"`       // 用于指定返回的字段名，每个字段名之间用逗号 "," 分隔，如：“id,name”，"*" 表示返回全部字段，可选字段有："id,name,description,floors"，默认返回所有字段
}

type batchGetMeetingRoomBuildingResp struct {
	Code int64                            `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetMeetingRoomBuildingResp `json:"data,omitempty"` // 返回业务信息
}

type BatchGetMeetingRoomBuildingResp struct {
	Buildings *BatchGetMeetingRoomBuildingRespBuildings `json:"buildings,omitempty"` // 建筑列表
}

type BatchGetMeetingRoomBuildingRespBuildings struct {
	BuildingID  string   `json:"building_id,omitempty"` // 建筑物 ID
	Description string   `json:"description,omitempty"` // 建筑物的相关描述
	Floors      []string `json:"floors,omitempty"`      // 属于当前建筑物的所有楼层列表
	Name        string   `json:"name,omitempty"`        // 建筑物名称
}