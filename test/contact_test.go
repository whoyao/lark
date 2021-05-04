package test

import (
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Contact(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		contactCli := cli.Contact()

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.CreateDepartment(ctx, &lark.CreateDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.DeleteDepartment(ctx, &lark.DeleteDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetDepartment(ctx, &lark.GetDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetDepartmentList(ctx, &lark.GetDepartmentListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetParentDepartment(ctx, &lark.GetParentDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.SearchDepartment(ctx, &lark.SearchDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.UpdateDepartment(ctx, &lark.UpdateDepartmentReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.UpdateDepartmentPatch(ctx, &lark.UpdateDepartmentPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.CreateUser(ctx, &lark.CreateUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.DeleteUser(ctx, &lark.DeleteUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetUser(ctx, &lark.GetUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.GetUserList(ctx, &lark.GetUserListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.UpdateUser(ctx, &lark.UpdateUserReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := contactCli.UpdateUserPatch(ctx, &lark.UpdateUserPatchReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("No permission", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().Message().SendRawMessage(ctx, &lark.SendRawMessageReq{
			ReceiveIDType: lark.IDTypePtr(lark.IDTypeChatID),
			ReceiveID:     ptrString("x"),
			Content:       `{"text":"hi"}`,
			MsgType:       lark.MsgTypeText,
		})
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("No permission", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().Message().GetMessage(ctx, &lark.GetMessageReq{
			MessageID: "1",
		})
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("No permission", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().Message().GetMessageFile(ctx, &lark.GetMessageFileReq{
			Type:      "s",
			MessageID: "s",
			FileKey:   "s",
		})
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("No permission", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().Message().GetMessageList(ctx, &lark.GetMessageListReq{})
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("ids not existed", func(t *testing.T) {
		_, _, err := AppALLPermission.Ins().Message().GetMessage(ctx, &lark.GetMessageReq{
			MessageID: "1",
		})
		as.NotNil(err)
		as.Contains(err.Error(), "these ids not existed")
	})

	t.Run("send-raw-message", func(t *testing.T) {
		_, _, err := AppALLPermission.Ins().Message().SendRawMessage(ctx, &lark.SendRawMessageReq{
			ReceiveIDType: lark.IDTypePtr(lark.IDTypeChatID),
			ReceiveID:     ptrString(ChatForSendMessage.ChatID),
			Content:       fmt.Sprintf(`{"text":"%d"}`, time.Now().Unix()),
			MsgType:       lark.MsgTypeText,
		})
		as.Nil(err)
	})

	t.Run("get-message-read", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Message().GetMessageReadUser(ctx, &lark.GetMessageReadUserReq{
			UserIDType: lark.IDTypeUserID,
			MessageID:  MessageAdminSendTextInChatContainAllPermissionApp.MessageID,
		})
		spew.Dump(resp, err)
		as.NotNil(err)
		as.Contains(err.Error(), "Bot is NOT the sender of the message")
	})

	t.Run("get-message-read", func(t *testing.T) {
		// resp, _, err := AppALLPermission.Ins().Message().GetMessageRead(ctx, &lark.GetMessageReadReq{
		// 	UserIDType: lark.IDTypeUserID,
		// 	MessageID: MessageAdminSendTextInChatContainAllPermissionApp.MessageID,
		// })
		// spew.Dump(resp, err)
		// as.NotNil(err)
		// as.Contains(err.Error(), "Bot is NOT the sender of the message")
	})

	t.Run("get-message-text", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Message().GetMessage(ctx, &lark.GetMessageReq{
			MessageID: MessageAdminSendTextInChatContainAllPermissionApp.MessageID,
		})
		spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.Len(resp.Items, 1)
		as.Equal(lark.MsgTypeText, resp.Items[0].MsgType)
		as.Equal(MessageAdminSendTextInChatContainAllPermissionApp.ChatID, resp.Items[0].ChatID)
		msgContent, err := lark.UnwrapMessageContent(resp.Items[0])
		as.Nil(err)
		as.Equal("test", msgContent.Text)
	})

	t.Run("get-message-image", func(t *testing.T) {
		messageFile := ""
		{
			resp, _, err := AppALLPermission.Ins().Message().GetMessage(ctx, &lark.GetMessageReq{
				MessageID: MessageAdminSendImageInChatContainAllPermissionApp.MessageID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
			as.NotNil(resp)
			as.Len(resp.Items, 1)
			as.Equal(lark.MsgTypeImage, resp.Items[0].MsgType)
			as.Equal(MessageAdminSendImageInChatContainAllPermissionApp.ChatID, resp.Items[0].ChatID)
			as.Contains(resp.Items[0].Body.Content, "image_key")
			msgContent, err := lark.UnwrapMessageContent(resp.Items[0])
			as.Nil(err)
			messageFile = msgContent.ImageKey
		}

		{
			resp, _, err := AppALLPermission.Ins().Message().GetMessageFile(ctx, &lark.GetMessageFileReq{
				Type:      "image",
				MessageID: MessageAdminSendImageInChatContainAllPermissionApp.MessageID,
				FileKey:   messageFile,
			})
			as.Nil(err)
			as.NotNil(resp)
			bs, err := io.ReadAll(resp.File)
			as.Nil(err)
			as.NotEmpty(bs)
		}
	})

	t.Run("get-message-list", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Message().GetMessageList(ctx, &lark.GetMessageListReq{
			ContainerIDType: lark.ContainerIDTypeChat,
			ContainerID:     ChatContainALLPermissionApp.ChatID,
			StartTime:       nil,
			EndTime:         nil,
			PageToken:       nil,
			PageSize:        nil,
		})
		spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.True(len(resp.Items) > 0)
	})
}
