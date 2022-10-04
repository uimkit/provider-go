package uim

// Provider 需要的接口

// 新账号
func (client *Client) NewAccount(account *IMAccount) error {
	return client.SendEvent(client.newEvent(ProviderEventNewAccount, account))
}

// 账号更新
func (client *Client) AccountUpdated(account *IMAccountUpdate) error {
	return client.SendEvent(client.newEvent(ProviderEventAccountUpdated, account))
}

// 新的好友申请
func (client *Client) NewFriendApply(apply *FriendApply) error {
	return client.SendEvent(client.newEvent(ProviderEventNewFriendApply, apply))
}

// 新好友
func (client *Client) NewContact(contact *Contact) error {
	return client.SendEvent(client.newEvent(ProviderEventNewContact, contact))
}

// 好友更新
func (client *Client) ContactUpdated(contact *ContactUpdate) error {
	return client.SendEvent(client.newEvent(ProviderEventContactUpdated, contact))
}

// 好友删除
func (client *Client) ContactDeleted(contact *ContactDeleted) error {
	return client.SendEvent(client.newEvent(ProviderEventContactDeleted, contact))
}

// 新群组
func (client *Client) NewGroup(group *Group) error {
	return client.SendEvent(client.newEvent(ProviderEventNewGroup, group))
}

// 群组更新
func (client *Client) GroupUpdated(group *GroupUpdate) error {
	return client.SendEvent(client.newEvent(ProviderEventGroupUpdated, group))
}

// 群组删除
func (client *Client) GroupDeleted(group *GroupDelete) error {
	return client.SendEvent(client.newEvent(ProviderEventGroupDeleted, group))
}

// 新群成员
func (client *Client) NewGroupMember(member *GroupMember) error {
	return client.SendEvent(client.newEvent(ProviderEventNewGroupMember, member))
}

// 群成员更新
func (client *Client) GroupMemberUpdated(member *GroupMemberUpdate) error {
	return client.SendEvent(client.newEvent(ProviderEventGroupMemberUpdated, member))
}

// 群成员删除
func (client *Client) GroupMemberDeleted(member *GroupMemberDelete) error {
	return client.SendEvent(client.newEvent(ProviderEventGroupMemberDeleted, member))
}

// 收到入群邀请
func (client *Client) NewGroupInvitation(invitation *GroupInvitation) error {
	return client.SendEvent(client.newEvent(ProviderEventNewGroupInvitation, invitation))
}

// 收到入群申请
func (client *Client) NewJoinGroupApply(apply *JoinGroupApply) error {
	return client.SendEvent(client.newEvent(ProviderEventNewJoinGroupApply, apply))
}

// 新会话
func (client *Client) NewConversation(conversation *Conversation) error {
	return client.SendEvent(client.newEvent(ProviderEventNewConversation, conversation))
}

// 会话更新
func (client *Client) ConversationUpdated(conversation *ConversationUpdate) error {
	return client.SendEvent(client.newEvent(ProviderEventConversationUpdated, conversation))
}

// 新消息
func (client *Client) NewMessage(message *Message) error {
	return client.SendEvent(client.newEvent(ProviderEventNewMessage, message))
}

// 消息更新
func (client *Client) MessageUpdated(message *MessageUpdate) error {
	return client.SendEvent(client.newEvent(ProviderEventMessageUpdated, message))
}

// 新动态
func (client *Client) NewMoment(moment *Moment) error {
	return client.SendEvent(client.newEvent(ProviderEventNewMoment, moment))
}

// 动态更新
func (client *Client) MomentUpdated(moment *MomentUpdate) error {
	return client.SendEvent(client.newEvent(ProviderEventMomentUpdated, moment))
}

// 动态删除
func (client *Client) MomentDeleted(moment *MomentDelete) error {
	return client.SendEvent(client.newEvent(ProviderEventMomentDeleted, moment))
}

// 收到动态评论
func (client *Client) NewMomentComment(comment *MomentComment) error {
	return client.SendEvent(client.newEvent(ProviderEventNewMomentComment, comment))
}

// 动态评论更新
func (client *Client) MomentCommentUpdated(comment *MomentCommentUpdate) error {
	return client.SendEvent(client.newEvent(ProviderEventMomentCommentUpdated, comment))
}

// 动态评论删除
func (client *Client) MomentCommentDeleted(comment *MomentCommentDelete) error {
	return client.SendEvent(client.newEvent(ProviderEventMomentCommentDeleted, comment))
}

// 收到动态点赞
func (client *Client) NewMomentLike(like *MomentLike) error {
	return client.SendEvent(client.newEvent(ProviderEventNewMomentLike, like))
}

// 动态点赞删除
func (client *Client) MomentLikeDeleted(like *MomentLikeDelete) error {
	return client.SendEvent(client.newEvent(ProviderEventMomentLikeDeleted, like))
}

// 新的元数据
func (client *Client) NewMetafield(metafield *Metafield) error {
	return client.SendEvent(client.newEvent(ProviderEventNewMetafield, metafield))
}

// 元数据更新
func (client *Client) MetafieldUpdated(metafield *MetafieldUpdate) error {
	return client.SendEvent(client.newEvent(ProviderEventMetafieldUpdated, metafield))
}

// 查询元数据
func (client *Client) GetMetafield(metafield *GetMetafieldRequest) (*GetMetafieldResponse, error) {
	return castCommandResponse[*GetMetafieldResponse](
		client.InvokeCommand(
			client.newEvent(ProviderCommandGetMetafield, metafield),
			&GetMetafieldResponse{},
		),
	)
}
