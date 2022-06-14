package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

/*
可参考融云:
https://www.rongcloud.cn/docs/api/js/imlib/v5/interfaces/ILocationMessageBody.html
*/

// 消息收发身份识别，序列化规则： provider|provider_id/strategy，如：wechat|wx_id/cs
type MessageIdentity struct {
	ID       string `json:"id,omitempty"`       // 平台唯一ID
	Provider string `json:"provider,omitempty"` // 即服务提供者
	Strategy string `json:"strategy,omitempty"` // 收发消息使用的服务提供者策略
}

func (m *MessageIdentity) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%s|%s/%s", m.Provider, m.ID, m.Strategy)
	return json.Marshal(s)
}

func (m *MessageIdentity) UnmarshalJSON(data []byte) error {
	str := ""
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	strs := strings.Split(str, "|")
	if len(strs) != 2 {
		return errors.New("invalid message indentity")
	}
	m.Provider = strs[0]
	strs = strings.Split(strs[1], "/")
	if len(strs) != 2 {
		return errors.New("invalid message indentity")
	}
	m.ID = strs[0]
	m.Strategy = strs[1]
	return nil
}

// 消息收发对象信息
type MessageTarget struct {
	ID       string           `json:"id,omitempty"`       // 对象唯一ID，如：用户ID、群组ID
	Name     string           `json:"name,omitempty"`     // 对象名称
	Avatar   string           `json:"avatar,omitempty"`   // 对象头像
	Identity *MessageIdentity `json:"identity,omitempty"` // 对象的身份识别
}

// 消息
type Message struct {
	ID               string           `json:"id,omitempty"`                // 消息唯一ID
	MessageId        string           `json:"message_id,omitempty"`        // 平台消息ID
	ConversationType ConversationType `json:"conversation_type,omitempty"` // 所属的会话类型
	Seq              int              `json:"seq,omitempty"`               // 序列号，在会话中唯一且有序增长，用于确保消息顺序
	From             *MessageTarget   `json:"from,omitempty"`              // 发送方信息，私聊会话时是用户，群聊会话时是用户
	To               *MessageTarget   `json:"to,omitempty"`                // 接收方信息，私聊会话时是用户，群聊会话时是群组
	MentionedType    MentionedType    `json:"mentioned_type,omitempty"`    // @用户类型
	MentionedUsers   []*IMUser        `json:"mentioned_users"`             // @用户列表
	SentAt           *time.Time       `json:"sent_at,omitempty"`           // 发送时间
	Payload          *MessagePayload  `json:"payload,omitempty"`           // 消息内容
	Revoked          bool             `json:"revoked,omitempty"`           // 是否撤回
	Metadata         map[string]any   `json:"metadata,omitempty"`          // 公开元数据
	PrivateMetadata  map[string]any   `json:"private_metadata,omitempty"`  // 私有元数据
	CreatedAt        *time.Time       `json:"created_at,omitempty"`        // 创建时间
	UpdatedAt        *time.Time       `json:"updated_at,omitempty"`        // 最后更新时间
}

// 发送消息请求，要么提供 conversion_id，要么提供 conversation_type & from & to
type PostMessage struct {
	MessageId        string           `json:"message_id,omitempty"`         // 平台消息ID
	ConversationID   string           `json:"conversation_id,omitempty"`    // 所属的会话ID
	ConversationType ConversationType `json:"conversation_type,omitempty"`  // 所属的会话类型
	From             *MessageIdentity `json:"from,omitempty"`               // 发送人，私聊会话、群聊会话时是用户
	To               *MessageIdentity `json:"to,omitempty"`                 // 接收人，私聊会话时是用户，群聊会话时是群组
	Seq              int              `json:"seq,omitempty"`                // 序列号，在会话中唯一且有序增长，用于确保消息顺序
	MentionedType    MentionedType    `json:"mentioned_type,omitempty"`     // @用户类型
	MentionedUserIds []string         `json:"mentioned_user_ids,omitempty"` // @用户列表
	SentAt           *time.Time       `json:"sent_at,omitempty"`            // 发送时间
	Payload          *MessagePayload  `json:"payload,omitempty"`            // 消息内容
	Revoked          bool             `json:"revoked,omitempty"`            // 是否撤回
	Metadata         map[string]any   `json:"metadata,omitempty"`           // 公开元数据
	PrivateMetadata  map[string]any   `json:"private_metadata,omitempty"`   // 私有元数据
}

type UpdateMessage struct {
	ID              string         `json:"id,omitempty"`               // 消息唯一ID
	Revoked         *bool          `json:"revoked,omitempty"`          // 是否撤回
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// 消息内容
type MessagePayload struct {
	Type MessageType `json:"type,omitempty"` // 消息类型
	Body any         `json:"body,omitempty"` // 消息体
}

// 文本消息内容
type TextMessageBody struct {
	Content string `json:"content,omitempty"` // 文本内容
}

// 图片消息内容
type ThumbMessageBody struct {
	URL    string `json:"url,omitempty"`    // 图片URL
	Width  int    `json:"width,omitempty"`  // 宽度（像素）
	Height int    `json:"height,omitempty"` // 高度（像素）
	Ext    string `json:"ext,omitempty"`    // 类型，如：png、jpeg
}

type ImageMessageBody struct {
	URL    string            `json:"url,omitempty"`    // 图片URL
	Width  int               `json:"width,omitempty"`  // 宽度（像素）
	Height int               `json:"height,omitempty"` // 高度（像素）
	Size   int               `json:"size,omitempty"`   // 大小（字节）
	Ext    string            `json:"ext,omitempty"`    // 类型，如：png、jpeg
	MD5    string            `json:"md5,omitempty"`    // 文件内容MD5
	Thumb  *ThumbMessageBody `json:"thumb,omitempty"`  // 缩略图
}

// 语音消息内容
type VoiceMessageBody struct {
	URL      string `json:"url,omitempty"`      // 语音URL
	Duration int    `json:"duration,omitempty"` // 时长（毫秒）
	Size     int    `json:"size,omitempty"`     // 大小（字节）
	Ext      string `json:"ext,omitempty"`      // 类型，如：mp3
	MD5      string `json:"md5,omitempty"`      // 文件内容MD5
}

// 视频消息内容
type VideoMessageBody struct {
	URL      string            `json:"url,omitempty"`      // 视频URL
	Duration int               `json:"duration,omitempty"` // 时长（毫秒）
	Width    int               `json:"width,omitempty"`    // 宽度（像素）
	Height   int               `json:"height,omitempty"`   // 高度（像素）
	Size     int               `json:"size,omitempty"`     // 大小（字节）
	Ext      string            `json:"ext,omitempty"`      // 类型，如：mp4
	MD5      string            `json:"md5,omitempty"`      // 文件内容MD5
	Thumb    *ThumbMessageBody `json:"thumb,omitempty"`    // 缩略图
}
