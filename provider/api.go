package provider

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	uim "github.com/uimkit/provider-go"
)

// 新账号
func (client *Client) NewAccount(account *uim.IMAccount, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewAccount, account, opts...)
}

// 账号更新
func (client *Client) AccountUpdated(account *uim.IMAccountUpdate, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventAccountUpdated, account, opts...)
}

// 新好友
func (client *Client) NewContact(contact *uim.Contact, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewContact, contact, opts...)
}

// 新粉丝
func (client *Client) NewFollower(follower *uim.Follower, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewFollower, follower, opts...)
}

// 新关注的人
func (client *Client) NewFollwing(following *uim.Following, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewFollowing, following, opts...)
}

// 新的好友申请
func (client *Client) NewFriendApply(apply *uim.FriendApply, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewFriendApply, apply, opts...)
}

// 新消息
func (client *Client) NewMessage(message *uim.Message, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewMessage, message, opts...)
}

// 消息更新
func (client *Client) MessageUpdated(message *uim.MessageUpdate, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventMessageUpdated, message, opts...)
}

// 新群组
func (client *Client) NewGroup(group *uim.Group, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewGroup, group, opts...)
}

// 群组更新
func (client *Client) GroupUpdated(group *uim.GroupUpdate, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventGroupUpdated, group, opts...)
}

// 新群成员
func (client *Client) NewGroupMember(member *uim.GroupMember, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewGroupMember, member, opts...)
}

// 群成员更新
func (client *Client) GroupMemberUpdated(member *uim.GroupMemberUpdate, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventGroupMemberUpdated, member, opts...)
}

// 群成员删除
func (client *Client) GroupMemberDeleted(member *uim.GroupMemberDelete, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventGroupMemberDeleted, member, opts...)
}

// 收到入群邀请
func (client *Client) NewGroupInvitation(invitation *uim.GroupInvitation, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewGroupInvitation, invitation, opts...)
}

// 收到入群申请
func (client *Client) NewJoinGroupApply(apply *uim.JoinGroupApply, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewJoinGroupApply, apply, opts...)
}

// 新的元数据
func (client *Client) NewMetafield(metafield *uim.Metafield, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventNewMetafield, metafield, opts...)
}

// 元数据更新
func (client *Client) MetafieldUpdated(metafield *uim.MetafieldUpdate, opts ...uim.RequestOption) error {
	return client.SendEvent(uim.ProviderEventMetafieldUpdated, metafield, opts...)
}

// 查询元数据
func (client *Client) GetMetafield(metafield *uim.GetMetafieldRequest, opts ...uim.RequestOption) (*uim.GetMetafieldResponse, error) {
	return uim.CastCommandResponse[*uim.GetMetafieldResponse](
		client.Invoke(
			uim.ProviderCommandGetMetafield,
			metafield,
			&uim.GetMetafieldResponse{},
			opts...,
		),
	)
}

// 发送消息
type SendMessageHandler func(*cloudevents.Event, *uim.SendMessageRequest) (*uim.SendMessageResponse, error)

func (client *Client) OnSendMessage(handler SendMessageHandler) {
	client.OnEvent(uim.UIMCommandSendMessage, uim.CastCommandHandler(handler))
}

// 添加好友
type AddContactHandler func(*cloudevents.Event, *uim.AddContactRequest) (*uim.AddContactResponse, error)

func (client *Client) OnAddContact(handler AddContactHandler) {
	client.OnEvent(uim.UIMCommandAddContact, uim.CastCommandHandler(handler))
}
