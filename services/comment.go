package services

import (
	"time"
)

// 评论
type Comment struct {
	ID              string          `json:"id,omitempty"`               // 评论唯一ID
	CommentId       string          `json:"comment_id,omitempty"`       // 平台评论ID
	User            *IMUser         `json:"user,omitempty"`             // 发表人用户信息
	Resource        string          `json:"resource,omitempty"`         // 评论资源类型
	ResourceId      string          `json:"resource_id,omitempty"`      // 评论资源ID
	ReplyId         string          `json:"reply_id,omitempty"`         // 直接回复的评论ID
	OriginId        string          `json:"origin_id,omitempty"`        // 多级回复时，最初的评论ID
	PublishedAt     *time.Time      `json:"published_at,omitempty"`     // 发表时间
	Content         *CommentContent `json:"content,omitempty"`          // 动态内容
	Metadata        map[string]any  `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any  `json:"private_metadata,omitempty"` // 私有元数据
	CreatedAt       *time.Time      `json:"created_at,omitempty"`       // 创建时间
	UpdatedAt       *time.Time      `json:"updated_at,omitempty"`       // 最后更新时间
}

// 发表评论
type PostComment struct {
	CommentId       string          `json:"comment_id,omitempty"`       // 平台评论ID
	UserId          string          `json:"user_id,omitempty"`          // 发表人用户ID
	Resource        string          `json:"resource,omitempty"`         // 评论资源类型
	ResourceId      string          `json:"resource_id,omitempty"`      // 评论资源ID
	ReplyId         string          `json:"reply_id,omitempty"`         // 直接回复的评论ID
	OriginId        string          `json:"origin_id,omitempty"`        // 多级回复时，最初的评论ID
	PublishedAt     *time.Time      `json:"published_at,omitempty"`     // 发表时间
	Content         *CommentContent `json:"content,omitempty"`          // 评论内容
	Metadata        map[string]any  `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any  `json:"private_metadata,omitempty"` // 私有元数据
}

// 修改评论
type UpdateComment struct {
	ID              string          `json:"id,omitempty"`               // 评论唯一ID
	Content         *CommentContent `json:"content,omitempty"`          // 评论内容
	Metadata        map[string]any  `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any  `json:"private_metadata,omitempty"` // 私有元数据
}

// 评论内容
type CommentContent struct {
	Type CommentType `json:"type,omitempty"` // 类型
	Body any         `json:"body,omitempty"` // 内容
}

// 文本动态
type TextCommentBody struct {
	Content string `json:"content,omitempty"` // 文本
}
