package services

import "time"

// 群组
type Group struct {
	ID              string         `json:"id,omitempty"`               // 群组唯一ID
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	Owner           *IMUser        `json:"owner,omitempty"`            // 群主信息
	Name            string         `json:"name,omitempty"`             // 名称
	Avatar          string         `json:"avatar,omitempty"`           // 头像URL
	Announcement    string         `json:"announcement,omitempty"`     // 群公告
	Description     string         `json:"description,omitempty"`      // 群介绍
	MemberCount     int32          `json:"member_count,omitempty"`     // 群成员数量
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
	CreatedAt       *time.Time     `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`       // 最后更新时间
}

// 创建群组
type CreateGroup struct {
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	OwnerId         string         `json:"owner_id,omitempty"`         // 操作账号对应的用户ID
	Name            string         `json:"name,omitempty"`             // 名称
	Avatar          string         `json:"avatar,omitempty"`           // 头像URL
	Announcement    string         `json:"announcement,omitempty"`     // 群公告
	Description     string         `json:"description,omitempty"`      // 群介绍
	UserIds         []string       `json:"user_ids,omitempty"`         // 加入群组的用户ID列表
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 修改群组信息
type UpdateGroup struct {
	ID              string         `json:"id,omitempty"`               // 群组唯一ID
	OwnerId         *string        `json:"owner_id,omitempty"`         // 操作账号对应的用户ID
	Name            *string        `json:"name,omitempty"`             // 名称
	Avatar          *string        `json:"avatar,omitempty"`           // 头像URL
	Announcement    *string        `json:"announcement,omitempty"`     // 群公告
	Description     *string        `json:"description,omitempty"`      // 群介绍
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 群组成员
type GroupMember struct {
	ID              string         `json:"id,omitempty"`               // 群成员唯一ID
	MemberId        string         `json:"member_id,omitempty"`        // 平台群成员ID
	GroupId         string         `json:"group_id,omitempty"`         // 群组ID
	User            *IMUser        `json:"user,omitempty"`             // 关联的用户信息
	IsOwner         bool           `json:"is_owner,omitempty"`         // 是否群主
	IsAdmin         bool           `json:"is_admin,omitempty"`         // 是否管理员
	Alias           string         `json:"alias,omitempty"`            // 群内备注名
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
	CreatedAt       *time.Time     `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`       // 最后更新时间
}

// 创建群组成员
type CreateGroupMember struct {
	MemberId        string         `json:"member_id,omitempty"`        // 平台群成员ID
	GroupId         string         `json:"group_id,omitempty"`         // 群组ID
	UserId          string         `json:"user_id,omitempty"`          // 关联的用户ID
	IsOwner         bool           `json:"is_owner,omitempty"`         // 是否群主
	IsAdmin         bool           `json:"is_admin,omitempty"`         // 是否管理员
	Alias           string         `json:"alias,omitempty"`            // 群内备注名
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 修改群组成员信息
type UpdateGroupMember struct {
	ID              string         `json:"id,omitempty"`               // 群成员唯一ID
	IsOwner         *bool          `json:"is_owner,omitempty"`         // 是否群主
	IsAdmin         *bool          `json:"is_admin,omitempty"`         // 是否管理员
	Alias           *string        `json:"alias,omitempty"`            // 群内备注名
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 直接拉人入群
type AddGroupMembersRequest struct {
	AccountId string   `json:"account_id,omitempty"` // 操作账号ID
	GroupId   string   `json:"group_id,omitempty"`   // 群组ID
	UserIds   []string `json:"user_ids,omitempty"`   // 加入群组的用户ID列表
}

// 入群邀请
type GroupInvitation struct {
	ID           string    `json:"id,omitempty"`            // 入群邀请ID
	Group        *Group    `json:"group,omitempty"`         // 群组信息
	Inviter      *IMUser   `json:"inviter,omitempty"`       // 邀请人信息
	HelloMessage string    `json:"hello_message,omitempty"` // 打招呼留言
	CreatedAt    time.Time `json:"created_at,omitempty"`    // 邀请时间
}

// 创建入群邀请
type NewGroupInvitationRequest struct {
	AccountId    string   `json:"account_id,omitempty"`    // 操作账号ID
	GroupId      string   `json:"group_id,omitempty"`      // 群组ID
	UserIds      []string `json:"user_ids,omitempty"`      // 加入群组的用户ID列表
	HelloMessage string   `json:"hello_message,omitempty"` // 打招呼留言
}

// 接收入群邀请
type AcceptGroupInvitationRequest struct {
	AccountId    string `json:"account_id,omitempty"`    // 操作账号ID
	InvitationId string `json:"invitation_id,omitempty"` // 邀请ID
	GroupId      string `json:"group_id,omitempty"`      // 群组ID
}

// 转让群组
type TransferGroupRequest struct {
	AccountId string `json:"account_id,omitempty"` // 操作账号ID
	GroupId   string `json:"group_id,omitempty"`   // 群组ID
	UserId    string `json:"user_id,omitempty"`    // 新群主用户ID
}

// 加入群组申请
type JoinGroupApply struct {
	ID           string    `json:"id,omitempty"`            // 入群申请ID
	User         *IMUser   `json:"user,omitempty"`          // 申请用户
	Group        *Group    `json:"group,omitempty"`         // 群组信息
	HelloMessage string    `json:"hello_message,omitempty"` // 打招呼留言
	CreatedAt    time.Time `json:"created_at,omitempty"`    // 申请时间
}

// 创建入群申请
type NewJoinGroupApplyRequest struct {
	AccountId    string `json:"account_id,omitempty"`    // 操作账号ID
	GroupId      string `json:"group_id,omitempty"`      // 群组ID
	HelloMessage string `json:"hello_message,omitempty"` // 打招呼留言
}

// 通过入群申请
type ApproveJoinGroupApplyRequest struct {
	AccountId string `json:"account_id,omitempty"` // 操作账号ID
	RequestId string `json:"request_id,omitempty"` // 入群申请ID
	GroupId   string `json:"group_id,omitempty"`   // 群组ID
}

// 拒绝入群申请
type RejectJoinGroupApplyRequest struct {
	AccountId string `json:"account_id,omitempty"` // 操作账号ID
	RequestId string `json:"request_id,omitempty"` // 入群申请ID
	GroupId   string `json:"group_id,omitempty"`   // 群组ID
}

// 退群
type QuitGroupRequest struct {
	AccountId string `json:"account_id,omitempty"` // 操作账号ID
	GroupId   string `json:"group_id,omitempty"`   // 群组ID
}

// 解散群组
type DismissGroupRequest struct {
	AccountId string `json:"account_id,omitempty"` // 操作账号ID
	GroupId   string `json:"group_id,omitempty"`   // 群组ID
}

// 踢除群成员
type KickGroupMemberRequest struct {
	AccountId string `json:"account_id,omitempty"` // 操作账号ID
	GroupId   string `json:"group_id,omitempty"`   // 群组ID
	MemberId  string `json:"member_id,omitempty"`  // 群成员ID
}

// 任命群管理员
type AppointGroupAdminRequest struct {
	AccountId string `json:"account_id,omitempty"` // 操作账号ID
	GroupId   string `json:"group_id,omitempty"`   // 群组ID
	MemberId  string `json:"member_id,omitempty"`  // 群成员ID
}

// 解除群管理员
type RevokeGroupAdminRequest struct {
	AccountId string `json:"account_id,omitempty"` // 操作账号ID
	GroupId   string `json:"group_id,omitempty"`   // 群组ID
	MemberId  string `json:"member_id,omitempty"`  // 群成员ID
}
