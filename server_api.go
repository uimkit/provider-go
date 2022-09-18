package uim

// UIM 需要的接口

// 新账号事件
type NewAccountHandler func(*IMAccount) error

func (client *Client) OnNewAccount(handler NewAccountHandler) {
	client.OnEvent(ProviderEventNewAccount, castEventHandler(handler))
}

// 账号更新
type AccountUpdatedHandler func(*IMAccountUpdate) error

func (client *Client) OnAccountUpdated(handler AccountUpdatedHandler) {
	client.OnEvent(ProviderEventAccountUpdated, castEventHandler(handler))
}

// 新的好友申请
type NewFriendApplyHandler func(*FriendApply) error

func (client *Client) OnNewFriendApply(handler NewFriendApplyHandler) {
	client.OnEvent(ProviderEventNewFriendApply, castEventHandler(handler))
}

// 收到好友申请的消息回复
type NewFriendReplyHandler func(*FriendReply) error

func (client *Client) OnNewFriendReply(handler NewFriendReplyHandler) {
	client.OnEvent(ProviderEventNewFriendReply, castEventHandler(handler))
}

// 新好友
type NewContactHandler func(*Contact) error

func (client *Client) OnNewContact(handler NewContactHandler) {
	client.OnEvent(ProviderEventNewContact, castEventHandler(handler))
}

// 好友更新
type ContactUpdatedHandler func(*ContactUpdate) error

func (client *Client) OnContactUpdated(handler ContactUpdatedHandler) {
	client.OnEvent(ProviderEventContactUpdated, castEventHandler(handler))
}

// 好友删除
type ContactDeletedHandler func(*ContactDeleted) error

func (client *Client) OnContactDeleted(handler ContactDeletedHandler) {
	client.OnEvent(ProviderEventContactDeleted, castEventHandler(handler))
}

// 新群组
type NewGroupHandler func(*Group) error

func (client *Client) OnNewGroup(handler NewGroupHandler) {
	client.OnEvent(ProviderEventNewGroup, castEventHandler(handler))
}

// 群组更新
type GroupUpdatedHandler func(*GroupUpdate) error

func (client *Client) OnGroupUpdated(handler GroupUpdatedHandler) {
	client.OnEvent(ProviderEventGroupUpdated, castEventHandler(handler))
}

// 群组删除
type GroupDeletedHandler func(*GroupDelete) error

func (client *Client) OnGroupDeleted(handler GroupDeletedHandler) {
	client.OnEvent(ProviderEventGroupDeleted, castEventHandler(handler))
}

// 新群成员
type NewGroupMemberHandler func(*GroupMember) error

func (client *Client) OnNewGroupMember(handler NewGroupMemberHandler) {
	client.OnEvent(ProviderEventNewGroupMember, castEventHandler(handler))
}

// 群成员更新
type GroupMemberUpdatedHandler func(*GroupMemberUpdate) error

func (client *Client) OnGroupMemberUpdated(handler GroupMemberUpdatedHandler) {
	client.OnEvent(ProviderEventGroupMemberUpdated, castEventHandler(handler))
}

// 群成员删除
type GroupMemberDeletedHandler func(*GroupMemberDelete) error

func (client *Client) OnGroupMemberDeleted(handler GroupMemberDeletedHandler) {
	client.OnEvent(ProviderEventGroupMemberDeleted, castEventHandler(handler))
}

// 收到入群邀请
type NewGroupInvitationHandler func(*GroupInvitation) error

func (client *Client) OnNewGroupInvitation(handler NewGroupInvitationHandler) {
	client.OnEvent(ProviderEventNewGroupInvitation, castEventHandler(handler))
}

// 收到入群申请
type NewJoinGroupApplyHandler func(*JoinGroupApply) error

func (client *Client) OnNewJoinGroupApply(handler NewJoinGroupApplyHandler) {
	client.OnEvent(ProviderEventNewJoinGroupApply, castEventHandler(handler))
}

// 新消息
type NewMessageHandler func(*Message) error

func (client *Client) OnNewMessage(handler NewMessageHandler) {
	client.OnEvent(ProviderEventNewMessage, castEventHandler(handler))
}

// 消息更新
type MessageUpdatedHandler func(*MessageUpdate) error

func (client *Client) OnMessageUpdated(handler MessageUpdatedHandler) {
	client.OnEvent(ProviderEventMessageUpdated, castEventHandler(handler))
}

// 新动态
type NewMomentHandler func(*Moment) error

func (client *Client) OnNewMoment(handler NewMomentHandler) {
	client.OnEvent(ProviderEventNewMoment, castEventHandler(handler))
}

// 动态更新
type MomentUpdatedHandler func(*MomentUpdate) error

func (client *Client) OnMomentUpdated(handler MomentUpdatedHandler) {
	client.OnEvent(ProviderEventMomentUpdated, castEventHandler(handler))
}

// 动态删除
type MomentDeletedHandler func(*MomentDelete) error

func (client *Client) OnMomentDeleted(handler MomentDeletedHandler) {
	client.OnEvent(ProviderEventMomentDeleted, castEventHandler(handler))
}

// 收到动态评论
type NewMomentCommentHandler func(*MomentComment) error

func (client *Client) OnNewMomentComment(handler NewMomentCommentHandler) {
	client.OnEvent(ProviderEventNewMomentComment, castEventHandler(handler))
}

// 动态评论更新
type MomentCommentUpdatedHandler func(*MomentCommentUpdate) error

func (client *Client) OnMomentCommentUpdated(handler MomentCommentUpdatedHandler) {
	client.OnEvent(ProviderEventMomentCommentUpdated, castEventHandler(handler))
}

// 动态评论删除
type MomentCommentDeletedHandler func(*MomentCommentDelete) error

func (client *Client) OnMomentCommentDeleted(handler MomentCommentDeletedHandler) {
	client.OnEvent(ProviderEventMomentCommentDeleted, castEventHandler(handler))
}

// 收到动态点赞
type NewMomentLikeHandler func(*MomentLike) error

func (client *Client) OnNewMomentLike(handler NewMomentLikeHandler) {
	client.OnEvent(ProviderEventNewMomentLike, castEventHandler(handler))
}

// 动态点赞删除
type MomentLikeDeletedHandler func(*MomentLikeDelete) error

func (client *Client) OnMomentLikeDeleted(handler MomentLikeDeletedHandler) {
	client.OnEvent(ProviderEventMomentLikeDeleted, castEventHandler(handler))
}

// 新的元数据
type NewMetafieldHandler func(*Metafield) error

func (client *Client) OnNewMetafield(handler NewMetafieldHandler) {
	client.OnEvent(ProviderEventNewMetafield, castEventHandler(handler))
}

// 元数据更新
type MetafieldUpdatedHandler func(*MetafieldUpdate) error

func (client *Client) OnMetafieldUpdated(handler MetafieldUpdatedHandler) {
	client.OnEvent(ProviderEventMetafieldUpdated, castEventHandler(handler))
}

// 查询元数据
type GetMetafieldHandler func(*GetMetafield) (*GetMetafieldResponse, error)

func (client *Client) OnGetMetafield(handler GetMetafieldHandler) {
	client.OnEvent(ProviderCommandGetMetafield, castCommandHandler(handler))
}
