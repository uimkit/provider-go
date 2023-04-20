package server

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	uim "github.com/uimkit/provider-go"
)

// UIM 需要的接口

// 新账号事件
type NewAccountHandler func(*cloudevents.Event, *uim.IMAccount) error

func (client *Client) OnNewAccount(handler NewAccountHandler) {
	client.OnEvent(uim.ProviderEventNewAccount, uim.CastEventHandler(handler))
}

// 账号更新
type AccountUpdatedHandler func(*cloudevents.Event, *uim.IMAccountUpdate) error

func (client *Client) OnAccountUpdated(handler AccountUpdatedHandler) {
	client.OnEvent(uim.ProviderEventAccountUpdated, uim.CastEventHandler(handler))
}

// 新好友
type NewContactHandler func(*cloudevents.Event, *uim.Contact) error

func (client *Client) OnNewContact(handler NewContactHandler) {
	client.OnEvent(uim.ProviderEventNewContact, uim.CastEventHandler(handler))
}

// 新粉丝
type NewFollowerHandler func(*cloudevents.Event, *uim.Follower) error

func (client *Client) OnNewFollower(handler NewFollowerHandler) {
	client.OnEvent(uim.ProviderEventNewFollower, uim.CastEventHandler(handler))
}

// 新关注的人
type NewFollowingHandler func(*cloudevents.Event, *uim.Following) error

func (client *Client) OnNewFollowing(handler NewFollowingHandler) {
	client.OnEvent(uim.ProviderEventNewFollowing, uim.CastEventHandler(handler))
}

// 新的好友申请
type NewFriendApplyHandler func(*cloudevents.Event, *uim.FriendApply) error

func (client *Client) OnNewFriendApply(handler NewFriendApplyHandler) {
	client.OnEvent(uim.ProviderEventNewFriendApply, uim.CastEventHandler(handler))
}

// 新消息
type NewMessageHandler func(*cloudevents.Event, *uim.Message) error

func (client *Client) OnNewMessage(handler NewMessageHandler) {
	client.OnEvent(uim.ProviderEventNewMessage, uim.CastEventHandler(handler))
}

// 消息更新
type MessageUpdatedHandler func(*cloudevents.Event, *uim.MessageUpdate) error

func (client *Client) OnMessageUpdated(handler MessageUpdatedHandler) {
	client.OnEvent(uim.ProviderEventMessageUpdated, uim.CastEventHandler(handler))
}

// 新群组
type NewGroupHandler func(*cloudevents.Event, *uim.Group) error

func (client *Client) OnNewGroup(handler NewGroupHandler) {
	client.OnEvent(uim.ProviderEventNewGroup, uim.CastEventHandler(handler))
}

// 群组更新
type GroupUpdatedHandler func(*cloudevents.Event, *uim.GroupUpdate) error

func (client *Client) OnGroupUpdated(handler GroupUpdatedHandler) {
	client.OnEvent(uim.ProviderEventGroupUpdated, uim.CastEventHandler(handler))
}

// 新群成员
type NewGroupMemberHandler func(*cloudevents.Event, *uim.GroupMember) error

func (client *Client) OnNewGroupMember(handler NewGroupMemberHandler) {
	client.OnEvent(uim.ProviderEventNewGroupMember, uim.CastEventHandler(handler))
}

// 群成员更新
type GroupMemberUpdatedHandler func(*cloudevents.Event, *uim.GroupMemberUpdate) error

func (client *Client) OnGroupMemberUpdated(handler GroupMemberUpdatedHandler) {
	client.OnEvent(uim.ProviderEventGroupMemberUpdated, uim.CastEventHandler(handler))
}

// 收到入群邀请
type NewGroupInvitationHandler func(*cloudevents.Event, *uim.GroupInvitation) error

func (client *Client) OnNewGroupInvitation(handler NewGroupInvitationHandler) {
	client.OnEvent(uim.ProviderEventNewGroupInvitation, uim.CastEventHandler(handler))
}

// 收到入群申请
type NewGroupApplyHandler func(*cloudevents.Event, *uim.GroupApply) error

func (client *Client) OnNewGroupApply(handler NewGroupApplyHandler) {
	client.OnEvent(uim.ProviderEventNewGroupApply, uim.CastEventHandler(handler))
}

// 新的元数据
type NewMetafieldHandler func(*cloudevents.Event, *uim.Metafield) error

func (client *Client) OnNewMetafield(handler NewMetafieldHandler) {
	client.OnEvent(uim.ProviderEventNewMetafield, uim.CastEventHandler(handler))
}

// 元数据更新
type MetafieldUpdatedHandler func(*cloudevents.Event, *uim.MetafieldUpdate) error

func (client *Client) OnMetafieldUpdated(handler MetafieldUpdatedHandler) {
	client.OnEvent(uim.ProviderEventMetafieldUpdated, uim.CastEventHandler(handler))
}

// 查询元数据
type GetMetafieldHandler func(*cloudevents.Event, *uim.GetMetafieldRequest) (*uim.GetMetafieldResponse, error)

func (client *Client) OnGetMetafield(handler GetMetafieldHandler) {
	client.OnEvent(uim.ProviderCommandGetMetafield, uim.CastCommandHandler(handler))
}

// 查询消息地址关联的信息
func (client *Client) GetChannelInfo(req *uim.GetChannelInfoRequest, opts ...uim.RequestOption) (*uim.GetChannelInfoResponse, error) {
	return uim.CastCommandResponse[*uim.GetChannelInfoResponse](
		client.Invoke(
			uim.UIMCommandGetChannelInfo,
			req,
			&uim.GetChannelInfoResponse{},
			opts...,
		),
	)
}

// 发送消息
func (client *Client) SendMessage(req *uim.SendMessageRequest, opts ...uim.RequestOption) (*uim.SendMessageResponse, error) {
	return uim.CastCommandResponse[*uim.SendMessageResponse](
		client.Invoke(
			uim.UIMCommandSendMessage,
			req,
			&uim.SendMessageResponse{},
			opts...,
		),
	)
}

// 发布朋友圈
func (client *Client) PublishMoment(req *uim.PublishMomentRequest, opts ...uim.RequestOption) (*uim.PublishMomentResponse, error) {
	return uim.CastCommandResponse[*uim.PublishMomentResponse](
		client.Invoke(
			uim.UIMCommandPublishMoment,
			req,
			&uim.PublishMomentResponse{},
			opts...,
		),
	)
}

// 获取动态列表
func (client *Client) GetMomentList(req *uim.GetMomentListRequest, opts ...uim.RequestOption) (*uim.GetMomentListResponse, error) {
	return uim.CastCommandResponse[*uim.GetMomentListResponse](
		client.Invoke(
			uim.UIMCommandGetMomentList,
			req,
			&uim.GetMomentListResponse{},
			opts...,
		),
	)
}

// 申请好友
func (client *Client) AddContact(req *uim.AddContactRequest, opts ...uim.RequestOption) (*uim.AddContactResponse, error) {
	return uim.CastCommandResponse[*uim.AddContactResponse](
		client.Invoke(
			uim.UIMCommandAddContact,
			req,
			&uim.AddContactResponse{},
			opts...,
		),
	)
}

// 设置群组禁言
func (client *Client) SetGroupMute(req *uim.SetGroupMuteRequest, opts ...uim.RequestOption) (*uim.SetGroupMuteResponse, error) {
	return uim.CastCommandResponse[*uim.SetGroupMuteResponse](
		client.Invoke(
			uim.UIMCommandSetGroupMute,
			req,
			&uim.SetGroupMuteResponse{},
			opts...,
		),
	)
}
