package uim

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

const (
	ProviderEventSource = "provider.source/%s/%s" // Provider事件源
	UIMEventSource      = "uim.source"            // UIM事件源
)

// UIM传递给Provider的事件
const (
	UIMEventSendMessage           = "uim.send_message"             // 发送消息
	UIMEventApplyFriend           = "uim.apply_friend"             // 添加好友
	UIMEventApproveFriendApply    = "uim.approve_friend_apply"     // 通过好友请求
	UIMEventNewGroup              = "uim.new_group"                // 创建群组
	UIMEventInviteToGroup         = "uim.invite_to_group"          // 邀请加入群组
	UIMEventAcceptGroupInvitation = "uim.accept_group_invitation"  // 接受入群邀请
	UIMEventApplyJoinGroup        = "uim.apply_join_group"         // 申请加入群组
	UIMEventApproveJoinGroupApply = "uim.approve_join_group_apply" // 通过入群申请
	UIMEventPublishMoment         = "uim.publish_moment"           // 发布动态
)

// Provider传递给UIM的事件
const (
	ProviderEventNewAccount         = "provider.new_account"          // 添加账号
	ProviderEventAccountUpdated     = "provider.account_updated"      // 账号更新
	ProviderEventNewFriendApply     = "provider.new_friend_apply"     // 收到好友请求
	ProviderEventNewContact         = "provider.new_contact"          // 添加好友
	ProviderEventContactUpdated     = "provider.contact_updated"      // 好友更新
	ProviderEventNewGroup           = "provider.new_group"            // 添加群组
	ProviderEventGroupUpdated       = "provider.group_updated"        // 群组更新
	ProviderEventNewGroupInvitation = "provider.new_group_invitation" // 收到入群邀请
	ProviderEventNewJoinGroupApply  = "provider.new_join_group_apply" // 收到入群申请
	ProviderEventNewGroupMember     = "provider.new_group_member"     // 添加群组成员
	ProviderEventGroupMemberUpdated = "provider.group_member_updated" // 群成员更新
	ProviderEventNewMessage         = "provider.new_message"          // 收到新消息
	ProviderEventNewMoment          = "provider.new_moment"           // 新发布动态
)

func (client *Client) SendEvent(event *cloudevents.Event) (err error) {
	content, _ := json.Marshal(event)
	req := NewBaseRequestWithPath("/")
	req.SetContent(content)
	resp := &BaseResponse{}
	return client.DoAction(req, resp)
}

func (c *Client) WebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		event, err := c.Webhook(r.Header, body)
		if err != nil {
			fmt.Println("Webhook is invalid :(")
			return
		}
		err = c.processEvent(event)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (c *Client) Webhook(header http.Header, body []byte) (*cloudevents.Event, error) {
	for _, token := range header[http.CanonicalHeaderKey("X-UIM-Key")] {
		if token == c.AppId && checkSignature(header.Get("X-UIM-Signature"), c.Secret, body) {
			var event *cloudevents.Event
			err := json.Unmarshal(body, event)
			if err != nil {
				return nil, err
			}
			return event, nil
		}
	}
	return nil, errors.New("invalid webhook")
}

func (c *Client) triggerSendMessage(message *SendMessage) error {
	for _, handler := range c.sendMessageHandlers {
		if err := handler(message); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerApplyFriend(apply *NewFriendApply) error {
	for _, handler := range c.applyFriendHandlers {
		if err := handler(apply); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerApproveFriendApply(apply *ApproveFriendApply) error {
	for _, handler := range c.approveFriendApplyHandlers {
		if err := handler(apply); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerNewGroup(group *NewGroup) error {
	for _, handler := range c.newGroupHandlers {
		if err := handler(group); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerApplyJoinGroup(apply *NewJoinGroupApply) error {
	for _, handler := range c.applyJoinGroupHandlers {
		if err := handler(apply); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerApproveJoinGroupApply(apply *ApproveJoinGroupApply) error {
	for _, handler := range c.approveJoinGroupApplyHandlers {
		if err := handler(apply); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerInviteToGroup(invite *InviteToGroup) error {
	for _, handler := range c.inviteToGroupHandlers {
		if err := handler(invite); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerAcceptGroupInvitation(invite *AcceptGroupInvitation) error {
	for _, handler := range c.acceptGroupInvitationHandlers {
		if err := handler(invite); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) OnSendMessage(handler SendMessageHandler) {
	c.sendMessageHandlers = append(c.sendMessageHandlers, handler)
}

func (c *Client) OnApplyFriend(handler ApplyFriendHandler) {
	c.applyFriendHandlers = append(c.applyFriendHandlers, handler)
}

func (c *Client) OnApproveFriendApply(handler ApproveFriendApplyHandler) {
	c.approveFriendApplyHandlers = append(c.approveFriendApplyHandlers, handler)
}

func (c *Client) OnNewGroup(handler NewGroupHandler) {
	c.newGroupHandlers = append(c.newGroupHandlers, handler)
}

func (c *Client) OnApplyJoinGroup(handler ApplyJoinGroupHandler) {
	c.applyJoinGroupHandlers = append(c.applyJoinGroupHandlers, handler)
}

func (c *Client) OnApproveJoinGroupApply(handler ApproveJoinGroupApplyHandler) {
	c.approveJoinGroupApplyHandlers = append(c.approveJoinGroupApplyHandlers, handler)
}

func (c *Client) OnInviteToGroup(handler InviteToGroupHandler) {
	c.inviteToGroupHandlers = append(c.inviteToGroupHandlers, handler)
}

func (c *Client) OnAcceptGroupInvitation(handler AcceptGroupInvitationHandler) {
	c.acceptGroupInvitationHandlers = append(c.acceptGroupInvitationHandlers, handler)
}

func (c *Client) processEvent(event *cloudevents.Event) error {
	switch event.Type() {
	case UIMEventSendMessage:
		var message *SendMessage
		if err := event.DataAs(&message); err != nil {
			return err
		}
		return c.triggerSendMessage(message)
	case UIMEventApplyFriend:
		var apply *NewFriendApply
		if err := event.DataAs(&apply); err != nil {
			return err
		}
		return c.triggerApplyFriend(apply)
	case UIMEventApproveFriendApply:
		var apply *ApproveFriendApply
		if err := event.DataAs(&apply); err != nil {
			return err
		}
		return c.triggerApproveFriendApply(apply)
	case UIMEventNewGroup:
		var group *NewGroup
		if err := event.DataAs(&group); err != nil {
			return err
		}
		return c.triggerNewGroup(group)
	case UIMEventApplyJoinGroup:
		var apply *NewJoinGroupApply
		if err := event.DataAs(&apply); err != nil {
			return err
		}
		return c.triggerApplyJoinGroup(apply)
	case UIMEventApproveJoinGroupApply:
		var apply *ApproveJoinGroupApply
		if err := event.DataAs(&apply); err != nil {
			return err
		}
		return c.triggerApproveJoinGroupApply(apply)
	case UIMEventInviteToGroup:
		var invite *InviteToGroup
		if err := event.DataAs(&invite); err != nil {
			return err
		}
		return c.triggerInviteToGroup(invite)
	case UIMEventAcceptGroupInvitation:
		var invite *AcceptGroupInvitation
		if err := event.DataAs(&invite); err != nil {
			return err
		}
		return c.triggerAcceptGroupInvitation(invite)
	default:
	}
	return nil
}
