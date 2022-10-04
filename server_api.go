package uim

import cloudevents "github.com/cloudevents/sdk-go/v2"

// UIM 需要的接口

// 新账号事件
type NewAccountHandler func(*cloudevents.Event, *IMAccount) error

func (client *Client) OnNewAccount(handler NewAccountHandler) {
	client.OnEvent(ProviderEventNewAccount, castEventHandler(handler))
}

// 账号更新
type AccountUpdatedHandler func(*cloudevents.Event, *IMAccountUpdate) error

func (client *Client) OnAccountUpdated(handler AccountUpdatedHandler) {
	client.OnEvent(ProviderEventAccountUpdated, castEventHandler(handler))
}

// 新的好友申请
type NewFriendApplyHandler func(*cloudevents.Event, *FriendApply) error

func (client *Client) OnNewFriendApply(handler NewFriendApplyHandler) {
	client.OnEvent(ProviderEventNewFriendApply, castEventHandler(handler))
}

// 新好友
type NewContactHandler func(*cloudevents.Event, *Contact) error

func (client *Client) OnNewContact(handler NewContactHandler) {
	client.OnEvent(ProviderEventNewContact, castEventHandler(handler))
}

// 好友更新
type ContactUpdatedHandler func(*cloudevents.Event, *ContactUpdate) error

func (client *Client) OnContactUpdated(handler ContactUpdatedHandler) {
	client.OnEvent(ProviderEventContactUpdated, castEventHandler(handler))
}

// 好友删除
type ContactDeletedHandler func(*cloudevents.Event, *ContactDeleted) error

func (client *Client) OnContactDeleted(handler ContactDeletedHandler) {
	client.OnEvent(ProviderEventContactDeleted, castEventHandler(handler))
}

// 新群组
type NewGroupHandler func(*cloudevents.Event, *Group) error

func (client *Client) OnNewGroup(handler NewGroupHandler) {
	client.OnEvent(ProviderEventNewGroup, castEventHandler(handler))
}

// 群组更新
type GroupUpdatedHandler func(*cloudevents.Event, *GroupUpdate) error

func (client *Client) OnGroupUpdated(handler GroupUpdatedHandler) {
	client.OnEvent(ProviderEventGroupUpdated, castEventHandler(handler))
}

// 群组删除
type GroupDeletedHandler func(*cloudevents.Event, *GroupDelete) error

func (client *Client) OnGroupDeleted(handler GroupDeletedHandler) {
	client.OnEvent(ProviderEventGroupDeleted, castEventHandler(handler))
}

// 新群成员
type NewGroupMemberHandler func(*cloudevents.Event, *GroupMember) error

func (client *Client) OnNewGroupMember(handler NewGroupMemberHandler) {
	client.OnEvent(ProviderEventNewGroupMember, castEventHandler(handler))
}

// 群成员更新
type GroupMemberUpdatedHandler func(*cloudevents.Event, *GroupMemberUpdate) error

func (client *Client) OnGroupMemberUpdated(handler GroupMemberUpdatedHandler) {
	client.OnEvent(ProviderEventGroupMemberUpdated, castEventHandler(handler))
}

// 群成员删除
type GroupMemberDeletedHandler func(*cloudevents.Event, *GroupMemberDelete) error

func (client *Client) OnGroupMemberDeleted(handler GroupMemberDeletedHandler) {
	client.OnEvent(ProviderEventGroupMemberDeleted, castEventHandler(handler))
}

// 收到入群邀请
type NewGroupInvitationHandler func(*cloudevents.Event, *GroupInvitation) error

func (client *Client) OnNewGroupInvitation(handler NewGroupInvitationHandler) {
	client.OnEvent(ProviderEventNewGroupInvitation, castEventHandler(handler))
}

// 收到入群申请
type NewJoinGroupApplyHandler func(*cloudevents.Event, *JoinGroupApply) error

func (client *Client) OnNewJoinGroupApply(handler NewJoinGroupApplyHandler) {
	client.OnEvent(ProviderEventNewJoinGroupApply, castEventHandler(handler))
}

// 新会话
type NewConversationHandler func(*cloudevents.Event, *Conversation) error

func (client *Client) OnNewConversation(handler NewConversationHandler) {
	client.OnEvent(ProviderEventNewConversation, castEventHandler(handler))
}

// 会话更新
type ConversationUpdatedHandler func(*cloudevents.Event, *ConversationUpdate) error

func (client *Client) OnConversationUpdated(handler ConversationUpdatedHandler) {
	client.OnEvent(ProviderEventConversationUpdated, castEventHandler(handler))
}

// 新消息
type NewMessageHandler func(*cloudevents.Event, *Message) error

func (client *Client) OnNewMessage(handler NewMessageHandler) {
	client.OnEvent(ProviderEventNewMessage, castEventHandler(handler))
}

// 消息更新
type MessageUpdatedHandler func(*cloudevents.Event, *MessageUpdate) error

func (client *Client) OnMessageUpdated(handler MessageUpdatedHandler) {
	client.OnEvent(ProviderEventMessageUpdated, castEventHandler(handler))
}

// 新动态
type NewMomentHandler func(*cloudevents.Event, *Moment) error

func (client *Client) OnNewMoment(handler NewMomentHandler) {
	client.OnEvent(ProviderEventNewMoment, castEventHandler(handler))
}

// 动态更新
type MomentUpdatedHandler func(*cloudevents.Event, *MomentUpdate) error

func (client *Client) OnMomentUpdated(handler MomentUpdatedHandler) {
	client.OnEvent(ProviderEventMomentUpdated, castEventHandler(handler))
}

// 动态删除
type MomentDeletedHandler func(*cloudevents.Event, *MomentDelete) error

func (client *Client) OnMomentDeleted(handler MomentDeletedHandler) {
	client.OnEvent(ProviderEventMomentDeleted, castEventHandler(handler))
}

// 收到动态评论
type NewMomentCommentHandler func(*cloudevents.Event, *MomentComment) error

func (client *Client) OnNewMomentComment(handler NewMomentCommentHandler) {
	client.OnEvent(ProviderEventNewMomentComment, castEventHandler(handler))
}

// 动态评论更新
type MomentCommentUpdatedHandler func(*cloudevents.Event, *MomentCommentUpdate) error

func (client *Client) OnMomentCommentUpdated(handler MomentCommentUpdatedHandler) {
	client.OnEvent(ProviderEventMomentCommentUpdated, castEventHandler(handler))
}

// 动态评论删除
type MomentCommentDeletedHandler func(*cloudevents.Event, *MomentCommentDelete) error

func (client *Client) OnMomentCommentDeleted(handler MomentCommentDeletedHandler) {
	client.OnEvent(ProviderEventMomentCommentDeleted, castEventHandler(handler))
}

// 收到动态点赞
type NewMomentLikeHandler func(*cloudevents.Event, *MomentLike) error

func (client *Client) OnNewMomentLike(handler NewMomentLikeHandler) {
	client.OnEvent(ProviderEventNewMomentLike, castEventHandler(handler))
}

// 动态点赞删除
type MomentLikeDeletedHandler func(*cloudevents.Event, *MomentLikeDelete) error

func (client *Client) OnMomentLikeDeleted(handler MomentLikeDeletedHandler) {
	client.OnEvent(ProviderEventMomentLikeDeleted, castEventHandler(handler))
}

// 新的元数据
type NewMetafieldHandler func(*cloudevents.Event, *Metafield) error

func (client *Client) OnNewMetafield(handler NewMetafieldHandler) {
	client.OnEvent(ProviderEventNewMetafield, castEventHandler(handler))
}

// 元数据更新
type MetafieldUpdatedHandler func(*cloudevents.Event, *MetafieldUpdate) error

func (client *Client) OnMetafieldUpdated(handler MetafieldUpdatedHandler) {
	client.OnEvent(ProviderEventMetafieldUpdated, castEventHandler(handler))
}

// 查询元数据
type GetMetafieldHandler func(*cloudevents.Event, *GetMetafieldRequest) (*GetMetafieldResponse, error)

func (client *Client) OnGetMetafield(handler GetMetafieldHandler) {
	client.OnEvent(ProviderCommandGetMetafield, castCommandHandler(handler))
}
