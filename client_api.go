package uim

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

// 添加账号
func (client *Client) NewAccount(account *IMAccount) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventNewAccount)
	ce.SetData(cloudevents.ApplicationJSON, account)
	return client.SendEvent(&ce)
}

// 更新账号
func (client *Client) UpdateAccount(account *UpdateIMAccount) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventAccountUpdated)
	ce.SetData(cloudevents.ApplicationJSON, account)
	return client.SendEvent(&ce)
}

// 好友请求
func (client *Client) NewFriendApply(apply *FriendApply) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventNewFriendApply)
	ce.SetData(cloudevents.ApplicationJSON, apply)
	return client.SendEvent(&ce)
}

// 添加好友
func (client *Client) NewContact(contact *Contact) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventNewContact)
	ce.SetData(cloudevents.ApplicationJSON, contact)
	return client.SendEvent(&ce)
}

// 更新好友
func (client *Client) UpdateContact(contact *UpdateContact) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventContactUpdated)
	ce.SetData(cloudevents.ApplicationJSON, contact)
	return client.SendEvent(&ce)
}

// 新建群组
func (client *Client) NewGroup(group *Group) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventNewGroup)
	ce.SetData(cloudevents.ApplicationJSON, group)
	return client.SendEvent(&ce)
}

// 更新群组
func (client *Client) UpdateGroup(group *UpdateGroup) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventGroupUpdated)
	ce.SetData(cloudevents.ApplicationJSON, group)
	return client.SendEvent(&ce)
}

// 邀请入群
func (client *Client) NewGroupInvitation(invitation *GroupInvitation) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventNewGroupInvitation)
	ce.SetData(cloudevents.ApplicationJSON, invitation)
	return client.SendEvent(&ce)
}

// 申请入群
func (client *Client) NewJoinGroupApply(apply *JoinGroupApply) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventNewJoinGroupApply)
	ce.SetData(cloudevents.ApplicationJSON, apply)
	return client.SendEvent(&ce)
}

// 添加群组成员
func (client *Client) NewGroupMember(member *GroupMember) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventNewGroupMember)
	ce.SetData(cloudevents.ApplicationJSON, member)
	return client.SendEvent(&ce)
}

// 更新群组成员
func (client *Client) UpdateGroupMember(member *UpdateGroupMember) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventGroupMemberUpdated)
	ce.SetData(cloudevents.ApplicationJSON, member)
	return client.SendEvent(&ce)
}

// 新消息
func (client *Client) NewMessage(message *Message) error {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.providerEventSource)
	ce.SetType(ProviderEventNewMessage)
	ce.SetData(cloudevents.ApplicationJSON, message)
	return client.SendEvent(&ce)
}
