package server

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	uim "github.com/uimkit/provider-go"
)

// UIM 需要的接口

// 新账号事件
type NewAccountHandler func(*cloudevents.Event, *uim.NewIMAccount) error

func (client *Client) OnNewAccount(handler NewAccountHandler) {
	client.OnEvent(uim.ProviderEventNewAccount, uim.CastEventHandler(handler))
}

// 账号更新
type AccountUpdatedHandler func(*cloudevents.Event, *uim.IMAccountUpdate) error

func (client *Client) OnAccountUpdated(handler AccountUpdatedHandler) {
	client.OnEvent(uim.ProviderEventAccountUpdated, uim.CastEventHandler(handler))
}

// 新的好友申请
type NewFriendApplyHandler func(*cloudevents.Event, *uim.FriendApply) error

func (client *Client) OnNewFriendApply(handler NewFriendApplyHandler) {
	client.OnEvent(uim.ProviderEventNewFriendApply, uim.CastEventHandler(handler))
}

// 新好友
type NewContactHandler func(*cloudevents.Event, *uim.Contact) error

func (client *Client) OnNewContact(handler NewContactHandler) {
	client.OnEvent(uim.ProviderEventNewContact, uim.CastEventHandler(handler))
}

// 好友更新
type ContactUpdatedHandler func(*cloudevents.Event, *uim.ContactUpdate) error

func (client *Client) OnContactUpdated(handler ContactUpdatedHandler) {
	client.OnEvent(uim.ProviderEventContactUpdated, uim.CastEventHandler(handler))
}

// 好友删除
type ContactDeletedHandler func(*cloudevents.Event, *uim.ContactDeleted) error

func (client *Client) OnContactDeleted(handler ContactDeletedHandler) {
	client.OnEvent(uim.ProviderEventContactDeleted, uim.CastEventHandler(handler))
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

// 群组删除
type GroupDeletedHandler func(*cloudevents.Event, *uim.GroupDelete) error

func (client *Client) OnGroupDeleted(handler GroupDeletedHandler) {
	client.OnEvent(uim.ProviderEventGroupDeleted, uim.CastEventHandler(handler))
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

// 群成员删除
type GroupMemberDeletedHandler func(*cloudevents.Event, *uim.GroupMemberDelete) error

func (client *Client) OnGroupMemberDeleted(handler GroupMemberDeletedHandler) {
	client.OnEvent(uim.ProviderEventGroupMemberDeleted, uim.CastEventHandler(handler))
}

// 收到入群邀请
type NewGroupInvitationHandler func(*cloudevents.Event, *uim.GroupInvitation) error

func (client *Client) OnNewGroupInvitation(handler NewGroupInvitationHandler) {
	client.OnEvent(uim.ProviderEventNewGroupInvitation, uim.CastEventHandler(handler))
}

// 收到入群申请
type NewJoinGroupApplyHandler func(*cloudevents.Event, *uim.JoinGroupApply) error

func (client *Client) OnNewJoinGroupApply(handler NewJoinGroupApplyHandler) {
	client.OnEvent(uim.ProviderEventNewJoinGroupApply, uim.CastEventHandler(handler))
}

// 新会话
type NewConversationHandler func(*cloudevents.Event, *uim.Conversation) error

func (client *Client) OnNewConversation(handler NewConversationHandler) {
	client.OnEvent(uim.ProviderEventNewConversation, uim.CastEventHandler(handler))
}

// 会话更新
type ConversationUpdatedHandler func(*cloudevents.Event, *uim.ConversationUpdate) error

func (client *Client) OnConversationUpdated(handler ConversationUpdatedHandler) {
	client.OnEvent(uim.ProviderEventConversationUpdated, uim.CastEventHandler(handler))
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

// 新动态
type NewMomentHandler func(*cloudevents.Event, *uim.Moment) error

func (client *Client) OnNewMoment(handler NewMomentHandler) {
	client.OnEvent(uim.ProviderEventNewMoment, uim.CastEventHandler(handler))
}

// 动态更新
type MomentUpdatedHandler func(*cloudevents.Event, *uim.MomentUpdate) error

func (client *Client) OnMomentUpdated(handler MomentUpdatedHandler) {
	client.OnEvent(uim.ProviderEventMomentUpdated, uim.CastEventHandler(handler))
}

// 动态删除
type MomentDeletedHandler func(*cloudevents.Event, *uim.MomentDelete) error

func (client *Client) OnMomentDeleted(handler MomentDeletedHandler) {
	client.OnEvent(uim.ProviderEventMomentDeleted, uim.CastEventHandler(handler))
}

// 收到动态评论
type NewMomentCommentHandler func(*cloudevents.Event, *uim.MomentComment) error

func (client *Client) OnNewMomentComment(handler NewMomentCommentHandler) {
	client.OnEvent(uim.ProviderEventNewMomentComment, uim.CastEventHandler(handler))
}

// 动态评论更新
type MomentCommentUpdatedHandler func(*cloudevents.Event, *uim.MomentCommentUpdate) error

func (client *Client) OnMomentCommentUpdated(handler MomentCommentUpdatedHandler) {
	client.OnEvent(uim.ProviderEventMomentCommentUpdated, uim.CastEventHandler(handler))
}

// 动态评论删除
type MomentCommentDeletedHandler func(*cloudevents.Event, *uim.MomentCommentDelete) error

func (client *Client) OnMomentCommentDeleted(handler MomentCommentDeletedHandler) {
	client.OnEvent(uim.ProviderEventMomentCommentDeleted, uim.CastEventHandler(handler))
}

// 收到动态点赞
type NewMomentLikeHandler func(*cloudevents.Event, *uim.MomentLike) error

func (client *Client) OnNewMomentLike(handler NewMomentLikeHandler) {
	client.OnEvent(uim.ProviderEventNewMomentLike, uim.CastEventHandler(handler))
}

// 动态点赞删除
type MomentLikeDeletedHandler func(*cloudevents.Event, *uim.MomentLikeDelete) error

func (client *Client) OnMomentLikeDeleted(handler MomentLikeDeletedHandler) {
	client.OnEvent(uim.ProviderEventMomentLikeDeleted, uim.CastEventHandler(handler))
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
