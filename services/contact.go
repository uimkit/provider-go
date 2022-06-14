package services

import "time"

// 联系人
type Contact struct {
	ID              string         `json:"id,omitempty"`               // 联系人唯一ID
	FromUser        *IMUser        `json:"from_user,omitempty"`        // 所属账号用户
	ToUser          *IMUser        `json:"to_user,omitempty"`          // 联系人的用户
	Alias           string         `json:"alias,omitempty"`            // 备注名
	Remark          string         `json:"remark,omitempty"`           // 备注说明
	Tags            []string       `json:"tags,omitempty"`             // 标签
	Blocked         bool           `json:"blocked,omitempty"`          // 是否拉黑
	Marked          bool           `json:"marked,omitempty"`           // 是否星标
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
	CreatedAt       *time.Time     `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`       // 最后更新时间
}

// 创建联系人
type CreateContact struct {
	FromUserId      string         `json:"from_user_id,omitempty"`     // 所属账号的用户ID
	ToUserId        string         `json:"to_user_id,omitempty"`       // 联系人的用户ID
	Alias           string         `json:"alias,omitempty"`            // 备注名
	Remark          string         `json:"remark,omitempty"`           // 备注说明
	Tags            []string       `json:"tags,omitempty"`             // 标签
	Blocked         bool           `json:"blocked,omitempty"`          // 是否拉黑
	Marked          bool           `json:"marked,omitempty"`           // 是否星标
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 更新联系人
type UpdateContact struct {
	ID              string         `json:"id,omitempty"`               // 联系人唯一ID
	Alias           *string        `json:"alias,omitempty"`            // 备注名
	Remark          *string        `json:"remark,omitempty"`           // 备注说明
	Tags            []string       `json:"tags,omitempty"`             // 标签
	Blocked         *bool          `json:"blocked,omitempty"`          // 是否拉黑
	Marked          *bool          `json:"marked,omitempty"`           // 是否星标
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 好友请求
type FriendApply struct {
	ID           string    `json:"id,omitempty"`            // 好友请求ID
	User         *IMUser   `json:"user,omitempty"`          // 申请人信息
	HelloMessage string    `json:"hello_message,omitempty"` // 打招呼留言
	CreatedAt    time.Time `json:"created_at,omitempty"`    // 申请时间
}

// 创建好友请求
type NewFriendApply struct {
	AccountId    string   `json:"account_id,omitempty"`    // 操作账号ID
	Contacts     []string `json:"contacts,omitempty"`      // 联系人列表，可用于在平台添加联系人，如：手机号、抖音号等
	HelloMessage string   `json:"hello_message,omitempty"` // 打招呼留言
}
