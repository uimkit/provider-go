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
	UIMEventListAccounts          = "uim.list_accounts"            // 查询账号列表
	UIMEventUpdateUser            = "uim.update_user"              // 更新账号用户资料
	UIMEventUpdateContact         = "uim.update_contact"           // 更新联系人
	UIMEventListContacts          = "uim.list_contacts"            // 查询联系人列表
	UIMEventApplyFriend           = "uim.apply_friend"             // 添加好友
	UIMEventApproveFriendApply    = "uim.approve_friend_apply"     // 通过好友请求
	UIMEventNewGroup              = "uim.new_group"                // 创建群组
	UIMEventListGroups            = "uim.list_groups"              // 查询群组列表
	UIMEventUpdateGroup           = "uim.update_group"             // 更新群组
	UIMEventInviteToGroup         = "uim.invite_to_group"          // 邀请加入群组
	UIMEventAcceptGroupInvitation = "uim.accept_group_invitation"  // 接受入群邀请
	UIMEventApplyJoinGroup        = "uim.apply_join_group"         // 申请加入群组
	UIMEventApproveJoinGroupApply = "uim.approve_join_group_apply" // 通过入群申请
	UIMEventListGroupMembers      = "uim.list_group_members"       // 查询群成员列表
	UIMEventPublishMoment         = "uim.publish_moment"           // 发布动态
)

// Provider传递给UIM的事件
const (
	ProviderEventNewAccount         = "provider.new_account"          // 添加账号
	ProviderEventAccountList        = "provider.account_list"         // 账号列表
	ProviderEventAccountUpdated     = "provider.account_updated"      // 账号更新
	ProviderEventNewFriendApply     = "provider.new_friend_apply"     // 收到好友请求
	ProviderEventNewContact         = "provider.new_contact"          // 添加好友
	ProviderEventContactList        = "provider.contact_list"         // 好友列表
	ProviderEventContactUpdated     = "provider.contact_updated"      // 好友更新
	ProviderEventNewGroup           = "provider.new_group"            // 添加群组
	ProviderEventGroupUpdated       = "provider.group_updated"        // 群组更新
	ProviderEventGroupList          = "provider.group_list"           // 群组列表
	ProviderEventNewGroupInvitation = "provider.new_group_invitation" // 收到入群邀请
	ProviderEventNewJoinGroupApply  = "provider.new_join_group_apply" // 收到入群申请
	ProviderEventNewGroupMember     = "provider.new_group_member"     // 添加群组成员
	ProviderEventGroupMemberUpdated = "provider.group_member_updated" // 群成员更新
	ProviderEventGroupMemberList    = "provider.group_member_list"    // 群成员列表
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

func (c *Client) triggerListAccounts(query *ListIMAccounts) error {
	for _, handler := range c.listAccountsHandlers {
		if err := handler(query); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerUpdateUser(user *UpdateIMUser) error {
	for _, handler := range c.updateUserHandlers {
		if err := handler(user); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerUpdateContact(contact *UpdateContact) error {
	for _, handler := range c.updateContactHandlers {
		if err := handler(contact); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerListContacts(query *ListContacts) error {
	for _, handler := range c.listContactsHandlers {
		if err := handler(query); err != nil {
			return err
		}
	}
	return nil
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

func (c *Client) triggerUpdateGroup(group *UpdateGroup) error {
	for _, handler := range c.updateGroupHandlers {
		if err := handler(group); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) triggerListGroups(query *ListGroups) error {
	for _, handler := range c.listGroupsHandlers {
		if err := handler(query); err != nil {
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

func (c *Client) triggerListGroupMembers(query *ListGroupMembers) error {
	for _, handler := range c.listGroupMembersHandlers {
		if err := handler(query); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) OnListAccounts(handler ListAccountsHandler) {
	c.listAccountsHandlers = append(c.listAccountsHandlers, handler)
}

func (c *Client) OnUpdateUser(handler UpdateUserHandler) {
	c.updateUserHandlers = append(c.updateUserHandlers, handler)
}

func (c *Client) OnUpdateContact(handler UpdateContactHandler) {
	c.updateContactHandlers = append(c.updateContactHandlers, handler)
}

func (c *Client) OnListContacts(handler ListContactsHandler) {
	c.listContactsHandlers = append(c.listContactsHandlers, handler)
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

func (c *Client) OnUpdateGroup(handler UpdateGroupHandler) {
	c.updateGroupHandlers = append(c.updateGroupHandlers, handler)
}

func (c *Client) OnListGroups(handler ListGroupsHandler) {
	c.listGroupsHandlers = append(c.listGroupsHandlers, handler)
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

func (c *Client) OnListGroupMembers(handler ListGroupMembersHandler) {
	c.listGroupMembersHandlers = append(c.listGroupMembersHandlers, handler)
}

func (c *Client) processEvent(event *cloudevents.Event) error {
	switch event.Type() {
	case UIMEventSendMessage:
		var message *SendMessage
		if err := event.DataAs(&message); err != nil {
			return err
		}
		return c.triggerSendMessage(message)
	case UIMEventListAccounts:
		var query *ListIMAccounts
		if err := event.DataAs(&query); err != nil {
			return err
		}
		return c.triggerListAccounts(query)
	case UIMEventUpdateUser:
		var user *UpdateIMUser
		if err := event.DataAs(&user); err != nil {
			return err
		}
		return c.triggerUpdateUser(user)
	case UIMEventUpdateContact:
		var contact *UpdateContact
		if err := event.DataAs(&contact); err != nil {
			return err
		}
		return c.triggerUpdateContact(contact)
	case UIMEventListContacts:
		var query *ListContacts
		if err := event.DataAs(&query); err != nil {
			return err
		}
		return c.triggerListContacts(query)
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
	case UIMEventUpdateGroup:
		var group *UpdateGroup
		if err := event.DataAs(&group); err != nil {
			return err
		}
		return c.triggerUpdateGroup(group)
	case UIMEventApplyJoinGroup:
		var apply *NewJoinGroupApply
		if err := event.DataAs(&apply); err != nil {
			return err
		}
		return c.triggerApplyJoinGroup(apply)
	case UIMEventListGroups:
		var query *ListGroups
		if err := event.DataAs(&query); err != nil {
			return err
		}
		return c.triggerListGroups(query)
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
	case UIMEventListGroupMembers:
		var query *ListGroupMembers
		if err := event.DataAs(&query); err != nil {
			return err
		}
		return c.triggerListGroupMembers(query)
	default:
	}
	return nil
}
