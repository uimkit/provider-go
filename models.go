package uim

import "time"

const (
	AddAccountRequest     = "add_account"
	AddGroupRequest       = "add_group"
	AddGroupMemberRequest = "add_group_member"
	PostMessageRequest    = "post_message"
	GetAccountsIQ         = "get_accounts"
)

// IM用户
type IMUser struct {
	UserId          string         `json:"user_id,omitempty"`          // 平台用户ID，如：微信ID
	CustomId        string         `json:"custom_id,omitempty"`        // 用户自定义ID，如：微信号
	Username        string         `json:"username,omitempty"`         // 用户账户
	Name            string         `json:"name,omitempty"`             // 名称
	Mobile          string         `json:"mobile,omitempty"`           // 手机号
	Email           string         `json:"email,omitempty"`            // 邮箱
	Avatar          string         `json:"avatar,omitempty"`           // 头像URL
	QRCode          string         `json:"qrcode,omitempty"`           // 二维码URL
	Gender          Gender         `json:"gender,omitempty"`           // 性别
	Country         string         `json:"country,omitempty"`          // 国家
	Province        string         `json:"province,omitempty"`         // 省份
	City            string         `json:"city,omitempty"`             // 城市
	District        string         `json:"district,omitempty"`         // 区
	Address         string         `json:"address,omitempty"`          // 地址
	Signature       string         `json:"signature,omitempty"`        // 签名
	Birthday        *time.Time     `json:"birthday,omitempty"`         // 生日
	Language        string         `json:"language,omitempty"`         // 语言
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

// IM账号
type IMAccount struct {
	User            *IMUser        `json:"user,omitempty"`             // 用户信息
	Presence        Presence       `json:"presence,omitempty"`         // 状态
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

func (client *Client) AddAccount(account *IMAccount) error {
	return client.SendMessage(NewMessageRequest(AddAccountRequest, "", "", account))
}

// 群组
type Group struct {
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	Owner           *IMUser        `json:"owner,omitempty"`            // 群主信息
	Name            string         `json:"name,omitempty"`             // 名称
	Avatar          string         `json:"avatar,omitempty"`           // 头像URL
	Announcement    string         `json:"announcement,omitempty"`     // 群公告
	Description     string         `json:"description,omitempty"`      // 群介绍
	MemberCount     int32          `json:"member_count,omitempty"`     // 群成员数量
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

func (client *Client) AddGroup(group *Group) error {
	return client.SendMessage(NewMessageRequest(AddGroupRequest, "", "", group))
}

// 群组成员
type GroupMember struct {
	GroupId         string         `json:"group_id,omitempty"`         // 平台群组ID
	MemberId        string         `json:"member_id,omitempty"`        // 平台群成员ID
	User            *IMUser        `json:"user,omitempty"`             // 关联的用户信息
	IsOwner         bool           `json:"is_owner,omitempty"`         // 是否群主
	IsAdmin         bool           `json:"is_admin,omitempty"`         // 是否管理员
	Alias           string         `json:"alias,omitempty"`            // 群内备注名
	Metadata        map[string]any `json:"metadata,omitempty"`         // 公开元数据
	PrivateMetadata map[string]any `json:"private_metadata,omitempty"` // 私有元数据
}

func (client *Client) AddGroupMember(member *GroupMember) error {
	return client.SendMessage(NewMessageRequest(AddGroupMemberRequest, "", "", member))
}

// 消息
type Message struct {
	MessageId        string           `json:"message_id,omitempty"`        // 平台消息ID
	ConversationType ConversationType `json:"conversation_type,omitempty"` // 所属的会话类型
	Seq              int              `json:"seq,omitempty"`               // 序列号，在会话中唯一且有序增长，用于确保消息顺序
	MentionedType    MentionedType    `json:"mentioned_type,omitempty"`    // @用户类型
	MentionedUsers   []*IMUser        `json:"mentioned_users"`             // @用户列表
	SentAt           time.Time        `json:"sent_at,omitempty"`           // 发送时间
	Payload          *MessagePayload  `json:"payload,omitempty"`           // 消息内容
	Revoked          bool             `json:"revoked,omitempty"`           // 是否撤回
	Metadata         map[string]any   `json:"metadata,omitempty"`          // 公开元数据
	PrivateMetadata  map[string]any   `json:"private_metadata,omitempty"`  // 私有元数据
}

func (client *Client) PostMessage(message *Message) error {
	return client.SendMessage(NewMessageRequest(PostMessageRequest, "", "", message))
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
