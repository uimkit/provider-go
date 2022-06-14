package services

import (
	"time"
)

// IM账号
type IMAccount struct {
	ID              string         `json:"id,omitempty"`               // 账号唯一ID
	User            *IMUser        `json:"user,omitempty"`             // 关联的用户
	Presence        Presence       `json:"presence,omitempty"`         // 状态
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
	CreatedAt       *time.Time     `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt       *time.Time     `json:"updated_at,omitempty"`       // 最后更新时间
}

// 创建IM账号
type CreateIMAccount struct {
	UserId          string         `json:"user_id,omitempty"`          // 关联的用户ID
	Presence        Presence       `json:"presence,omitempty"`         // 状态
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 更新IM账号
type UpdateIMAccount struct {
	ID              string         `json:"id,omitempty"`               // 账号唯一ID
	Presence        Presence       `json:"presence,omitempty"`         // 状态
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// TODO
type AccountBinding struct {
	Provider  string    // IM提供者标识符
	AccountId string    // IM账户ID
	UserId    string    // 系统用户ID
	CreatedAt time.Time // 绑定时间
}
