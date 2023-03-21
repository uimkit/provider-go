package uim

// Provider 发送给 UIM 的事件
const (
	ProviderEventNewAccount       = "provider.new_account"       // 新账号
	ProviderEventAccountUpdated   = "provider.account_updated"   // 账号更新
	ProviderEventNewContact       = "provider.new_contact"       // 新好友
	ProviderEventNewFollower      = "provider.new_follower"      // 新粉丝
	ProviderEventNewFollowing     = "provider.new_following"     // 新关注的人
	ProviderEventNewFriendApply   = "provider.new_friend_apply"  // 新的好友申请
	ProviderEventNewMessage       = "provider.new_message"       // 收新消息
	ProviderEventMessageUpdated   = "provider.message_updated"   // 消息更新，如：撤回消息
	ProviderEventNewMetafield     = "provider.new_metafield"     // 新的元信息
	ProviderEventMetafieldUpdated = "provider.metafield_updated" // 元信息更新

	ProviderEventNewFriendReply       = "provider.new_friend_reply"       // 收到好友申请回复
	ProviderEventNewGroup             = "provider.new_group"              // 新群组
	ProviderEventGroupUpdated         = "provider.group_updated"          // 群组更新
	ProviderEventGroupDeleted         = "provider.group_deleted"          // 群组删除
	ProviderEventNewGroupMember       = "provider.new_group_member"       // 新群成员
	ProviderEventGroupMemberUpdated   = "provider.group_member_updated"   // 群成员更新
	ProviderEventGroupMemberDeleted   = "provider.group_member_deleted"   // 群成员删除
	ProviderEventNewGroupInvitation   = "provider.new_group_invitation"   // 收到入群邀请
	ProviderEventNewGroupApply        = "provider.new_group_apply"        // 收到入群申请
	ProviderEventNewMoment            = "provider.new_moment"             // 新动态
	ProviderEventMomentUpdated        = "provider.moment_updated"         // 动态更新
	ProviderEventMomentDeleted        = "provider.moment_deleted"         // 动态删除
	ProviderEventNewMomentComment     = "provider.new_moment_comment"     // 收到动态评论
	ProviderEventMomentCommentUpdated = "provider.moment_comment_updated" // 动态评论更新
	ProviderEventMomentCommentDeleted = "provider.moment_comment_deleted" // 动态评论被删除
	ProviderEventNewMomentLike        = "provider.new_moment_like"        // 收到动态点赞
	ProviderEventMomentLikeDeleted    = "provider.moment_like_deleted"    // 动态点赞被删除
)

// Provider 调用 UIM 的指令
const (
	ProviderCommandGetMetafield = "provider.get_metafield" // 查询元信息
)

// UIM 发送给 Provider 的事件
const ()

// UIM 调用 Provider 的指令
const (
	UIMCommandGetChannelInfo    = "uim.get_channel_info"    // 查询消息地址关联的信息
	UIMCommandSendMessage       = "uim.send_message"        // 发送消息
	UIMCommandAddContact        = "uim.add_contact"         // 发起好友申请
	UIMCommandAcceptFriendApply = "uim.accept_friend_apply" // 通过好友请求
	UIMCommandGetMomentList     = "uim.get_moment_list"     // 获取动态列表

	// deprecated
	UIMCommandUpdateAccount         = "uim.update_account"          // 更新账号资料
	UIMCommandUpdateContact         = "uim.update_contact"          // 更新联系人资料
	UIMCommandListContacts          = "uim.list_contacts"           // 查询账号的联系人列表
	UIMCommandApplyFriend           = "uim.apply_friend"            // 添加好友
	UIMCommandCreateGroup           = "uim.create_group"            // 创建群组
	UIMCommandUpdateGroup           = "uim.update_group"            // 更新群组资料
	UIMCommandListGroups            = "uim.list_groups"             // 查询账号的群组列表
	UIMCommandInviteToGroup         = "uim.invite_to_group"         // 邀请加入群组
	UIMCommandAcceptGroupInvitation = "uim.accept_group_invitation" // 接受入群邀请
	UIMCommandApplyJoinGroup        = "uim.apply_join_group"        // 申请加入群组
	UIMCommandAcceptGroupApply      = "uim.accept_group_apply"      // 通过入群申请
	UIMCommandListGroupMembers      = "uim.list_group_members"      // 查询群成员列表
	UIMCommandPublishMoment         = "uim.publish_moment"          // 发布动态
)
