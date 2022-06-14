package services

import (
	"time"
)

// 会话对象
type ConversationTarget struct {
	ID     string `json:"id,omitempty"`     // 对象ID，私聊会话时为对方用户ID，群聊会话时为群组ID
	Name   string `json:"name,omitempty"`   // 对象名称
	Avatar string `json:"avatar,omitempty"` // 对象头像
}

// 会话
type Conversation struct {
	ID              string                 `json:"id,omitempty"`               // 会话唯一ID
	Type            ConversationType       `json:"type,omitempty"`             // 会话类型
	User            *IMUser                `json:"user,omitempty"`             // 会话归属用户
	Target          *ConversationTarget    `json:"target,omitempty"`           // 会话对象
	LastMessage     *Message               `json:"last_message,omitempty"`     // 最新消息
	LastMessageAt   *time.Time             `json:"last_message_at,omitempty"`  // 最新消息时间
	Unread          int                    `json:"unread,omitempty"`           // 未读消息数
	Pinned          bool                   `json:"pinned,omitempty"`           // 是否置顶
	Metadata        map[string]interface{} `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]interface{} `json:"private_metadata,omitempty"` // 私有元数据
	CreatedAt       *time.Time             `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt       *time.Time             `json:"updated_at,omitempty"`       // 最后更新时间
}

// 新建会话
type CreateConversation struct {
	UserId          string                 `json:"user_id,omitempty"`          // 会话归属的用户ID
	Type            ConversationType       `json:"type,omitempty"`             // 会话类型
	TargetId        string                 `json:"target_id,omitempty"`        // 对象ID，私聊会话时为对方用户ID，群聊会话时为群组ID
	Pinned          bool                   `json:"pinned,omitempty"`           // 是否置顶
	Metadata        map[string]interface{} `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]interface{} `json:"private_metadata,omitempty"` // 私有元数据
}

// 更新会话
type UpdateConversation struct {
	ID              string                 `json:"id,omitempty"`               // 会话唯一ID
	LastMessageId   *string                `json:"last_message_id,omitempty"`  // 最新消息ID
	LastMessageAt   *time.Time             `json:"last_message_at,omitempty"`  // 最新消息时间
	Unread          *int                   `json:"unread,omitempty"`           // 未读消息数
	Pinned          *bool                  `json:"pinned,omitempty"`           // 是否置顶
	Metadata        map[string]interface{} `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]interface{} `json:"private_metadata,omitempty"` // 私有元数据
}
